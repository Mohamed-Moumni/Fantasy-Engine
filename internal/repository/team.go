package repository

import (
	"fantasy-engine/internal/models"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type TeamRepository struct {
	db             *gorm.DB
	CountryService *CountryRepository
}

// -------------------------------------------  team  ----------------------------------------

func NewTeamRepository(db *gorm.DB) *TeamRepository {
	return &TeamRepository{
		db:             db,
		CountryService: NewCountryRepository(db),
	}
}

// create - update - delete - get(team)

func (s *TeamRepository) Create(team *models.Team) error {

	teamColor, err := s.CreateTeamColor(&team.TeamColors)
	if err != nil {
		log.Printf("Error creating team colors: %v", err)
		return err
	}

	team.TeamColorsID = team.ID
	team.TeamColors = *teamColor

	country, err := s.CountryService.Create(&team.Country)
	if err != nil {
		if err.Error() != "Country with this Name does Exist" {
			log.Printf("Error creating country: %v", err)
			return err
		}
	}

	team.CountryID = country.ID
	team.Country = *country

	if err := s.db.Create(team).Error; err != nil {
		return fmt.Errorf("failed to create team: %w", err)
	}
	return nil
}

func (s *TeamRepository) DeleteByName(name string) error {
	var team models.Team
	res := s.db.Where("name = ?", name).First(&team)
	if res.Error != nil {
		fmt.Errorf("Team not Found with this Name: %s", name)
		return res.Error
	}
	s.db.Delete(&team)
	return nil
}

func (s *TeamRepository) GetTeamByNameCode(nameCode string) (*models.Team, error) {
	var team models.Team
	res := s.db.Where("name_code = ?", nameCode).First(&team)
	if res.Error != nil {
		fmt.Errorf("Team not Found with this Name Code: %s", nameCode)
		return &models.Team{}, res.Error
	}
	return &team, nil
}

func (s *TeamRepository) GetTeamByName(name string) (*models.Team, error) {
	var team models.Team

	res := s.db.First(&team, name)
	if res.Error != nil {
		fmt.Errorf("Team not Found with this Name: %s", name)
		return &models.Team{}, res.Error
	}
	return &team, nil
}

// -------------------------------------------  team colors   ----------------------------------------
func (s *TeamRepository) CreateTeamColor(teamColor *models.TeamColor) (*models.TeamColor, error) {
	var TeamColor models.TeamColor
	s.db.FirstOrCreate(&TeamColor, *teamColor)
	if err := s.db.Create(teamColor).Error; err != nil {
		return teamColor, fmt.Errorf("Failed to create TeamColor: %w", err)
	}
	return &TeamColor, nil
}

func (s *TeamRepository) GetTeamByID(id uint) (*models.Team, error) {
	var team models.Team

	res := s.db.First(&team, id)
	if res.Error != nil {
		fmt.Errorf("Team not Found with this ID: %d", id)
		return &models.Team{}, res.Error
	}
	return &team, nil
}

func (s *TeamRepository) GetAllTeams() ([]models.Team, error) {
	var teams []models.Team
	res := s.db.Find(&teams)
	if res.Error != nil {
		return []models.Team{}, fmt.Errorf("Failed to get all teams: %w", res.Error)
	}
	return teams, nil
}