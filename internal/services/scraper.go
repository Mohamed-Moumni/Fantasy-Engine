package services

import (
	"encoding/json"
	"fantasy-engine/internal/models"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type PlayerDTO struct {
    ID                     uint    `json:"id"`
    Name                   string  `json:"name"`
    Slug                   string  `json:"slug"`
    ShortName              string  `json:"shortName"`
    Position               string  `json:"position"`
    JerseyNumber           string  `json:"jerseyNumber"`
    Height                 float32 `json:"height"`
    Gender                 string  `json:"gender"`
    DateOfBirthTimestamp   int64   `json:"dateOfBirthTimestamp"` // Unix timestamp as int
    ShirtNumber            uint    `json:"shirtNumber"`
    ProposedMarketValueRaw *struct {
        Value    uint   `json:"value"`
        Currency string `json:"currency"`
    } `json:"proposedMarketValueRaw"`
    Country struct {
        Alpha2 string `json:"alpha2"`
        Alpha3 string `json:"alpha3"`
        Name   string `json:"name"`
        Slug   string `json:"slug"`
    } `json:"country"`
}

// ------------------------------------------------ SCRAP TEAMS ------------------------------- //
func ScrapTeams() []models.Team {
	resp, err := http.Get(os.Getenv("SCRAPER_TEAM_URL"))
	if err != nil {
		fmt.Println("Error:", err)
		return []models.Team{}
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response: ", err)
		return []models.Team{}
	}

	var response struct {
		Teams []models.Team `json:"teams"`
	}

	fmt.Println("Body: ", string(body))

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error: ", err)
		return []models.Team{}
	}

	return response.Teams
}

// ------------------------------------------------ SCRAP PLAYERS ------------------------------- //

func ScrapPlayers(slug string, teamId uint) []PlayerDTO {
	url := os.Getenv("SCRAPER_PLAYER_URL") + slug + "/" + strconv.FormatUint(uint64(teamId), 10)
	fmt.Println("URL: ", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return []PlayerDTO{}
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	var response struct {
		Players []PlayerDTO `json:"players"`
		Total   int `json:"total"`
	}

	fmt.Println("Body: ", string(body))

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error: ", err)
		return []PlayerDTO{}
	}
	return response.Players
}


// ------------------------------------------------ SCRAP MATCHES ------------------------------- //