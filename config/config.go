package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	DB *DatabaseConfig
}

func Load() (*Config, error) {
	err := godotenv.Load(".env", ".env.services")
	if err != nil {
		log.Println("Config: some env files not loaded: ", err.Error())
	}

	// let viper auto load the env files
	viper.AutomaticEnv()

	return &Config{
		DB: &DatabaseConfig{
			Driver: viper.GetString("DB_DRIVER"),
			Host: viper.GetString("DB_HOST"),
			Port: viper.GetInt("DB_PORT"),
			User: viper.GetString("DB_USER"),
			Pass: viper.GetString("DB_PASSWORD"), 
			Name: viper.GetString("DB_NAME"),
			Charset: viper.GetString("DB_CHARSET"),
		},
	}, nil
}