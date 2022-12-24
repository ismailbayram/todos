package main

import (
	"github.com/ismailbayram/todos/config"
	"github.com/ismailbayram/todos/src/database"
	"github.com/ismailbayram/todos/src/models"
)

func main() {
	cfg := config.Init()
	db := database.New(&cfg.Database)

	tx := db.DBConn.Begin()

	user1 := &models.User{
		Username: "user1",
		Password: "123456",
	}
	user2 := &models.User{
		Username: "user2",
		Password: "123456",
	}
	tx.Create(user1)
	tx.SavePoint("sp1")
	tx.Create(user2)
	tx.Commit()
	tx.RollbackTo("sp1")
	tx.Commit()
}
