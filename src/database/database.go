package database

import (
	"fmt"
	"github.com/ismailbayram/todos/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DBConn *gorm.DB
}

func New(dsn string) *Database {
	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic(fmt.Sprintf("Database error: %s", err))
	}

	return &Database{
		DBConn: db,
	}
}

func (db *Database) Migrate() {
	err := db.DBConn.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}
}
