package users

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
)

func (s *UserTestSuite) TestLoginView() {
	ur := NewUserRepository(s.DB)
	user, err := ur.Create("hilal", "123456", true)
	assert.Nil(s.T(), err)

	handler := LoginView(s.DB)
	var payload map[string]any

	// Wrong data scenario
	reqBody, _ := json.Marshal(map[string]string{"wrong": "wrong"})
	req, _ := http.NewRequest(http.MethodPost, "/api/login/", bytes.NewReader(reqBody))
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	assert.Equal(s.T(), http.StatusBadRequest, response.Code)
	resp, _ := io.ReadAll(response.Body)
	err = json.Unmarshal(resp, &payload)
	if err != nil {
		panic(err)
	}
	assert.Equal(s.T(), payload["username"], "required")
	assert.Equal(s.T(), payload["password"], "required")

	// Wrong password scenario
	reqBody, _ = json.Marshal(map[string]string{
		"username": user.Username,
		"password": "WRONG",
	})
	req, _ = http.NewRequest(http.MethodPost, "/api/login/", bytes.NewReader(reqBody))
	response = httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	assert.Equal(s.T(), http.StatusBadRequest, response.Code)
	resp, _ = io.ReadAll(response.Body)
	err = json.Unmarshal(resp, &payload)
	if err != nil {
		panic(err)
	}
	assert.Equal(s.T(), payload["username"], "Incorrect username or password.")

	// true password scenario
	reqBody, _ = json.Marshal(map[string]string{
		"username": user.Username,
		"password": "123456",
	})
	req, _ = http.NewRequest(http.MethodPost, "/api/login/", bytes.NewReader(reqBody))
	response = httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	assert.Equal(s.T(), http.StatusOK, response.Code)
	resp, _ = io.ReadAll(response.Body)
	err = json.Unmarshal(resp, &payload)
	if err != nil {
		panic(err)
	}
	assert.NotNil(s.T(), payload["token"])
}
