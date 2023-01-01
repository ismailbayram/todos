package api

import (
	"github.com/ismailbayram/todos/src/tests"
	"github.com/ismailbayram/todos/src/todos"
	"github.com/ismailbayram/todos/src/users"
	"github.com/stretchr/testify/suite"
	"testing"
)

type APITestSuite struct {
	tests.AppTestSuite
}

func TestAPITestSuite(t *testing.T) {
	userTestSuite := new(APITestSuite)
	userTestSuite.Models = []interface{}{&users.User{}, &todos.ToDo{}}
	suite.Run(t, userTestSuite)
}
