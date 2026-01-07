package repository

import (
	"encoding/json"
	"fantasy-engine/internal/models"
	"fmt"
	"io"
	"net/http"
	"os"
)

// ------------------------------------------------ SCRAP TEAMS ------------------------------- //
func Scrap_teams() []models.Team {
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


// ------------------------------------------------ SCRAP MATCHES ------------------------------- //