package main

import (
	"fmt"
	"github.com/ismailbayram/todos/config"
)

func main() {
	//dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", "postgres", "123456", "localhost", "5432", "todos")
	//db := database.New(dsn)
	//
	//ur := repositories.NewUserRepository(db.DBConn)
	////user, _ := ur.Create("ismail", "123456", true)
	//var user models.User
	//db.DBConn.First(&user)
	////if result.Error != nil {
	////	panic(result.Error)
	////}
	//ur.Deactivate(&user)

	cfg := config.Init()
	fmt.Println(cfg.Database.Host)
	fmt.Println(cfg.Database.Username)
	fmt.Println(cfg.Database.Password)
}
