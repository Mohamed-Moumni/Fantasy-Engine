package repository

import (
	"fantasy-engine/internal/models"

	"gorm.io/gorm"
)

type PlayerRepository struct {
	db *gorm.DB
}

// -------------------------------------------  player  ----------------------------------------

func NewPlayerRepository(db *gorm.DB) *PlayerRepository {
	return &PlayerRepository{
		db: db,
	}
}

// create - update - delete - get(player)

func (s *PlayerRepository) Create(player *models.Player) (*models.Player, error) {
	var createPlayer models.Player
	result := s.db.FirstOrCreate(&createPlayer, *player)
	if result.Error != nil {
		return nil, result.Error
	}
	return &createPlayer, nil
}

// func (s *PlayerService) Update(player *models.Player) (*models.Player, error) {

// }


