package service

import (
	"core_api/graph/model"
	"fmt"
)

type SquadValidator struct {
	maxBudget  float64
	maxPlayers int
}

func NewSquadValidator(maxBudget float64, maxPlayers int) *SquadValidator {
	return &SquadValidator{maxBudget: maxBudget, maxPlayers: maxPlayers}
}

func (sv *SquadValidator) Validate(squad *model.CreateSquadInput) error {
	if err := sv.validatePlayerCount(squad.Players); err != nil {
		return err
	}
	if err := sv.validateCaptaincy(squad.Players); err != nil {
		return err
	}
	if err := sv.validateNoDuplicatePlayers(squad.Players); err != nil {
		return err
	}
	if err := sv.validateBudget(squad.Players); err != nil {
		return err
	}
	return nil
}

// func (sv *SquadValidator) validateComposition(squad *model.CreateSquadInput) error {
// 	goalkeeperCount := 0
// 	defenderCount := 0
// 	midfielderCount := 0
// 	forwardCount := 0

// 	for _, player := range squad.Players {
// 		switch player.Player.Position {
// 		case "G":
// 			goalkeeperCount++
// 		case "D":
// 			defenderCount++
// 		case "M":
// 			midfielderCount++
// 		case "F":
// 			forwardCount++
// 		}
// 	}
// 	if goalkeeperCount != 2 {
// 		return fmt.Errorf("invalid goalkeeper count: %w", ErrInvalidGoalkeeperCount)
// 	}
// 	if defenderCount != 5 {
// 		return fmt.Errorf("invalid defender count: %w", ErrInvalidDefenderCount)
// 	}
// 	if midfielderCount != 5 {
// 		return fmt.Errorf("invalid midfielder count: %w", ErrInvalidMidfielderCount)
// 	}
// 	if forwardCount != 3 {
// 		return fmt.Errorf("invalid forward count: %w", ErrInvalidForwardCount)
// 	}
// 	return nil
// }

func (sv *SquadValidator) validatePlayerCount(players []*model.CreateSquadPlayerInput) error {
	if len(players) != sv.maxPlayers {
		return fmt.Errorf("invalid player count: %w", ErrInvalidPlayerCount)
	}
	return nil
}

func (sv *SquadValidator) validateCaptaincy(players []*model.CreateSquadPlayerInput) error {
	captainCount := 0
	viceCaptainCount := 0

	for _, player := range players {
		if player.Captain && player.ViceCaptain {
			return fmt.Errorf("captain and vice captain cannot be the same player: %w", ErrCaptainIsViceCaptain)
		}
		if player.Captain {
			captainCount++
		}
		if player.ViceCaptain {
			viceCaptainCount++
		}
	}
	if captainCount != 1 {
		return fmt.Errorf("no captain: %w", ErrNoCaptain)
	}
	if viceCaptainCount != 1 {
		return fmt.Errorf("no vice captain: %w", ErrNoViceCaptain)
	}
	return nil
}

func (sv *SquadValidator) validateNoDuplicatePlayers(players []*model.CreateSquadPlayerInput) error {
	playerMap := make(map[int]bool)

	for _, player := range players {
		if playerMap[int(player.PlayerID)] {
			return fmt.Errorf("duplicate player: %w", ErrDuplicatePlayer)
		}
		playerMap[int(player.PlayerID)] = true
	}
	return nil
}

func (sv *SquadValidator) validateBudget(players []*model.CreateSquadPlayerInput) error {
	totalPrice := 0.0
	for _, player := range players {
		totalPrice += player.Price
	}
	if totalPrice > sv.maxBudget {
		return fmt.Errorf("budget exceeded: %w", ErrBudgetExceeded)
	}
	return nil
}
