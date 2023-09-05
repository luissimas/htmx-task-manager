package config

import (
	"log"

	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Host string
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	viper.SetDefault("api.host", "127.0.0.1")
	viper.SetDefault("api.port", "3000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")

	if err := loadConfig(); err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}
}

func loadConfig() error {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Print("Config file not found, falling back to default config")
		} else {
			return err
		}
	}

	cfg = new(config)
	cfg.API = APIConfig{
		Host: viper.GetString("api.host"),
		Port: viper.GetString("api.port"),
	}
	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		Database: viper.GetString("database.database"),
	}

	return nil
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetAPI() APIConfig {
	return cfg.API
}
