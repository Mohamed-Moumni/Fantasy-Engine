package main

import (
	"log"

	"fantasy-engine/internal/config"
	"fantasy-engine/internal/database"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	_, err = database.NewConnection(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
}
