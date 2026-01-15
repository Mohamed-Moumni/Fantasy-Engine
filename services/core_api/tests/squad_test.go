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
