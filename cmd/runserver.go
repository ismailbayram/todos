package main

import (
	"fmt"
	"github.com/ismailbayram/todos/src/api"
	"github.com/ismailbayram/todos/src/config"
	"github.com/ismailbayram/todos/src/database"
	"log"
	"net/http"
	"time"
)

func main() {
	cfg := config.Init()

	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		WriteTimeout: time.Second * time.Duration(cfg.Server.Timeout),
		ReadTimeout:  time.Second * time.Duration(cfg.Server.Timeout),
		Handler:      api.NewRouter(database.New(&cfg.Database).Conn),
	}

	log.Println(fmt.Sprintf("Listening on http://%s:%s", cfg.Server.Host, cfg.Server.Port))
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
