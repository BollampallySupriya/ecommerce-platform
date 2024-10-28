package helpers

import "os"

type Config struct {
	Port         string
	DATABASE_URL string
}

func LoadConfig() *Config {
	cfg := &Config{
		Port: os.Getenv("PORT"),
		DATABASE_URL: os.Getenv("DATABASE_URL"),
	}
	return cfg
}