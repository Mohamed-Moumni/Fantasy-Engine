package service

import (
	"core_api/graph/model"
	"core_api/internal/repository"
	"core_api/internal/utils"
	"pkg/models"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

// register User
func (us *UserService) RegisterUser(user *model.UserRegistration) (*model.User, error) {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return &model.User{}, nil
	}
	createdUser := models.User{
		Username:        user.Username,
		Email:           user.Email,
		Age:             uint(user.Age),
		Gender:          user.Gender,
		PhoneNumber:     user.Phone,
		PasswordHash:    hashedPassword,
		PasswordActived: true,
	}
	if err := us.userRepository.CreateUser(&createdUser); err != nil {
		return nil, err
	}
	return &model.User{
		ID:       createdUser.ID,
		Email:    createdUser.Email,
		Username: createdUser.Username,
		Age:      int32(createdUser.Age),
		Gender:   createdUser.Gender,
	}, nil
}