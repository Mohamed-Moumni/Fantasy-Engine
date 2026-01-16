package repository

import (
	"pkg/models"
	"fmt"
	"gorm.io/gorm"
)

type SquadRepository struct {
	db *gorm.DB
}

func NewSquadRepository(db *gorm.DB) *SquadRepository {
	return &SquadRepository{db: db}
}

func (sr *SquadRepository) CreateSquad(squad *models.Squad, userId string) error {
	result := sr.db.Create(squad)
	if result.Error != nil {
		return result.Error
	}

	userRepository := NewUserRepository(sr.db)
	user, err := userRepository.GetUserByID(userId)
	if err != nil {
		return fmt.Errorf("failed to get user by id: %w")
	}

	// update the user's squad ID
	user.SquadID = &squad.ID
	userRepository.UpdateUser(user)
	return nil
}

func (sr *SquadRepository) CreateSquadPlayer(squadPlayer *models.SquadPlayer) error {
	result := sr.db.Create(squadPlayer)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (sr *SquadRepository) GetSquadByID(id int32) (*models.Squad, error) {
	var squad models.Squad
	if err := sr.db.First(&squad, id).Error; err != nil {
		return nil, err
	}
	return &squad, nil
}

func (sr *SquadRepository) GetSquadPlayersBySquadID(squadID int32) ([]models.SquadPlayer, error) {
	var squadPlayers []models.SquadPlayer
	if err := sr.db.Where("squad_id = ?", squadID).Preload("Player").Find(&squadPlayers).Error; err != nil {
		return nil, err
	}
	return squadPlayers, nil
}