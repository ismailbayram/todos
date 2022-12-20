package models

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

func (s *ToDoTestSuite) SetupSuite() {
	var err error

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s", "postgres", "123456", "localhost", "5432")
	s.db, err = gorm.Open(postgres.Open(dsn))
	require.NoError(s.T(), err)

	s.db.Exec("CREATE DATABASE test_todos;")
}

func (s *ToDoTestSuite) TearDownSuite() {
	//s.db.Exec("DROP DATABASE test_todos;")
}

func (s *ToDoTestSuite) SetupTest() {
	s.tx = s.db.Begin()
}

func (s *ToDoTestSuite) TearDownTest() {
	s.tx.Rollback()
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(ToDoTestSuite))
}
