package main

import (
	// "fmt"
	"encoding/json"
	// "fmt"
	"fmt"
	"log"
	"os"

	// "os"

	"fantasy-engine/internal/config"
	"fantasy-engine/internal/database"
	"fantasy-engine/internal/services"

	// "fantasy-engine/internal/services"
	"fantasy-engine/internal/models"
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

	body, err := os.ReadFile("/Users/mac/Desktop/Fantasy-Engine/cmd/server/teams.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Body: ", string(body))

	var teams []models.Team
	err = json.Unmarshal(body, &teams)
	if err != nil {
		log.Fatal(err)
	}

	teamService := services.NewTeamService(db)
	for _, team := range teams {
		teamService.Create(&team)
	}

	os.Exit(0)
}
