package api

import (
	"github.com/ismailbayram/todos/src/users"
	"gorm.io/gorm"
	"net/http"
)

func LoginView(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responseData := make(ResponseData)
		var loginDTO LoginDTO

		errors := ValidateRequestData(r.Body, &loginDTO)
		if errors != nil {
			Respond(w, errors, http.StatusBadRequest)
			return
		}

		ur := users.NewUserRepository(db)
		user, err := ur.GetByUsername(loginDTO.Username)
		if err != nil || !ur.CheckUserPassword(user, loginDTO.Password) {
			responseData["username"] = "Incorrect username or password."
			Respond(w, responseData, http.StatusBadRequest)
			return
		}

		token, err := ur.CreateToken(user)
		if err != nil {
			responseData["username"] = "Something went wrong, please try again."
			Respond(w, responseData, http.StatusInternalServerError)
			return
		}

		responseData["token"] = token
		Respond(w, responseData, http.StatusOK)
	}
}

func UserListView(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !IsAuthorized(w, r.Context()) {
			return
		}
		if !IsAdmin(w, r.Context()) {
			return
		}
		responseData := make(ResponseData)

		ur := users.NewUserRepository(db)
		userList, err := ur.All(ConvertQueryToFilter(
			r.URL.Query(),
			[]string{"id", "is_active", "is_admin"},
		))

		if err != nil {
			responseData["error"] = "An error has been occured, please try again."
			Respond(w, responseData, http.StatusInternalServerError)
			return
		}
		count := len(userList)
		responseData["count"] = count
		responseData["results"] = make([]UserDTO, count)
		for i, user := range userList {
			responseData["results"].([]UserDTO)[i] = ToUserDTO(&user)
		}
		Respond(w, responseData, http.StatusOK)
	}
}

func UserCreateView(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !IsAuthorized(w, r.Context()) {
			return
		}
		if !IsAdmin(w, r.Context()) {
			return
		}
		responseData := make(ResponseData)
		var userCreateDTO UserCreateDTO

		errors := ValidateRequestData(r.Body, &userCreateDTO)
		if errors != nil {
			Respond(w, errors, http.StatusBadRequest)
			return
		}

		ur := users.NewUserRepository(db)
		_, err := ur.GetByUsername(userCreateDTO.Username)
		if err == nil {
			responseData["username"] = "There is already a user registered with this username."
			Respond(w, responseData, http.StatusBadRequest)
			return
		}

		_, err = ur.Create(userCreateDTO.Username, userCreateDTO.Password, *userCreateDTO.IsAdmin)
		if err != nil {
			responseData["username"] = "Something went wrong, please try again."
			Respond(w, responseData, http.StatusInternalServerError)
			return
		}
		Respond(w, responseData, http.StatusCreated)
	}
}
