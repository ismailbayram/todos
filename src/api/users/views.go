package users

import (
	"encoding/json"
	"log"
	"net/http"
)

func LoginView(w http.ResponseWriter, r *http.Request) {
	response, _ := json.Marshal(map[string]string{"name": "iso"})
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(response)
	if err != nil {
		log.Fatalln(err)
	}
}
