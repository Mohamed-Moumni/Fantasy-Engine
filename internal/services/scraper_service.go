package services

import (
	"encoding/json"
	"fantasy-engine/internal/models"
	"fmt"
	"io"
	"net/http"
	"os"
)


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

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error: ", err)
		return []models.Team{}
	}

	fmt.Println("Status: ", resp.Status)
	// fmt.Println("Body: ", string(body))

	for _, team := range response.Teams {
		fmt.Println(team.Name)
		fmt.Println(team.Slug)
		fmt.Println(team.ShortName)
		fmt.Println(team.Gender)
		fmt.Println(team.NameCode)
		fmt.Println(team.Country.Alpha2)
		fmt.Println(team.Country.Alpha3)
		fmt.Println(team.Country.Name)
		fmt.Println(team.Country.Slug)
		fmt.Println(team.TeamColors.Primary)
		fmt.Println(team.TeamColors.Secondary)
		fmt.Println(team.TeamColors.TextColor)
	}
	return response.Teams
}