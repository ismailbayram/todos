package repositories

import (
	"fmt"
	"github.com/ismailbayram/todos/src/database"
	"github.com/stretchr/testify/suite"
	"testing"
)

func (s *ToDoTestSuite) SetupSuite() {
	dbName := "test_todos"

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s", "postgres", "123456", "localhost", "5432")
	s.db = database.New(dsn)
	s.db.DBConn.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbName))

	dsn = fmt.Sprintf("%s/%s", dsn, dbName)
	s.db = database.New(dsn)
	s.db.Migrate()
}

func (s *ToDoTestSuite) TearDownSuite() {
	db, _ := s.db.DBConn.DB()
	err := db.Close()
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s", "postgres", "123456", "localhost", "5432")
	s.db = database.New(dsn)
	s.db.DBConn.Exec("DROP DATABASE test_todos;")
}

func (s *ToDoTestSuite) SetupTest() {
	s.tx = s.db.DBConn.Begin()
	result := s.tx.SavePoint("sp")
	if result.Error != nil {
		panic(result.Error)
	}
}

func (s *ToDoTestSuite) TearDownTest() {
	result := s.tx.RollbackTo("sp")
	if result.Error != nil {
		panic(result.Error)
	}
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(ToDoTestSuite))
}
