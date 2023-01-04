package config

import (
	"github.com/spf13/viper"
	"os"
)

type Configuration struct {
	SecretKey string
	Database  DatabaseConfiguration
	Server    ServerConfiguration
}

type ServerConfiguration struct {
	Port    string
	Timeout uint
}

type DatabaseConfiguration struct {
	Name     string
	Username string
	Password string
	Host     string
	Port     string
}

func Init() *Configuration {
	config := getDefaultConfig()
	readConfiguration()

	err := viper.Unmarshal(config)
	if err != nil {
		panic(err)
	}

	return config
}

func getDefaultConfig() *Configuration {
	return &Configuration{
		SecretKey: "TOP_SECRET",
		Database: DatabaseConfiguration{
			Name:     "todos",
			Username: "postgres",
			Password: "123456",
			Host:     "localhost",
			Port:     "5432",
		},
		Server: ServerConfiguration{
			Port:    "8000",
			Timeout: 10,
		},
	}
}

func readConfiguration() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	bindEnvs()

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		if _, err := os.Stat("./config/config.yml"); os.IsNotExist(err) {
			os.Create("./config/config.yml")
		}
	}
}

func bindEnvs() {
	viper.BindEnv("database.host", "DB_HOST")
	viper.BindEnv("database.port", "DB_PORT")
	viper.BindEnv("database.name", "DB_NAME")
	viper.BindEnv("database.username", "DB_USER")
	viper.BindEnv("database.password", "DB_PASSWORD")

	viper.BindEnv("server.port", "SW_PORT")
}
