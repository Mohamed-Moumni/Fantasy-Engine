package main

import (
	// "fmt"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"fantasy-engine/internal/config"
	"fantasy-engine/internal/database"
	"fantasy-engine/internal/models"
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

	content, err := os.ReadFile("/home/moumni/Desktop/Fantasy-Engine/cmd/server/team.json")

	if err != nil {
		log.Fatal(err)
	}

	var team models.Team

	err = json.Unmarshal(content, &team)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Team: %+v\n", team)

	var TeamService *services.TeamService

	TeamService = services.NewTeamService(db)

	err = TeamService.Create(&team)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Team created successfully")


	// log.Println("Connected to the database")

	// playerService := services.NewPlayerService(db)

	// playerService.

	// services.Scrap_teams()

	// if err := database.AutoMigrate(db); err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("Database migrated")

	// countryService := services.NewCountryService(db)

	// var countryObject = models.Country{
	// 	ID: 1,
	// 	Alpha2: "MA",
	// 	Alpha3: "MAR",
	// 	Name: "Morocco",
	// 	Slug: "morocco",
	// }

	// CountryErr := countryService.Create(&countryObject)

	// delErr := countryService.DeleteByID(1)

	// if delErr != nil {
		// fmt.Println(delErr.Error())
	// }
	// fmt.Println(country.Name)
	// fmt.Println(country)
	// countryObject.Name = "Marecous"
	// countryService.Update(&countryObject)

	// if countryService != nil {
		// fmt.Println(CountryErr)
	// }
}
