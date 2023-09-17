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
	DBUrl      string `mapstructure:"DB_URL"`
}

var cfg Config

func init() {
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	viper.SetDefault("SERVER_HOST", "127.0.0.1")
	viper.SetDefault("SERVER_PORT", "3000")
	viper.SetDefault("DB_URL", "")

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
