package main

import (
	"github.com/ismailbayram/todos/src/config"
	"github.com/ismailbayram/todos/src/database"
	"github.com/ismailbayram/todos/src/users"
	"log"
)

func main() {
	cfg := config.Init()
	db := database.New(&cfg.Database)
	models := []interface{}{users.User{}}
	db.Migrate(models)

	db = database.New(&cfg.Database)
	ur := users.NewUserRepository(db.Conn)
	_, err := ur.Create("admin", "123456", true)
	if err != nil {
		log.Fatalln(err)
	}
}
