package users

import (
	"encoding/json"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
)

func LoginView(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var responseData = map[string]string{}
		requestData, _ := io.ReadAll(r.Body)
		var loginDTO LoginDTO
		err := json.Unmarshal(requestData, &loginDTO)
		if err != nil {
			log.Fatalln(err)
		}
		requestDataIsValid := true
		if loginDTO.Username == "" {
			responseData["username"] = "Username is required."
			requestDataIsValid = false
		}
		if loginDTO.Password == "" {
			responseData["password"] = "Password is required."
			requestDataIsValid = false
		}
		if !requestDataIsValid {
			response, _ := json.Marshal(responseData)
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write(response)
			return
		}

		ur := NewUserRepository(db)
		user, err := ur.GetByUsername(loginDTO.Username)
		if err != nil {
			responseData["username"] = "Incorrect username or password."
			response, _ := json.Marshal(responseData)
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write(response)
			return
		}
		if !ur.CheckUserPassword(user, loginDTO.Password) {
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
