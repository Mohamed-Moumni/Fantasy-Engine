package service

import (
	"core_api/graph/model"
	"core_api/internal/repository"
	"core_api/internal/service"
	"core_api/tests/testutil"
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestSquadService(t *testing.T) {
	db, cleanup := testutil.SetupTestDB(t)
	t.Cleanup(cleanup)

	t.Run("Create Squad", func(t *testing.T) { CreateSquad(t, db) })
}

func CreateSquad(t *testing.T, db *gorm.DB) {
	squadService := service.NewSquadService(repository.NewSquadRepository(db))
	squad := model.CreateSquadInput{
		Name: "Test Squad",
		Players: []*model.CreateSquadPlayerInput{
			{
				PlayerID:      1,
				Captain:       false,
				ViceCaptain:   false,
				Injured:       false,
				Suspended:     false,
				Score:         0,
				Price:         0.0,
				IsStarting:    false,
				PositionOrder: 0,
			},
			{
				PlayerID:      2,
				Captain:       false,
				ViceCaptain:   false,
				Injured:       false,
				Suspended:     false,
				Score:         0,
				Price:         0.0,
				IsStarting:    false,
				PositionOrder: 1,
			},
			{
				PlayerID:      3,
				Captain:       false,
				ViceCaptain:   false,
				Injured:       false,
				Suspended:     false,
				Score:         0,
				Price:         0.0,
				IsStarting:    false,
				PositionOrder: 2,
			},
			{
				PlayerID:      4,
				Captain:       false,
				ViceCaptain:   false,
				Injured:       false,
				Suspended:     false,
				Score:         0,
				Price:         0.0,
				IsStarting:    false,
				PositionOrder: 3,
			},
			{
				PlayerID:      5,
				Captain:       false,
				ViceCaptain:   false,
				Injured:       false,
				Suspended:     false,
				Score:         0,
				Price:         0.0,
				IsStarting:    false,
				PositionOrder: 4,
			},
			{
				PlayerID:      6,
				Captain:       false,
				ViceCaptain:   false,
				Injured:       false,
				Suspended:     false,
				Score:         0,
				Price:         0.0,
				IsStarting:    false,
				PositionOrder: 5,
			},
			{
				PlayerID:      7,
				Captain:       false,
				ViceCaptain:   false,
				Injured:       false,
				Suspended:     false,
				Score:         0,
				Price:         0.0,
				IsStarting:    false,
				PositionOrder: 6,
			},
			{
				PlayerID:      8,
				Captain:       false,
				ViceCaptain:   false,
				Injured:       false,
				Suspended:     false,
				Score:         0,
				Price:         0.0,
				IsStarting:    false,
				PositionOrder: 7,
			},
			{
				PlayerID:      9,
				Captain:       false,
				ViceCaptain:   false,
				Injured:       false,
				Suspended:     false,
				Score:         0,
				Price:         0.0,
				IsStarting:    false,
				PositionOrder: 8,
			},
			{
				PlayerID:      10,
				Captain:       false,
				ViceCaptain:   false,
				Injured:       false,
				Suspended:     false,
				Score:         0,
				Price:         0.0,
				IsStarting:    false,
				PositionOrder: 9,
			},
			{
				PlayerID:      11,
				Captain:       false,
				ViceCaptain:   false,
				Injured:       false,
				Suspended:     false,
				Score:         0,
				Price:         0.0,
				IsStarting:    false,
				PositionOrder: 10,
			},
			{
				PlayerID:      12,
				Captain:       false,
				ViceCaptain:   false,
				Injured:       false,
				Suspended:     false,
				Score:         0,
				Price:         0.0,
				IsStarting:    false,
				PositionOrder: 11,
			},
			{
				PlayerID:      13,
				Captain:       false,
				ViceCaptain:   false,
				Injured:       false,
				Suspended:     false,
				Score:         0,
				Price:         0.0,
				IsStarting:    false,
				PositionOrder: 12,
			},
			{
				PlayerID:      14,
				Captain:       false,
				ViceCaptain:   false,
				Injured:       false,
				Suspended:     false,
				Score:         0,
				Price:         0.0,
				IsStarting:    false,
				PositionOrder: 13,
			},
			{
				PlayerID:      15,
				Captain:       false,
				ViceCaptain:   false,
				Injured:       false,
				Suspended:     false,
				Score:         0,
				Price:         0.0,
				IsStarting:    false,
				PositionOrder: 14,
			},
		},
	}

	squadResult, err := squadService.CreateSquad(&squad, testutil.USER_ID)

	fmt.Println(squadResult)
	if err != nil {
		t.Fatalf("Failed to create squad: %v", err)
	}
	assert.NoError(t, err)
	assert.Equal(t, "Test Squad", squadResult.Name)
	assert.Equal(t, len(squadResult.Players), 15)
	for _, player := range squadResult.Players {
		assert.NotEmpty(t, player.ID)
		assert.GreaterOrEqual(t, player.PositionOrder, int32(0))
	}
}
