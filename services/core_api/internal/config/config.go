package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env            string
	Port           string
	DatabaseURL    string
	LogLevel       string
	ScraperBaseURL string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load(".core_api.env")

	return &Config{
		Port:           os.Getenv("PORT"),
		DatabaseURL:    os.Getenv("DATABASE_URL"),
		LogLevel:       os.Getenv("LOG_LEVEL"),
		ScraperBaseURL: os.Getenv("SCRAPER_BASE_URL"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
