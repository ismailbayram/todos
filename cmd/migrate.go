package main

import (
	"github.com/ismailbayram/todos/config"
	"github.com/ismailbayram/todos/src/database"
	"github.com/ismailbayram/todos/src/todos"
	"github.com/ismailbayram/todos/src/users"
	"log"
)

func main() {
	cfg := config.Init()
	db := database.New(&cfg.Database)
	models := []interface{}{users.User{}, todos.ToDo{}}
	db.Migrate(models)

	db = database.New(&cfg.Database)
	ur := users.NewUserRepository(db.Conn)
	_, err := ur.Create("admin", "123456", true)
	if err != nil {
		log.Fatalln(err)
	}
}
