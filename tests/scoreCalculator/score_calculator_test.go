package scorecalculator

import (
	"testing"

	"fantasy-engine/internal/models"
	"fantasy-engine/internal/services"

	"github.com/stretchr/testify/assert"
)


func TestMidfielderScoreCalculator(t *testing.T) {
	scoreCalculator := services.NewScoreCalculator()
	player := models.Player{
		ID:       1,
		Name:     "Midfielder Player",
		Position: "M",
	}

	stats := models.MatchStats{
		ID: 0,
		PlayerID: 1,
		Player: player,
		MatchID: 1,
		TotalPass: 29,
		AccuratePass: 21,
		TotalLongBalls: 5,
		AccurateLongBalls: 3,
		GoalAssist: 0,
		AccurateOwnHalfPasses: 15,
		TotalOwnHalfPasses: 10,
		AccurateOppositionHalfPasses: 11,
		TotalOppositionHalfPasses: 17,
		TotalCross: 1,
		AccurateCross: 0,
		AerialLost: 2,
		AerialWon: 1,
		DuelLost: 13,
		DuelWon: 6,
		ChallengeLost: 1,
		TotalContest: 3,
		WonContest: 1,
		OnTargetScoringAttempt: 1,
		Goals: 1,
		TotalClearance: 1,
		OutfielderBlock: 0,
		BallRecovery: 9,
		TotalTackle: 1,
		WonTackle: 0,
		ErrorLeadToAShot: 0,
		OwnGoals: 0,
		UnsuccessfulTouch: 1,
		WasFouled: 3,
		Fouls: 2,
		TotalOffside: 0,
		MinutesPlayed: 90,
		Touches: 46,
		Rating: 6,
		PossessionLostCtrl: 16,
		ExpectedGoals: 0.76,
		ExpectedAssists: 0,
		TotalShots: 2,
		Saves: 0,
		PenaltySave: 0,
		RedCard: false,
		YellowCard: false,			
	}

	score := scoreCalculator.CalculatePlayerScore(stats)

	var expectedScore int32 = 8
	assert.Equal(t, expectedScore, score)
}

func TestGoalKeeperScoreCalculator(t *testing.T) {
	scoreCalculator := services.NewScoreCalculator()
	player := models.Player{
		ID:       1,
		Name:     "GoalKeeper Player",
		Position: "G",
	}

	stats := models.MatchStats{
		ID: 0,
		PlayerID: 1,
		Player: player,
		MatchID: 1,
		TotalPass: 5,
		AccuratePass: 5,
		TotalLongBalls: 3,
		AccurateLongBalls: 3,
		GoalAssist: 0,
		AccurateOwnHalfPasses: 2,
		TotalOwnHalfPasses: 2,
		AccurateOppositionHalfPasses: 3,
		TotalOppositionHalfPasses: 3,
		TotalCross: 0,
		AccurateCross: 0,
		AerialLost: 0,
		AerialWon: 0,
		DuelLost: 1,
		DuelWon: 0,
		ChallengeLost: 1,
		TotalContest: 0,
		WonContest: 0,
		OnTargetScoringAttempt: 0,
		Goals: 0,
		TotalClearance: 0,
		OutfielderBlock: 0,
		BallRecovery: 0,
		TotalTackle: 0,
		WonTackle: 0,
		ErrorLeadToAShot: 0,
		OwnGoals: 0,
		UnsuccessfulTouch: 0,
		WasFouled: 0,
		Fouls: 1,
		TotalOffside: 0,
		MinutesPlayed: 90,
		Touches: 5,
		Rating: 6.1,
		PossessionLostCtrl: 0,
		ExpectedGoals: 0,
		ExpectedAssists: 0,
		TotalShots: 0,
		Saves: 3,
		PenaltySave: 0,
		RedCard: false,
		YellowCard: false,	
	}

	score := scoreCalculator.CalculatePlayerScore(stats)
	
	var expectedScore int32 = 3
	assert.Equal(t, expectedScore, score)
}

// func TestForwardPlayerCalculator(t *testing.T) {

// }

// func TestDefenderPlayerCalcualtor(t *testing.T) {

// }