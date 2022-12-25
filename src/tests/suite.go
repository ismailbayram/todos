package tests

import (
	"fmt"
	"github.com/ismailbayram/todos/config"
	"github.com/ismailbayram/todos/src/database"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type AppTestSuite struct {
	suite.Suite
	Config *config.Configuration
	DB     *gorm.DB
	Models []interface{}
}

func (s *AppTestSuite) SetupSuite() {
	s.Config = config.Init()
	s.Config.Database.Name = fmt.Sprintf("test_%s", s.Config.Database.Name)
	database.SetupTestDatabase(s.Config.Database, s.Models)
}

func (s *AppTestSuite) TearDownSuite() {
	db, _ := s.DB.DB()
	err := db.Close()
	if err != nil {
		panic(err)
	}

	database.DropTestDatabase(s.Config.Database)
}

func (s *AppTestSuite) SetupTest() {
	db := database.New(&s.Config.Database)
	s.DB = db.Conn.Begin()
	s.DB.SavePoint("sp")
}

func (s *AppTestSuite) TearDownTest() {
	s.DB.RollbackTo("sp")
	s.DB.Commit()
	db, _ := s.DB.DB()
	_ = db.Close()
}
