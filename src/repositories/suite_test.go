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
	db, _ := s.db.DBConn.DB()
	_ = db.Close()
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
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", "postgres", "123456", "localhost", "5432", "test_todos")
	s.db = database.New(dsn)
}

func (s *ToDoTestSuite) TearDownTest() {
	db, _ := s.db.DBConn.DB()
	_ = db.Close()
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(ToDoTestSuite))
}
