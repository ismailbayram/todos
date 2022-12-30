package users

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"strings"
)

func LoginView(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var responseData = map[string]any{}
		requestData, _ := io.ReadAll(r.Body)
		var loginDTO LoginDTO
		err := json.Unmarshal(requestData, &loginDTO)
		if err != nil {
			log.Fatalln(err)
		}

		requestDataValidator := validator.New()

		errors := map[string]string{}
		if err := requestDataValidator.Struct(loginDTO); err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				errors[strings.ToLower(err.Field())] = err.Tag()
			}
			response, _ := json.Marshal(errors)
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write(response)
			return
		}

		ur := NewUserRepository(db)
		user, err := ur.GetByUsername(loginDTO.Username)
		if err != nil || !ur.CheckUserPassword(user, loginDTO.Password) {
			responseData["username"] = "Incorrect username or password."
			response, _ := json.Marshal(responseData)
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write(response)
			return
		}

		token, err := ur.CreateToken(user, "a")
		if err != nil {
			responseData["username"] = "Something went wrong, please try again."
			response, _ := json.Marshal(responseData)
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write(response)
			return
		}

		responseData["token"] = token
		response, _ := json.Marshal(responseData)
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(response)
	}
}
