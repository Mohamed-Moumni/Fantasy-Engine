package services

import (
	"fantasy-engine/internal/models"
)

type ScoreCalculator struct {
}

func NewScoreCalculator() *ScoreCalculator {
	return &ScoreCalculator{}
}

func (sc *ScoreCalculator) CalculatePlayerScore(stats models.MatchStats) int32 {
	switch stats.Player.Position {
	case "M":
		return sc.calculateMidfielderScore(stats)
	case "G":
		return sc.calcualtorGoalKeeperScore(stats)
	// Add other positions like "F", "D", "G" here
	default:
		return 0
	}
	return 0
}

func (sc *ScoreCalculator) calculateMidfielderScore(stats models.MatchStats) int32 {
	var score int32 = 0

	if stats.MinutesPlayed <= 60 {
		score += 1
	} else {
		score += 2
	}

	score += int32(stats.Goals * 6)
	score += int32(stats.GoalAssist * 3)

	// clearness + Ball recoveries + interceptions + tackles
	accumulationBonus := stats.TotalClearance + stats.BallRecovery + stats.InterceptionWon + stats.TotalTackle
	if accumulationBonus >= 12 {
		score += 2
	}

	if stats.YellowCard {
		score -= 1
	}

	if stats.RedCard {
		score -= 3
	}

	score -= int32(stats.OwnGoals * 2)
	return score
}

func (sc *ScoreCalculator) calcualtorGoalKeeperScore(stats models.MatchStats) int32 {
	var score int32 = 0

	if stats.MinutesPlayed <= 60 {
		score += 1
	} else {
		score += 2
	}

	score += int32(stats.Goals * 15)
	score += int32(stats.GoalAssist * 3)

	score += int32(stats.Saves) / 3

	score += int32(stats.PenaltySave * 5)

	if stats.YellowCard {
		score -= 1
	}

	if stats.RedCard {
		score -= 3
	}

	score -= int32(stats.OwnGoals * 2)
	return score
}


// func (sc *ScoreCalculator) calculateDefenderScore(stats models.MatchStats) int32 {

// }

// func (sc *ScoreCalculator) calculateForwardScore(stats models.MatchStats) int32 {

// }