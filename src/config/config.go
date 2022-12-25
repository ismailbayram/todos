package config

import (
	"github.com/spf13/viper"
	"os"
)

type Configuration struct {
	Database DatabaseConfiguration
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

	// TODO: bindEnvs()

	return config
}

func getDefaultConfig() *Configuration {
	return &Configuration{
		Database: DatabaseConfiguration{
			Name:     "todos",
			Username: "postgres",
			Password: "123456",
			Host:     "localhost",
			Port:     "5432",
		},
	}
}

func readConfiguration() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		if _, err := os.Stat("./config/config.yml"); os.IsNotExist(err) {
			os.Create("./config/config.yml")
		} else {
			panic(err)
		}
	}
}

//func bindEnvs() {
//	viper.BindEnv("database.host", "DB_HOST")
//	viper.BindEnv("database.username", "DB_USER")
//}
