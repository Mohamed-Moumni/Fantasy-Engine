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
	case "F":
		return sc.calculateForwardScore(stats)
	case "D":
		return sc.calculateDefenderScore(stats)
	default:
		return 0
	}
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


	// clearness + OutfielderBlcok +  Ball recoveries + interceptions + tackles
	accumulationBonus := stats.OutfielderBlock + stats.TotalClearance + stats.BallRecovery + stats.InterceptionWon + stats.TotalTackle
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

func (sc *ScoreCalculator) calculateForwardScore(stats models.MatchStats) int32 {
	var score int32 = 0

	if stats.MinutesPlayed <= 60 {
		score += 1
	} else {
		score += 2
	}

	score += int32(stats.Goals * 4)
	score += int32(stats.GoalAssist * 3)

	// clearness + OutfielderBlcok +  Ball recoveries + interceptions + tackles
	accumulationBonus := stats.OutfielderBlock + stats.TotalClearance + stats.BallRecovery + stats.InterceptionWon + stats.TotalTackle
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

func (sc *ScoreCalculator) calculateDefenderScore(stats models.MatchStats) int32 {
	var score int32 = 0
	
	if stats.MinutesPlayed <= 60 {
		score += 1
	} else {
		score += 2
	}
	score += int32(stats.Goals * 7)
	score += int32(stats.GoalAssist * 3)

	// clearness + OutfielderBlock +  Ball recoveries + interceptions + tackles
	accumulationBonus := stats.OutfielderBlock + stats.TotalClearance + stats.BallRecovery + stats.InterceptionWon + stats.TotalTackle
	if accumulationBonus >= 10 {
		score += 2
	}

	if stats.Rating >= 8.0 {
		score += 2
	}

	if stats.WonTackle > 5 {
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


// // apply the characteristics of the team on the player's score
// func (sc *ScoreCalculator) calculateTeamLevelScore(squad ) int32 {
// }
