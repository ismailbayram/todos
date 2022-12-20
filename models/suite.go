package models

import (
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type ToDoTestSuite struct {
	suite.Suite
	db *gorm.DB
	tx *gorm.DB
}
