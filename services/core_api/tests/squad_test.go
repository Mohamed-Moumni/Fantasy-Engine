package core_api

import (
	"core_api/internal/repository"
	"core_api/tests/testutil"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"pkg/models"
	"testing"
)

func TestSquadAll(t *testing.T) {
	db, cleanup := testutil.SetupTestDB(t)
	t.Cleanup(cleanup)

	t.Run("Create Squad", func(t *testing.T) { CreateSquad(t, db) })
	t.Run("Create Squad Player", func(t *testing.T) { CreateSquadPlayer(t, db) })
}

func CreateSquad(t *testing.T, db *gorm.DB) {
	squad := models.Squad{
		Name: "Test Squad",
	}
	squadRepository := repository.NewSquadRepository(db)
	err := squadRepository.CreateSquad(&squad, testutil.USER_ID)
	if err != nil {
		t.Fatalf("Failed to create squad: %v", err)
	}
	assert.NoError(t, err)
	assert.Equal(t, "Test Squad", squad.Name)
}

func CreateSquadPlayer(t *testing.T, db *gorm.DB) {
	squadPlayer := models.SquadPlayer{
		SquadID:       1,
		PlayerID:      1,
		Captain:       false,
		ViceCaptain:   false,
		Injured:       false,
		Suspended:     false,
		Score:         0,
		Price:         5.2,
		IsStarting:    false,
		PositionOrder: 0,
	}
	squadPlayerRepository := repository.NewSquadRepository(db)
	err := squadPlayerRepository.CreateSquadPlayer(&squadPlayer)
	if err != nil {
		t.Fatalf("Failed to create squad player: %v", err)
	}
	assert.NoError(t, err)
	assert.Equal(t, uint(1), squadPlayer.SquadID)
	assert.Equal(t, uint(1), squadPlayer.PlayerID)
	assert.Equal(t, false, squadPlayer.Captain)
	assert.Equal(t, false, squadPlayer.ViceCaptain)
	assert.Equal(t, false, squadPlayer.Injured)
	assert.Equal(t, false, squadPlayer.Suspended)
	assert.Equal(t, uint(0), squadPlayer.Score)
	assert.Equal(t, float32(5.2), squadPlayer.Price)
	assert.Equal(t, false, squadPlayer.IsStarting)
	assert.Equal(t, uint(0), squadPlayer.PositionOrder)
}
