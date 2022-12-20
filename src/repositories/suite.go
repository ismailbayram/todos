package repositories

import (
	"github.com/ismailbayram/todos/src/database"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type ToDoTestSuite struct {
	suite.Suite
	db *database.Database
	tx *gorm.DB
}
