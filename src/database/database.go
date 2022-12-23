package database

import (
	"fmt"
	"github.com/ismailbayram/todos/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Database struct {
	DBConn *gorm.DB
}

func New(dsn string) *Database {
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level (Silent, Error, Warn, Info)
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: dbLogger,
	})

	if err != nil {
		panic(fmt.Sprintf("Database error: %s", err))
	}

	return &Database{
		DBConn: db,
	}
}

func (db *Database) Migrate() {
	err := db.DBConn.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}
}
