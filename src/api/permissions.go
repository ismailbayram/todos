package api

import (
	"context"
	"github.com/ismailbayram/todos/src/users"
	"net/http"
)

func IsAuthorized(w http.ResponseWriter, ctx context.Context) bool {
	requestUser := ctx.Value("user")
	if requestUser == nil {
		RespondWithError(w, nil, http.StatusUnauthorized)
		return false
	}
	return true
}

func IsAdmin(w http.ResponseWriter, ctx context.Context) bool {
	requestUser := ctx.Value("user")
	if !requestUser.(*users.User).IsAdmin {
		RespondWithError(w, nil, http.StatusForbidden)
		return false
	}
	return true
}
