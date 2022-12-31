package users

import (
	"github.com/ismailbayram/todos/src/api"
	"gorm.io/gorm"
	"net/http"
)

func LoginView(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responseData := make(api.ResponseData)
		var loginDTO LoginDTO

		errors := api.ValidateRequestData(r.Body, &loginDTO)
		if errors != nil {
			api.RespondWithError(w, errors, http.StatusBadRequest)
			return
		}

		ur := NewUserRepository(db)
		user, err := ur.GetByUsername(loginDTO.Username)
		if err != nil || !ur.CheckUserPassword(user, loginDTO.Password) {
			responseData["username"] = "Incorrect username or password."
			api.RespondWithError(w, responseData, http.StatusBadRequest)
			return
		}

		token, err := ur.CreateToken(user)
		if err != nil {
			responseData["username"] = "Something went wrong, please try again."
			api.RespondWithError(w, responseData, http.StatusInternalServerError)
			return
		}

		responseData["token"] = token
		api.RespondWithSuccess(w, responseData)
	}
}

func UserListView(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestUser := r.Context().Value("user")
		if requestUser == nil {
			api.RespondWithError(w, nil, http.StatusUnauthorized)
			return

		}
		if !requestUser.(*User).IsAdmin {
			api.RespondWithError(w, nil, http.StatusForbidden)
			return
		}
		responseData := make(api.ResponseData)

		ur := NewUserRepository(db)
		users, err := ur.All(api.ConvertQueryToFilter(
			r.URL.Query(),
			[]string{"id", "is_active", "is_admin"},
		))

		if err != nil {
			responseData["error"] = "An error has been occured, please try again."
			api.RespondWithError(w, responseData, http.StatusInternalServerError)
			return
		}
		count := len(users)
		responseData["count"] = count
		responseData["results"] = make([]UserDTO, count)
		for i, user := range users {
			responseData["results"].([]UserDTO)[i] = ToUserDTO(&user)
		}
		api.RespondWithSuccess(w, responseData)
	}
}
