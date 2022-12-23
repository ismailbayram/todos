package main

import (
	"fmt"
	"github.com/ismailbayram/todos/src/database"
	"github.com/ismailbayram/todos/src/models"
	"github.com/ismailbayram/todos/src/repositories"
)

func main() {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", "postgres", "123456", "localhost", "5432", "todos")
	db := database.New(dsn)

	ur := repositories.NewUserRepository(db.DBConn)
	//user, _ := ur.CreateUser("ismail", "123456", true)
	var user models.User
	db.DBConn.First(&user)
	//if result.Error != nil {
	//	panic(result.Error)
	//}
	ur.DeactivateUser(&user)
}
