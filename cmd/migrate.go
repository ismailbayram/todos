package main

import (
	"github.com/ismailbayram/todos/src/config"
	"github.com/ismailbayram/todos/src/database"
	"github.com/ismailbayram/todos/src/users"
)

func main() {
	cfg := config.Init()
	db := database.New(&cfg.Database)
	models := []interface{}{users.User{}}
	db.Migrate(models)
}
