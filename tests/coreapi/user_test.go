package core_api

import (
	"testing"
)

func TestCreateUser(t *testing.T) {
	Country := models.Country{
		ID: 1,
		Name: "Test Country",
		Slug: "Test Code",
	}
	favoriteTeam := models.Team{
		ID: 1,
		Name: "Test Team",
		CountryID: 1,
		Logo: "https://example.com/logo.png",
	}
	favoritePlayer := models.Player{
		ID: 1,
		Name: "Test Player",
		Position: "Test Position",
		CountryID: 1
	}
	user := models.User{
		Email: "test@example.com",
		Username: "testuser",
		Age: 20,
		FavoriteTeamID: 1,
		FavoritePlayerID: 1,
		CountryID: 1,
		Gender: "M",
		PasswordHash: "password",
	}

	userRepository := repository.NewUserRepository(db)
	user, err := userRepository.CreateUser(&user)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}
	
	assert.NoError(t, err)
	assert.Equal(t, user.Email, "test@example.com")
	assert.Equal(t, user.Username, "testuser")
	assert.Equal(t, user.Age, 20)
	assert.Equal(t, user.FavoriteTeamID, 1)
	assert.Equal(t, user.FavoritePlayerID, 1)
	assert.Equal(t, user.CountryID, 1)
	assert.Equal(t, user.Gender, "M")
	assert.Equal(t, user.PasswordHash, "password")
}