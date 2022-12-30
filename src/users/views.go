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
			api.RespondWithError(w, errors)
			return
		}

		ur := NewUserRepository(db)
		user, err := ur.GetByUsername(loginDTO.Username)
		if err != nil || !ur.CheckUserPassword(user, loginDTO.Password) {
			responseData["username"] = "Incorrect username or password."
			api.RespondWithError(w, responseData)
			return
		}

		token, err := ur.CreateToken(user, "a")
		if err != nil {
			responseData["username"] = "Something went wrong, please try again."
			api.RespondWithError(w, responseData)
			return
		}

		responseData["token"] = token
		api.RespondWithSuccess(w, responseData)
	}
}
