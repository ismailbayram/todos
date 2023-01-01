package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ismailbayram/todos/src/todos"
	"github.com/ismailbayram/todos/src/users"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
)

func (s *APITestSuite) TestToDoListView() {
	ur := users.NewUserRepository(s.DB)
	admin, _ := ur.Create("ismail", "123456", true)
	hilal, _ := ur.Create("hilal", "123456", false)
	tdr := todos.NewToDoRepository(s.DB)
	todo1, _ := tdr.Create("admin todo 1", *admin)
	todo2, _ := tdr.Create("admin todo 2", *admin)
	todo3, _ := tdr.Create("hilal todo 1", *hilal)

	handler := ToDoListView(s.DB)
	var payload struct {
		Count   int
		Results []ToDoDTO
	}
	ctx := context.Background()

	// without token
	ctx = context.WithValue(ctx, "user", nil)
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "/api/todos/", nil)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	assert.Equal(s.T(), http.StatusUnauthorized, response.Code)

	// with admin
	ctx = context.WithValue(ctx, "user", admin)
	req, _ = http.NewRequestWithContext(ctx, http.MethodGet, "/api/todos/", nil)
	response = httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	assert.Equal(s.T(), http.StatusOK, response.Code)
	resp, _ := io.ReadAll(response.Body)
	json.Unmarshal(resp, &payload)

	assert.Equal(s.T(), 2, payload.Count)
	assert.Equal(s.T(), todo1.ID, payload.Results[0].ID)
	assert.Equal(s.T(), todo1.Name, payload.Results[0].Name)
	assert.Equal(s.T(), todo1.IsDone, payload.Results[0].IsDone)

	assert.Equal(s.T(), todo2.ID, payload.Results[1].ID)
	assert.Equal(s.T(), todo2.Name, payload.Results[1].Name)
	assert.Equal(s.T(), todo2.IsDone, payload.Results[1].IsDone)

	// with hilal
	ctx = context.WithValue(ctx, "user", hilal)
	req, _ = http.NewRequestWithContext(ctx, http.MethodGet, "/api/todos/", nil)
	response = httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	assert.Equal(s.T(), http.StatusOK, response.Code)
	resp, _ = io.ReadAll(response.Body)
	json.Unmarshal(resp, &payload)

	assert.Equal(s.T(), 1, payload.Count)
	assert.Equal(s.T(), todo3.ID, payload.Results[0].ID)
	assert.Equal(s.T(), todo3.Name, payload.Results[0].Name)
	assert.Equal(s.T(), todo3.IsDone, payload.Results[0].IsDone)
}

func (s *APITestSuite) TestToDoCreateView() {
	ur := users.NewUserRepository(s.DB)
	admin, _ := ur.Create("ismail", "123456", true)

	handler := ToDoCreateView(s.DB)
	var payload map[string]any
	ctx := context.Background()

	// without token
	ctx = context.WithValue(ctx, "user", nil)
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, "/api/todos/", nil)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	assert.Equal(s.T(), http.StatusUnauthorized, response.Code)

	// with unappropriated data
	reqBody, _ := json.Marshal(map[string]string{
		"username": "new",
		"password": "password1",
	})
	ctx = context.WithValue(ctx, "user", admin)
	req, _ = http.NewRequestWithContext(ctx, http.MethodPost, "/api/todos/", bytes.NewReader(reqBody))
	response = httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	assert.Equal(s.T(), http.StatusBadRequest, response.Code)
	resp, _ := io.ReadAll(response.Body)
	err := json.Unmarshal(resp, &payload)
	if err != nil {
		panic(err)
	}
	assert.Equal(s.T(), payload["name"], "required")

	// final scenario
	reqBody, _ = json.Marshal(map[string]any{
		"name": "new to do",
	})
	ctx = context.WithValue(ctx, "user", admin)
	req, _ = http.NewRequestWithContext(ctx, http.MethodPost, "/api/todos/", bytes.NewReader(reqBody))
	response = httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	assert.Equal(s.T(), http.StatusCreated, response.Code)
	resp, _ = io.ReadAll(response.Body)
	err = json.Unmarshal(resp, &payload)
	if err != nil {
		panic(err)
	}
}

func (s *APITestSuite) TestToDoUpdateView() {
	ur := users.NewUserRepository(s.DB)
	user1, _ := ur.Create("user1", "123456", true)
	user2, _ := ur.Create("user2", "123456", true)
	tdr := todos.NewToDoRepository(s.DB)
	todo1, err := tdr.Create("user1 todo", *user1)
	todo2, err := tdr.Create("user2 todo", *user2)

	handler := ToDoUpdateView(s.DB)
	var payload map[string]any
	ctx := context.Background()

	// without token
	ctx = context.WithValue(ctx, "user", nil)
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("/api/todos/%d/", todo1.ID), nil)
	req = mux.SetURLVars(req, map[string]string{
		"id": fmt.Sprint(todo1.ID),
	})
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	assert.Equal(s.T(), http.StatusUnauthorized, response.Code)

	// wrong ownership
	reqBody, _ := json.Marshal(map[string]string{
		"name": "new name",
	})
	ctx = context.WithValue(ctx, "user", user1)
	req, _ = http.NewRequestWithContext(ctx, http.MethodPut, fmt.Sprintf("/api/todos/%d/", todo2.ID), bytes.NewReader(reqBody))
	req = mux.SetURLVars(req, map[string]string{
		"id": fmt.Sprint(todo2.ID),
	})
	response = httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	assert.Equal(s.T(), http.StatusNotFound, response.Code)

	// non-existent todo
	reqBody, _ = json.Marshal(map[string]string{
		"name": "new name",
	})
	ctx = context.WithValue(ctx, "user", user1)
	req, _ = http.NewRequestWithContext(ctx, http.MethodPut, fmt.Sprintf("/api/todos/%d/", 321), bytes.NewReader(reqBody))
	req = mux.SetURLVars(req, map[string]string{
		"id": "321",
	})
	response = httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	assert.Equal(s.T(), http.StatusNotFound, response.Code)

	// final scenario
	reqBody, _ = json.Marshal(map[string]any{
		"name":    "new to do",
		"is_done": true,
	})
	ctx = context.WithValue(ctx, "user", user1)
	req, _ = http.NewRequestWithContext(ctx, http.MethodPut, fmt.Sprintf("/api/todos/%d/", todo1.ID), bytes.NewReader(reqBody))
	req = mux.SetURLVars(req, map[string]string{
		"id": fmt.Sprint(todo1.ID),
	})
	response = httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	assert.Equal(s.T(), http.StatusOK, response.Code)
	resp, _ := io.ReadAll(response.Body)
	err = json.Unmarshal(resp, &payload)
	if err != nil {
		panic(err)
	}
}
