package api

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"io"
	"log"
	"reflect"
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
			errors[getFieldName(s, err)] = err.Tag()
		}
		return errors
	}
	return nil
}

func getFieldName(s any, err validator.FieldError) string {
	field, _ := reflect.TypeOf(s).Elem().FieldByName(err.Field())
	fieldName := field.Tag.Get("json")
	if fieldName == "" {
		fieldName = strings.ToLower(err.Field())
	}
	return fieldName
}
