package repository

import (
	"fmt"

	"pkg/models"
	"gorm.io/gorm"
)

type SquadRepository struct {
	db *gorm.DB
}

func NewSquadRepository(db *gorm.DB) *SquadRepository {
	return &SquadRepository{db: db}
}

// create squad:
// first create the squad
// then associate the squad to specific user by updating the user's squad ID
func (sr *SquadRepository) CreateSquad(squad *models.Squad, userId string) error {
	result := sr.db.Create(squad)
	if result.Error != nil {
		return result.Error
	}

	userRepository := NewUserRepository(sr.db)
	user, err := userRepository.GetUserByID(userId)
	fmt.Println("user", user)
	if err != nil {
		return err
	}

	// update the user's squad ID
	user.SquadID = &squad.ID
	userRepository.UpdateUser(user)
	return nil
}
