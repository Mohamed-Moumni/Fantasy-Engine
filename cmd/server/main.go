package main

import (
	"log"

	"fantasy-engine/internal/config"
	"fantasy-engine/internal/database"
	"fantasy-engine/internal/services"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewConnection(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to the database")

	// playerService := services.NewPlayerService(db)

	// playerService.

	// services.Scrap_teams()

	if err := database.AutoMigrate(db); err != nil {
		log.Fatal(err)
	}
	log.Println("Database migrated")

	teams := services.Scrap_teams()
	teamService := services.NewTeamService(db)
	for _, team := range teams {
		teamService.Create(&team)
	}
}
