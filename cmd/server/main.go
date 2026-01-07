package main

import (
	// "fmt"
	// "encoding/json"
	// "time"

	// "fmt"
	"fmt"
	"log"
	"time"

	// "os"

	// "os"

	"fantasy-engine/internal/config"
	"fantasy-engine/internal/database"
	"fantasy-engine/internal/models"
	"fantasy-engine/internal/repository"
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

	teamRepo := repository.NewTeamService(db)
	playerRepo := repository.NewPlayerService(db)
	countryRepo := repository.NewCountryRepository(db)
	teams, err := teamRepo.GetAllTeams()
	if err != nil {
		log.Fatal(err)
	}
	for _, team := range teams {
		fmt.Println("Scraping players for team: ", team.Name, team.ID)
		players := services.ScrapPlayers(team.Slug, team.ID)
		for _, playerDto := range players {
			var country *models.Country
			country, err = countryRepo.GetCountryByName(playerDto.Country.Name)
			if err != nil {
				log.Printf("Country not Found with this Name: %s", playerDto.Country.Name)
				country, err = countryRepo.Create(&models.Country{
					Alpha2: playerDto.Country.Alpha2,
					Alpha3: playerDto.Country.Alpha3,
					Name:   playerDto.Country.Name,
					Slug:   playerDto.Country.Slug,
				})
			}
			player := models.Player{
				ID:                   playerDto.ID,
				Name:                 playerDto.Name,
				Slug:                 playerDto.Slug,
				ShortName:            playerDto.ShortName,
				Position:             playerDto.Position,
				JerseyNumber:         playerDto.JerseyNumber,
				Height:               playerDto.Height,
				Gender:               playerDto.Gender,
				CountryID:            country.ID,
				DateOfBirthTimestamp: time.Unix(playerDto.DateOfBirthTimestamp, 0),
				ShirtNumber:          playerDto.ShirtNumber,
				TeamID:               team.ID,
				Team:                 team,
			}
			if playerDto.ProposedMarketValueRaw != nil {
				player.ProposedMarketValueRaw = playerDto.ProposedMarketValueRaw.Value
			} else {
				player.ProposedMarketValueRaw = 0
			}

			playerRepo.Create(&player)
		}
	}
}
