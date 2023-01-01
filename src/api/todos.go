package api

import (
	"github.com/gorilla/mux"
	"github.com/ismailbayram/todos/src/todos"
	"github.com/ismailbayram/todos/src/users"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

func ToDoListView(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !IsAuthorized(w, r.Context()) {
			return
		}
		responseData := make(ResponseData)

		tdr := todos.NewToDoRepository(db)
		toDoList, err := tdr.GetByUserID(r.Context().Value("user").(*users.User).ID)

		if err != nil {
			responseData["error"] = "An error has been occured, please try again."
			Respond(w, responseData, http.StatusInternalServerError)
			log.Fatalln(err)
			return
		}
		count := len(toDoList)
		responseData["count"] = count
		responseData["results"] = make([]ToDoDTO, count)
		for i, todo := range toDoList {
			responseData["results"].([]ToDoDTO)[i] = ToToDoDTO(&todo)
		}
		Respond(w, responseData, http.StatusOK)
	}
}

func ToDoCreateView(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !IsAuthorized(w, r.Context()) {
			return
		}
		responseData := make(ResponseData)
		var toDoCreateDTO ToDoCreateDTO

		errors := ValidateRequestData(r.Body, &toDoCreateDTO)
		if errors != nil {
			Respond(w, errors, http.StatusBadRequest)
			return
		}

		tdr := todos.NewToDoRepository(db)
		user := r.Context().Value("user").(*users.User)
		_, err := tdr.Create(toDoCreateDTO.Name, *user)
		if err != nil {
			responseData["error"] = "Something went wrong, please try again."
			Respond(w, responseData, http.StatusInternalServerError)
			log.Fatalln(err)
			return
		}
		Respond(w, responseData, http.StatusCreated)
	}
}

func ToDoUpdateView(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !IsAuthorized(w, r.Context()) {
			return
		}
		responseData := make(ResponseData)

		tdr := todos.NewToDoRepository(db)
		toDoID, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			responseData["error"] = "To Do Not Found"
			Respond(w, responseData, http.StatusNotFound)
			return
		}
		toDo, err := tdr.GetByID(toDoID)
		if err != nil {
			responseData["error"] = "To Do Not Found"
			Respond(w, responseData, http.StatusNotFound)
			return
		}
		requestUser := r.Context().Value("user").(*users.User)
		if requestUser.ID != toDo.UserID {
			responseData["error"] = "To Do Not Found"
			Respond(w, responseData, http.StatusNotFound)
			return
		}

		var toDoUpdateDTO ToDoUpdateDTO
		errors := ValidateRequestData(r.Body, &toDoUpdateDTO)
		if errors != nil {
			Respond(w, errors, http.StatusBadRequest)
			return
		}
		toDo.IsDone = *toDoUpdateDTO.IsDone
		toDo.Name = toDoUpdateDTO.Name
		err = tdr.Update(toDo)
		if err != nil {
			responseData["error"] = "Something went wrong, please try again."
			Respond(w, responseData, http.StatusInternalServerError)
			log.Fatalln(err)
			return
		}
		Respond(w, responseData, http.StatusOK)
	}
}
