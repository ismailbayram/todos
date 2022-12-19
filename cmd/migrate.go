package main

import (
	"fmt"
	"github.com/ismailbayram/todos/apps/users"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", "postgres", "123456", "localhost", "5432", "todos")
	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic(fmt.Sprintf("Database error: %s", err))
	}
	db.AutoMigrate(&users.User{})
}
