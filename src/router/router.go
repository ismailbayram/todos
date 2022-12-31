package router

import (
	"github.com/gorilla/mux"
	"github.com/ismailbayram/todos/src/users"
	"gorm.io/gorm"
	"net/http"
)

func NewRouter(db *gorm.DB) *mux.Router {
	router := mux.NewRouter().PathPrefix("/api").Subrouter()
	router.Use(loggingMiddleware)
	router.Use(jsonMiddleware)
	router.Use(authenticationMiddleware(db))

	router.HandleFunc("/login/", users.LoginView(db)).Methods(http.MethodPost)
	router.HandleFunc("/users/", users.UserListView(db)).Methods(http.MethodGet)

	//router.NotFoundHandler = router.NewRoute().HandlerFunc(http.NotFound).GetHandler()

	return router
}
