package services

import (
	"fantasy-engine/internal/models"
	// "fmt"
	"log"

	"gorm.io/gorm"
)

// TODO- CREATE, UPDATE, DELETE, GET

type CountryService struct {
	db *gorm.DB
}

func NewCountryService(db *gorm.DB) *CountryService {
	return &CountryService{db: db}
}

func (s *CountryService) Create(country *models.Country) (*models.Country, error) {
	var Country models.Country
	s.db.FirstOrCreate(&Country, *country)
	return &Country, nil
}

func (s *CountryService) Update(country *models.Country) error {
	result := s.db.Save(country)

	if result.Error != nil {
		return result.Error
	}
	log.Printf("Country with ID: %d is updated", country.ID)
	return nil
}

func (s *CountryService) DeleteByID(id uint) error {
	var country models.Country
	result := s.db.First(&country, id)

	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	s.db.Delete(&country)
	return nil
}

func (s *CountryService) GetCountryByID(id uint) (*models.Country, error) {
	var country models.Country

	res := s.db.First(&country, id)
	if res.Error != nil {
		log.Fatalf("Country not Found with this ID: %d", id)
		return &models.Country{}, res.Error
	}
	return &country, nil
}