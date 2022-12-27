package users

import (
	"github.com/ismailbayram/todos/src/tests"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserTestSuite struct {
	tests.AppTestSuite
}

func TestUserTestSuite(t *testing.T) {
	userTestSuite := new(UserTestSuite)
	userTestSuite.Models = []interface{}{&User{}}
	suite.Run(t, userTestSuite)
}
