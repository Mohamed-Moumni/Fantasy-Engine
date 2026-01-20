package core_api

import (
	"core_api/tests/testutil"
	"testing"

	"core_api/internal/repository"
	"pkg/models"

	"github.com/stretchr/testify/assert"

	"gorm.io/gorm"
)

func TestAll(t *testing.T) {
	db, cleanup := testutil.SetupTestDB(t)
	t.Cleanup(cleanup)

	t.Run("CreateUser", func(t *testing.T) { CreateUser(t, db) })
	t.Run("GetUserByID", func(t *testing.T) { GetUserByID(t, db) })
	t.Run("GetUserByEmail", func(t *testing.T) { GetUserByEmail(t, db) })
	t.Run("DuplicateEmailUser", func(t *testing.T) { DuplicateEmailUser(t, db) })
	t.Run("UpdateUser", func(t *testing.T) { UpdateUser(t, db) })
}

func CreateUser(t *testing.T, db *gorm.DB) {
	// Create test data
	country := models.Country{
		ID:   1,
		Name: "Test Country",
		Slug: "test-country",
	}
	favoriteTeam := models.Team{
		ID:        1,
		Name:      "Test Team",
		CountryID: country.ID,
	}
	favoritePlayer := models.Player{
		ID:        1,
		Name:      "Test Player",
		Position:  "Test Position",
		CountryID: country.ID,
	}
	// Create user to test
	user := models.User{
		ID:               testutil.USER_ID,
		Email:            "test@example.com",
		Username:         "testuser",
		Age:              20,
		FavoriteTeamID:   favoriteTeam.ID,
		FavoritePlayerID: favoritePlayer.ID,
		CountryID:        country.ID,
		Gender:           "M",
		PasswordHash:     "password",
	}

	userRepository := repository.NewUserRepository(db)
	err := userRepository.CreateUser(&user)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	assert.NoError(t, err)
	assert.Equal(t, "test@example.com", user.Email)
	assert.Equal(t, "testuser", user.Username)
	assert.Equal(t, uint(20), user.Age)
	assert.Equal(t, favoriteTeam.ID, user.FavoriteTeamID)
	assert.Equal(t, favoritePlayer.ID, user.FavoritePlayerID)
	assert.Equal(t, country.ID, user.CountryID)
	assert.Equal(t, "M", user.Gender)
	assert.Equal(t, "password", user.PasswordHash)
}

func GetUserByID(t *testing.T, db *gorm.DB) {

	userRepository := repository.NewUserRepository(db)
	user, err := userRepository.GetUserByID(testutil.USER_ID)
	if err != nil {
		t.Fatalf("Failed to get user: %v", err)
	}
	assert.NoError(t, err)
	assert.Equal(t, "test@example.com", user.Email)
}

func GetUserByEmail(t *testing.T, db *gorm.DB) {

	userRepository := repository.NewUserRepository(db)
	user, err := userRepository.GetUserByEmail("test@example.com")
	if err != nil {
		t.Fatalf("Failed to get user: %v", err)
	}
	assert.NoError(t, err)
	assert.Equal(t, "test@example.com", user.Email)
}

func DuplicateEmailUser(t *testing.T, db *gorm.DB) {
	user := models.User{
		Email:            "test@example.com",
		Username:         "testuser",
		Age:              20,
		FavoriteTeamID:   1,
		FavoritePlayerID: 1,
		CountryID:        1,
		Gender:           "M",
		PasswordHash:     "password",
	}
	userRepository := repository.NewUserRepository(db)
	err := userRepository.CreateUser(&user)
	assert.Error(t, err)
}

func UpdateUser(t *testing.T, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	user, err := userRepository.GetUserByID(testutil.USER_ID)
	if err != nil {
		t.Fatalf("Failed to get user: %v", err)
	}
	user.Email = "test2@example.com"
	err = userRepository.UpdateUser(user)
	assert.NoError(t, err)
	assert.Equal(t, "test2@example.com", user.Email)
}
