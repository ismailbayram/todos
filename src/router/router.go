package router

import (
	"github.com/gorilla/mux"
	"github.com/ismailbayram/todos/src/api"
	"gorm.io/gorm"
	"net/http"
)

func NewRouter(db *gorm.DB) *mux.Router {
	router := mux.NewRouter().PathPrefix("/api").Subrouter()
	router.Use(loggingMiddleware)
	router.Use(jsonMiddleware)
	router.Use(authenticationMiddleware(db))

	router.HandleFunc("/login/", api.LoginView(db)).Methods(http.MethodPost)
	router.HandleFunc("/users/", api.UserListView(db)).Methods(http.MethodGet)
	router.HandleFunc("/users/", api.UserCreateView(db)).Methods(http.MethodPost)
	router.HandleFunc("/todos/", api.ToDoListView(db)).Methods(http.MethodGet)
	router.HandleFunc("/todos/", api.ToDoCreateView(db)).Methods(http.MethodPost)
	router.HandleFunc("/todos/{id}/", api.ToDoUpdateView(db)).Methods(http.MethodPut)

	//router.NotFoundHandler = router.NewRoute().HandlerFunc(http.NotFound).GetHandler()

	return router
}
