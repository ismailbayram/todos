package api

import (
	"encoding/json"
	"net/http"
)

type ResponseData map[string]any

func RespondWithError(w http.ResponseWriter, resp ResponseData, code int) {
	response, _ := json.Marshal(resp)
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithSuccess(w http.ResponseWriter, resp ResponseData) {
	response, _ := json.Marshal(resp)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
