package services

import (
	"fantasy-engine/internal/models"
	"fmt"

	"gorm.io/gorm"
)

type TeamService struct {
	db *gorm.DB
}

func NewTeamService(db *gorm.DB) *TeamService {
	return &TeamService{db: db}
}

func (s *TeamService) Create(team *models.Team) error {
	if err := s.db.Create(team).Error; err != nil {
		return fmt.Errorf("failed to create team: %w", err)
	}
	return nil
}


func (s *TeamService) CreateTeamColors(teamColors *models.TeamColor) error {
	if 
}