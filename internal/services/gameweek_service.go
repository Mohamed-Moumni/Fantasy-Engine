package services

import (
	"fantasy-engine/internal/models"

	"gorm.io/gorm"
)

type GameweekService struct {
	db *gorm.DB
}

func NewGameweekService(db *gorm.DB) *GameweekService {
	return &GameweekService{db: db}
}

func (s *GameweekService) Create(gameweek *models.GameWeek) (*models.GameWeek, error) {
	var Gameweek models.GameWeek
	s.db.FirstOrCreate(&Gameweek, *gameweek)
	return &Gameweek, nil
}