package repository

import (
	// "fantasy-engine/internal/database"
	"fantasy-engine/internal/models"

	"gorm.io/gorm"
)

type MatchRepository struct {
	db *gorm.DB
}

func NewMatchService(db *gorm.DB) *MatchRepository {
	return &MatchRepository{db: db}
}

func (s *MatchRepository) Create(match *models.Match) (*models.Match, error) {
	var Match models.Match
	s.db.FirstOrCreate(&Match, *match)
	return &Match, nil
}

func (s *MatchRepository) Update(match *models.Match) error {
	result := s.db.Save(match)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *MatchRepository) GetByTeamsAndGameWeek(homeID uint, gameWeekRound uint) (*models.Match, error) {
	var match models.Match
	result := s.db.Where(s.db.Where("home_team_id = ? AND gameweek = ?", homeID, gameWeekRound)).Or(s.db.Where("away_team_id = ? AND gameweek = ?", homeID, gameWeekRound)).First(&match)
	if result.Error != nil {
		return nil, result.Error
	}
	return &match, nil
}
