package core_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	"core_api/internal/repository"
	"pkg/database"
	"pkg/models"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/gorm"
)

const USER_ID = "845df16a-db46-4767-a553-92d822023461"

func setupTestDB(t *testing.T) (*gorm.DB, func()) {
	ctx := context.Background()

	// Start PostgreSQL container
	pgContainer, err := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithDatabase("test_db"),
		postgres.WithUsername("test_user"),
		postgres.WithPassword("test_password"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(60*time.Second),
		),
	)
	if err != nil {
		t.Fatalf("Failed to start postgres container: %v", err)
	}

	// Get connection string
	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to get connection string: %v", err)
	}

	fmt.Println(connStr)

	// Connect with GORM
	db, err := database.NewConnection(connStr)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations
	err = database.AutoMigrate(db)
	if err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}

	// Return cleanup function that terminates the container
	cleanup := func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Logf("Failed to terminate container: %v", err)
		}
	}
	return db, cleanup
}

func TestAll(t *testing.T) {
	db, cleanup := setupTestDB(t)
	t.Cleanup(cleanup)

	t.Run("CreateUser", func(t *testing.T) { CreateUser(t, db) })
	t.Run("GetUserByID", func(t *testing.T) { GetUserByID(t, db) })
	t.Run("GetUserByEmail", func(t *testing.T) { GetUserByEmail(t, db) })
	t.Run("DuplicateEmailUser", func(t *testing.T) { DuplicateEmailUser(t, db) })
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
		ID:               USER_ID,
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
	user, err := userRepository.GetUserByID(USER_ID)
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
