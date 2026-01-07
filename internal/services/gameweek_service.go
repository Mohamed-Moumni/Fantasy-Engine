package services

import (
	"fantasy-engine/internal/models"
	"log"

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

func (s *GameweekService) GetByRound(round uint) (*models.GameWeek, error) {
	var gameweek models.GameWeek
	result := s.db.Where("round = ?", round).First(&gameweek)
	if result.Error != nil {
		return nil, result.Error
	}
	return &gameweek, nil
}

func (s *GameweekService) Update(gameweek *models.GameWeek) error {
	result := s.db.Save(gameweek)

	if result.Error != nil {
		return result.Error
	}
	log.Printf("GameWeek with ID: %d is updated", gameweek.ID)
	return nil
}