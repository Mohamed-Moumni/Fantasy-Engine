package services

import (
	"errors"
	"fantasy-engine/internal/models"
	"fmt"

	"gorm.io/gorm"
)

type PlayerService struct {
	db *gorm.DB
}

func NewPlayerService(db *gorm.DB) *PlayerService {
	return &PlayerService{db: db}
}

func (s *PlayerService) Create(player *models.Player) error {
	if err := s.db.Create(player).Error; err != nil {
		return fmt.Errorf("failed to create player: %w", err)
	}
	return nil
}

func (s *PlayerService) GetByID(id uint) (*models.Player, error) {
	var player models.Player

	if err := s.db.First(&player, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Player not found")
		}
		return nil, fmt.Errorf("Failed to get player: %w", err)
	}
	return &player, nil
}
