package repository

import (
	"fantasy-engine/internal/models"
	// "fmt"
	"log"

	"gorm.io/gorm"
)

// TODO- CREATE, UPDATE, DELETE, GET

type CountryRepository struct {
	db *gorm.DB
}

func NewCountryRepository(db *gorm.DB) *CountryRepository {
	return &CountryRepository{db: db}
}

func (s *CountryRepository) Create(country *models.Country) (*models.Country, error) {
	var Country models.Country
	s.db.FirstOrCreate(&Country, *country)
	return &Country, nil
}

func (s *CountryRepository) Update(country *models.Country) error {
	result := s.db.Save(country)

	if result.Error != nil {
		return result.Error
	}
	log.Printf("Country with ID: %d is updated", country.ID)
	return nil
}

func (s *CountryRepository) DeleteByID(id uint) error {
	var country models.Country
	result := s.db.First(&country, id)

	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	s.db.Delete(&country)
	return nil
}

func (s *CountryRepository) GetCountryByID(id uint) (*models.Country, error) {
	var country models.Country

	res := s.db.First(&country, id)
	if res.Error != nil {
		log.Printf("Country not Found with this ID: %d", id)
		return &models.Country{}, res.Error
	}
	return &country, nil
}
func (s *CountryRepository) GetCountryByName(name string) (*models.Country, error) {
	var country models.Country

	res := s.db.First(&country, "name = ?", name)
	if res.Error != nil {
		log.Printf("Country not Found with this Name: %s", name)
		return &models.Country{}, res.Error
	}
	return &country, nil
}