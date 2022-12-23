package repositories

import (
	"github.com/ismailbayram/todos/config"
	"github.com/ismailbayram/todos/src/database"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type ToDoTestSuite struct {
	suite.Suite
	config *config.Configuration
	db     *database.Database
	tx     *gorm.DB
}
