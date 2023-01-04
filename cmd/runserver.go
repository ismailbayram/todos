package main

import (
	"fmt"
	"github.com/ismailbayram/todos/config"
	"github.com/ismailbayram/todos/src/database"
	"github.com/ismailbayram/todos/src/router"
	"log"
	"net/http"
	"time"
)

func main() {
	cfg := config.Init()

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Server.Port),
		WriteTimeout: time.Second * time.Duration(cfg.Server.Timeout),
		ReadTimeout:  time.Second * time.Duration(cfg.Server.Timeout),
		Handler:      router.NewRouter(database.New(&cfg.Database).Conn),
	}

	log.Println(fmt.Sprintf("Listening on http://127.0.0.1:%s", cfg.Server.Port))
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
