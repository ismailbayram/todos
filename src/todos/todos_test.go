package todos

import (
	"github.com/ismailbayram/todos/src/tests"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ToDoTestSuite struct {
	tests.AppTestSuite
}

func TestUserTestSuite(t *testing.T) {
	toDoTestSuite := new(ToDoTestSuite)
	toDoTestSuite.Models = []interface{}{&ToDo{}}
	suite.Run(t, toDoTestSuite)
}
