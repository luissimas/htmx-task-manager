package config

import (
	"errors"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	ServerHost string `mapstructure:"SERVER_HOST"`
	ServerPort string `mapstructure:"SERVER_PORT"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBDatabase string `mapstructure:"DB_DATABASE"`
}

var cfg Config

func init() {
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	viper.SetDefault("SERVER_HOST", "127.0.0.1")
	viper.SetDefault("SERVER_PORT", "3000")

	viper.SetDefault("DB_HOST", "127.0.0.1")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_USER", "")
	viper.SetDefault("DB_PASSWORD", "")
	viper.SetDefault("DB_DATABASE", "")

	if err := loadConfig(); err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}
}

func loadConfig() error {
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	viper.Unmarshal(&cfg)

	return nil
}

func GetConfig() Config {
	return cfg
}
