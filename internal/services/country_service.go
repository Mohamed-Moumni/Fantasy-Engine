package services

import (
	"fantasy-engine/internal/models"
	"fmt"

	"gorm.io/gorm"
)

type CountryService struct {
	db *gorm.DB
}

func NewCountryService(db *gorm.DB) *CountryService {
	return &CountryService{db: db}
}

func (s *CountryService) Create(country *models.Country) error {
	if err := s.db.Where("name = ?", country.Name).First(&models.Country{}).Error; err == nil {
		return nil
	}
	if err := s.db.Create(country).Error; err != nil {
		return fmt.Errorf("failed to create country: %w", err)
	}
	return nil
}
