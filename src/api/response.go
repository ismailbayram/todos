package api

import (
	"encoding/json"
	"net/http"
)

type ResponseData map[string]any

func Respond(w http.ResponseWriter, resp ResponseData, code int) {
	response, _ := json.Marshal(resp)
	w.WriteHeader(code)
	w.Write(response)
}
