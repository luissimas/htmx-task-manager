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
	Path string
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	viper.SetDefault("api.host", "127.0.0.1")
	viper.SetDefault("api.port", "3000")
	viper.SetDefault("database.path", "database.db")

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
		Port: viper.GetString("api.port"),
	}
	cfg.DB = DBConfig{
		Path: viper.GetString("database.path"),
	}

	return nil
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetAPI() APIConfig {
	return cfg.API
}
