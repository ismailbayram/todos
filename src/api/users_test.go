package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ismailbayram/todos/src/users"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
)

func (s *APITestSuite) TestLoginView() {
	ur := users.NewUserRepository(s.DB)
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

func (s *APITestSuite) TestUserListView() {
	ur := users.NewUserRepository(s.DB)
	admin, _ := ur.Create("ismail", "123456", true)
	hilal, _ := ur.Create("hilal", "123456", true)
	fatih, _ := ur.Create("fatih", "123456", false)
	ur.Deactivate(fatih)

	handler := UserListView(s.DB)
	var payload struct {
		Count   int
		Results []UserDTO
	}
	ctx := context.Background()

	// without token
	ctx = context.WithValue(ctx, "user", nil)
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "/api/users/", nil)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	assert.Equal(s.T(), http.StatusUnauthorized, response.Code)

	// with unprivileged user's token
	ctx = context.WithValue(ctx, "user", fatih)
	req, _ = http.NewRequestWithContext(ctx, http.MethodGet, "/api/users/", nil)
	response = httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	assert.Equal(s.T(), http.StatusForbidden, response.Code)

	// with token
	ctx = context.WithValue(ctx, "user", admin)
	req, _ = http.NewRequestWithContext(ctx, http.MethodGet, "/api/users/", nil)
	response = httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	assert.Equal(s.T(), http.StatusOK, response.Code)
	resp, _ := io.ReadAll(response.Body)
	json.Unmarshal(resp, &payload)

	assert.Equal(s.T(), 3, payload.Count)
	assert.Equal(s.T(), admin.ID, payload.Results[0].ID)
	assert.Equal(s.T(), admin.Username, payload.Results[0].Username)

	assert.Equal(s.T(), hilal.ID, payload.Results[1].ID)
	assert.Equal(s.T(), hilal.Username, payload.Results[1].Username)

	assert.Equal(s.T(), fatih.ID, payload.Results[2].ID)
	assert.Equal(s.T(), fatih.Username, payload.Results[2].Username)

	// filtering by ID
	req, _ = http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("/api/users/?id=%d", admin.ID), nil)
	response = httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	assert.Equal(s.T(), http.StatusOK, response.Code)
	resp, _ = io.ReadAll(response.Body)
	json.Unmarshal(resp, &payload)
	assert.Equal(s.T(), 1, payload.Count)
	assert.Equal(s.T(), admin.ID, payload.Results[0].ID)
	assert.Equal(s.T(), admin.Username, payload.Results[0].Username)

	// filtering by is_admin
	req, _ = http.NewRequestWithContext(ctx, http.MethodGet, "/api/users/?is_admin=true", nil)
	response = httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	assert.Equal(s.T(), http.StatusOK, response.Code)
	resp, _ = io.ReadAll(response.Body)
	json.Unmarshal(resp, &payload)
	assert.Equal(s.T(), 2, payload.Count)
	assert.Equal(s.T(), admin.ID, payload.Results[0].ID)
	assert.Equal(s.T(), admin.Username, payload.Results[0].Username)
	assert.Equal(s.T(), hilal.ID, payload.Results[1].ID)
	assert.Equal(s.T(), hilal.Username, payload.Results[1].Username)

	// filtering by is_active
	req, _ = http.NewRequestWithContext(ctx, http.MethodGet, "/api/users/?is_active=false", nil)
	response = httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	assert.Equal(s.T(), http.StatusOK, response.Code)
	resp, _ = io.ReadAll(response.Body)
	json.Unmarshal(resp, &payload)
	assert.Equal(s.T(), 1, payload.Count)
	assert.Equal(s.T(), fatih.ID, payload.Results[0].ID)
	assert.Equal(s.T(), fatih.Username, payload.Results[0].Username)

	// filtering by is_admin and id
	req, _ = http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("/api/users/?is_admin=true&id=%d", admin.ID), nil)
	response = httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	assert.Equal(s.T(), http.StatusOK, response.Code)
	resp, _ = io.ReadAll(response.Body)
	json.Unmarshal(resp, &payload)
	assert.Equal(s.T(), 1, payload.Count)
	assert.Equal(s.T(), admin.ID, payload.Results[0].ID)
	assert.Equal(s.T(), admin.Username, payload.Results[0].Username)
}

func (s *APITestSuite) TestUserCreateView() {
	ur := users.NewUserRepository(s.DB)
	admin, _ := ur.Create("ismail", "123456", true)
	fatih, _ := ur.Create("fatih", "123456", false)

	handler := UserCreateView(s.DB)
	var payload map[string]any
	ctx := context.Background()

	// without token
	ctx = context.WithValue(ctx, "user", nil)
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, "/api/users/", nil)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	assert.Equal(s.T(), http.StatusUnauthorized, response.Code)

	// with unprivileged user's token
	ctx = context.WithValue(ctx, "user", fatih)
	req, _ = http.NewRequestWithContext(ctx, http.MethodPost, "/api/users/", nil)
	response = httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	assert.Equal(s.T(), http.StatusForbidden, response.Code)

	// with unappropriated data
	reqBody, _ := json.Marshal(map[string]string{
		"username": "new",
		"password": "password1",
	})
	ctx = context.WithValue(ctx, "user", admin)
	req, _ = http.NewRequestWithContext(ctx, http.MethodPost, "/api/users/", bytes.NewReader(reqBody))
	response = httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	assert.Equal(s.T(), http.StatusBadRequest, response.Code)
	resp, _ := io.ReadAll(response.Body)
	err := json.Unmarshal(resp, &payload)
	if err != nil {
		panic(err)
	}
	assert.Equal(s.T(), payload["is_admin"], "required")

	// with existed username
	reqBody, _ = json.Marshal(map[string]any{
		"username": "ismail",
		"password": "password1",
		"is_admin": false,
	})
	ctx = context.WithValue(ctx, "user", admin)
	req, _ = http.NewRequestWithContext(ctx, http.MethodPost, "/api/users/", bytes.NewReader(reqBody))
	response = httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	assert.Equal(s.T(), http.StatusBadRequest, response.Code)
	resp, _ = io.ReadAll(response.Body)
	err = json.Unmarshal(resp, &payload)
	if err != nil {
		panic(err)
	}
	assert.Equal(s.T(), payload["username"], "There is already a user registered with this username.")

	// final scenario
	reqBody, _ = json.Marshal(map[string]any{
		"username": "new",
		"password": "password1",
		"is_admin": false,
	})
	ctx = context.WithValue(ctx, "user", admin)
	req, _ = http.NewRequestWithContext(ctx, http.MethodPost, "/api/users/", bytes.NewReader(reqBody))
	response = httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	assert.Equal(s.T(), http.StatusCreated, response.Code)
	resp, _ = io.ReadAll(response.Body)
	err = json.Unmarshal(resp, &payload)
	if err != nil {
		panic(err)
	}
}
