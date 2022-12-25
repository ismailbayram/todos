package database

import (
	"fmt"
	"github.com/ismailbayram/todos/src/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Database struct {
	Conn *gorm.DB
}

func New(cfg *config.DatabaseConfiguration) *Database {
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level (Silent, Error, Warn, Info)
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	dsn := fmt.Sprintf("host=%s  port=%s user=%s password=%s dbname=%s", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Name)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: dbLogger,
	})

	if err != nil {
		panic(fmt.Sprintf("Database error: %s", err))
	}

	return &Database{
		Conn: db,
	}
}

func SetupTestDatabase(cfg config.DatabaseConfiguration, models []interface{}) {
	dsn := fmt.Sprintf("host=%s  port=%s user=%s password=%s", cfg.Host, cfg.Port, cfg.Username, cfg.Password)
	dbServer, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}
	dbServer.Exec(fmt.Sprintf("CREATE DATABASE %s;", cfg.Name))

	dsn = fmt.Sprintf("%s name=test_%s", dsn, cfg.Name)
	db := New(&cfg)
	db.Migrate(models)
}

func DropTestDatabase(cfg config.DatabaseConfiguration) {
	dsn := fmt.Sprintf("host=%s  port=%s user=%s password=%s", cfg.Host, cfg.Port, cfg.Username, cfg.Password)
	dbServer, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}
	dbServer.Exec(fmt.Sprintf("DROP DATABASE %s;", cfg.Name))
}

func (db *Database) Migrate(models []interface{}) {
	err := db.Conn.AutoMigrate(models...)
	if err != nil {
		panic(err)
	}
	conn, _ := db.Conn.DB()
	_ = conn.Close()
}
