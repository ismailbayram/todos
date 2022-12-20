package main

import (
	"fmt"
	"github.com/ismailbayram/todos/src/database"
)

func main() {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", "postgres", "123456", "localhost", "5432", "todos")
	db := database.New(dsn)
	db.Migrate()
}
