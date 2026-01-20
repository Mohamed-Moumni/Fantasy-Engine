package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env string
	Port string
	DatabaseURL string
	LogLevel string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load(".backend.env")

	return &Config{
		Env: os.Getenv("ENV"),
		Port: os.Getenv("PORT"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
		LogLevel: os.Getenv("LOG_LEVEL"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}