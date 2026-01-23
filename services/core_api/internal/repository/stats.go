package repository

import (
	"pkg/models"

	"gorm.io/gorm"
)

type StatsRepository struct {
	db *gorm.DB
}

func NewStatsRepository(db *gorm.DB) *StatsRepository {
	return &StatsRepository{db: db}
}

func (sr *StatsRepository) GetGameWeeks(gameWeeks *[]models.GameWeek) error {
	if err := sr.db.Preload("Matches").Find(&gameWeeks).Error; err != nil {
		return err
	}
	return nil
}

func (sr *StatsRepository) GetTables(tables *[]models.Table) error {
	if err := sr.db.Preload("Team").Find(&tables).Error; err != nil {
		return err
	}
	return nil
}
