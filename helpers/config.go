package helpers

import (
	"os"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	Port         string
	DATABASE_URL string
}

func LoadConfig() *Config {

	if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

	cfg := &Config{
		Port: os.Getenv("PORT"),
		DATABASE_URL: os.Getenv("DATABASE_URL"),
	}
	return cfg
}