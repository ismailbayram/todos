package tests

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

func (s *ToDoTestSuite) SetupSuite() {
	fmt.Println("setup todo suite")
}

func (s *ToDoTestSuite) TearDownSuite() {
	fmt.Println("tear down todo suite")
}

//func (s *ToDoTestSuite) SetupTest() {
//	s.tx = s.db.Begin()
//}
//
//func (s *ToDoTestSuite) TearDownTest() {
//	s.tx.Rollback()
//}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(ToDoTestSuite))
}
