package repository

import (
	"fantasy-engine/internal/models"
	"log"

	"gorm.io/gorm"
)

type GameWeekRepository struct {
	db *gorm.DB
}

func NewGameWeekRepository(db *gorm.DB) *GameWeekRepository {
	return &GameWeekRepository{db: db}
}

func (s *GameWeekRepository) Create(gameweek *models.GameWeek) (*models.GameWeek, error) {
	var Gameweek models.GameWeek
	s.db.FirstOrCreate(&Gameweek, *gameweek)
	return &Gameweek, nil
}

func (s *GameWeekRepository) GetByRound(round uint) (*models.GameWeek, error) {
	var gameweek models.GameWeek
	result := s.db.Where("round = ?", round).First(&gameweek)
	if result.Error != nil {
		return nil, result.Error
	}
	return &gameweek, nil
}

func (s *GameWeekRepository) Update(gameweek *models.GameWeek) error {
	result := s.db.Save(gameweek)

	if result.Error != nil {
		return result.Error
	}
	log.Printf("GameWeek with ID: %d is updated", gameweek.ID)
	return nil
}
