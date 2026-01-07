package services

import (
	// "fantasy-engine/internal/database"
	"fantasy-engine/internal/models"

	"gorm.io/gorm"
)

type MatchService struct {
	db *gorm.DB
}

func NewMatchService(db *gorm.DB) *MatchService {
	return &MatchService{db: db}
}

func (s *MatchService) Create(match *models.Match) (*models.Match, error) {
	var Match models.Match
	s.db.FirstOrCreate(&Match, *match)
	return &Match, nil
}