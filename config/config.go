package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type ConfigDatabase struct {
	Port     string `env:"DB_PORT"`
	Host     string `env:"DB_HOST" `
	Name     string `env:"DB_NAME" `
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
}

func LoadConfig() (ConfigDatabase, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	var cfg ConfigDatabase
	err = cleanenv.ReadEnv(&cfg)
	return cfg, err
}
