package service

import (
	"context"
	"core_api/graph/model"
	"core_api/internal/repository"
	"fmt"
	"pkg/models"
)

type SquadService struct {
	squadRepository *repository.SquadRepository
	userRepository  *repository.UserRepository
	squadValidator  *SquadValidator
}

func NewSquadService(userRepository *repository.UserRepository, squadRepository *repository.SquadRepository) *SquadService {
	return &SquadService{userRepository: userRepository, squadRepository: squadRepository, squadValidator: NewSquadValidator(100.00, 15)}
}

func (ss *SquadService) createSquadWithPlayers(ctx context.Context, squad *model.CreateSquadInput) (*model.Squad, error) {
	userId, ok := ctx.Value("userID").(string)
	if !ok {
		return nil, fmt.Errorf("User Unauthorized")
	}
	squadModel := models.Squad{Name: squad.Name}
	err := ss.squadRepository.CreateSquad(&squadModel, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to create squad: %w", err)
	}

	// Build the response players slice
	squadPlayers := make([]*model.SquadPlayer, 0, len(squad.Players))

	for _, player := range squad.Players {
		squadPlayer := models.SquadPlayer{
			SquadID:       squadModel.ID,
			PlayerID:      uint(player.PlayerID),
			Captain:       player.Captain,
			ViceCaptain:   player.ViceCaptain,
			Injured:       player.Injured,
			Suspended:     player.Suspended,
			Score:         uint(player.Score),
			Price:         float32(player.Price),
			IsStarting:    player.IsStarting,
			PositionOrder: uint(player.PositionOrder),
		}
		err := ss.squadRepository.CreateSquadPlayer(&squadPlayer)
		if err != nil {
			return nil, fmt.Errorf("failed to create squad player: %w", err)
		}

		// Add to response slice
		squadPlayers = append(squadPlayers, &model.SquadPlayer{
			ID:            int32(squadPlayer.ID),
			Captain:       squadPlayer.Captain,
			ViceCaptain:   squadPlayer.ViceCaptain,
			Injured:       squadPlayer.Injured,
			Suspended:     squadPlayer.Suspended,
			Score:         int32(squadPlayer.Score),
			Price:         float64(squadPlayer.Price),
			IsStarting:    squadPlayer.IsStarting,
			PositionOrder: int32(squadPlayer.PositionOrder),
			// Player: nil - GraphQL will resolve this via a field resolver if needed
		})
	}

	return &model.Squad{
		ID:      int32(squadModel.ID),
		Name:    squadModel.Name,
		Players: squadPlayers,
	}, nil
}

func (ss *SquadService) CreateSquad(ctx context.Context, squad *model.CreateSquadInput) (*model.Squad, error) {
	userId, ok := ctx.Value("userID").(string)
	if !ok {
		return nil, fmt.Errorf("User Unauthorized")
	}

	user, err := ss.userRepository.GetUserByID(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}
	if user.SquadID != nil {
		return nil, ErrSquadAlreadyExists // Fix: use the error directly, not fmt.Errorf with two args
	}
	return ss.createSquadWithPlayers(ctx, squad)
}

func (ss *SquadService) SubstitutePlayer(ctx context.Context, input *model.SubstitutePlayerInput) (*model.Squad, error) {

	// check if the user has a squad
	// check if the player to substitute and the player to replace are in the same squad
	// implement the logic to substitute the player

	// user, err := ss.userRepository.GetUserByID(userId)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to get user by id: %w", err)
	// }
	// if user.SquadID == nil {
	// 	return nil, ErrSquadNotFound
	// }

	// squadPlayers, err := ss.squadRepository.GetSquadPlayersBySquadID(int32(*user.SquadID))
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to get squad players by squad id: %w", err)
	// }

	// squadPlayerToSubstitute := models.SquadPlayer{}
	// squadPlayerToReplace := models.SquadPlayer{}

	// for _, squadPlayer := range squadPlayers {
	// 	if squadPlayer.ID == uint(input.SquadPlayerIDToSubstitute) {
	// 		squadPlayerToSubstitute = squadPlayer
	// 	}
	// 	if squadPlayer.ID == uint(input.SquadPlayerIDToSubstitute) {
	// 		squadPlayerToReplace = squadPlayer
	// 	}
	// }

	// if squadPlayerToSubstitute.ID == 0 || squadPlayerToReplace.ID == 0 {
	// 	return nil, ErrSquadPlayerNotFound
	// }
	return nil, nil
}