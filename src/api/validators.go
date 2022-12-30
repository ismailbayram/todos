package api

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"io"
	"log"
	"strings"
)

var requestDataValidator = validator.New()

func ValidateRequestData(data io.ReadCloser, s any) ResponseData {
	requestData, _ := io.ReadAll(data)

	err := json.Unmarshal(requestData, s)
	if err != nil {
		log.Fatalln(err)
	}

	errors := make(ResponseData)
	if err := requestDataValidator.Struct(s); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors[strings.ToLower(err.Field())] = err.Tag()
		}
		return errors
	}
	return nil
}
