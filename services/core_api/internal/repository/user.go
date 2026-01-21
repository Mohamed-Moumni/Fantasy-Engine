package repository

import (
	"pkg/models"

	"gorm.io/gorm"
)

// User Creation
// User Update
// Get User By Email, ID, Username

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// create user
func (r *UserRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

// get user by id
func (r *UserRepository) GetUserByID(id string) (*models.User, error) {
	var user models.User

	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// get user by email
func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	err := r.db.First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	result := r.db.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User

	err := r.db.First(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
