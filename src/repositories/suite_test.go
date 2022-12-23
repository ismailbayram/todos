package repositories

import (
	"fmt"
	"github.com/ismailbayram/todos/config"
	"github.com/ismailbayram/todos/src/database"
	"github.com/stretchr/testify/suite"
	"testing"
)

func (s *ToDoTestSuite) SetupSuite() {
	s.config = config.Init()
	s.config.Database.Name = fmt.Sprintf("test_%s", s.config.Database.Name)
	database.SetupTestDatabase(s.config.Database)
}

func (s *ToDoTestSuite) TearDownSuite() {
	db, _ := s.db.DBConn.DB()
	err := db.Close()
	if err != nil {
		panic(err)
	}

	database.DropTestDatabase(s.config.Database)
}

func (s *ToDoTestSuite) SetupTest() {
	s.db = database.New(&s.config.Database)
	s.tx = s.db.DBConn.Begin()
	s.tx.SavePoint("sp")
}

func (s *ToDoTestSuite) TearDownTest() {
	s.tx.RollbackTo("sp")
	s.tx.Commit()
	db, _ := s.db.DBConn.DB()
	_ = db.Close()
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(ToDoTestSuite))
}
