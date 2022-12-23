package main

import (
	"github.com/ismailbayram/todos/config"
	"github.com/ismailbayram/todos/src/database"
)

func main() {
	cfg := config.Init()
	db := database.New(&cfg.Database)
	db.Migrate()
}
