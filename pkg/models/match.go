// internal/models/match_stats.go
package models

import (
	"time"

	"gorm.io/datatypes"
)

// internal/models/match.go
type Match struct {
	ID         uint           `gorm:"primaryKey"`
	GameWeekID uint           `gorm:"not null;column:gameweek"`
	GameWeek   GameWeek       `gorm:"foreignKey:GameWeekID"`
	HomeTeamID uint           `gorm:"not null;column:home_team_id"`
	HomeTeam   Team           `gorm:"foreignKey:HomeTeamID"`
	AwayTeamID uint           `gorm:"not null;column:away_team_id"`
	AwayTeam   Team           `gorm:"foreignKey:AwayTeamID"`
	HomeScore  uint           `gorm:"not null;default:0"`
	AwayScore  uint           `gorm:"not null;default:0"`
	Date       datatypes.Date `gorm:"not null;type:date"`
	Time       datatypes.Time `gorm:"not null;type:time"`
	Completed  bool           `gorm:"not null;default:false"`
	MatchStats []MatchStats   `gorm:"foreignKey:MatchID"` // One-to-many relationship
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
}

func (Match) TableName() string {
	return "Match"
}

type MatchStats struct {
	ID                           uint      `gorm:"primaryKey"`
	PlayerID                     uint      `gorm:"not null;column:player"`
	Player                       Player    `gorm:"foreignKey:PlayerID"`
	MatchID                      uint      `gorm:"not null;column:match"`
	Match                        Match     `gorm:"foreignKey:MatchID"`
	TotalPass                    uint16    `gorm:"not null;column:totalPass"`
	AccuratePass                 uint16    `gorm:"not null;column:accuratePass"`
	TotalLongBalls               uint16    `gorm:"not null;column:totalLongBalls"`
	AccurateLongBalls            uint16    `gorm:"not null;column:accurateLongBalls"`
	GoalAssist                   uint16    `gorm:"not null;column:goalAssist"`
	AccurateOwnHalfPasses        uint16    `gorm:"not null;column:accurateOwnHalfPasses"`
	TotalOwnHalfPasses           uint16    `gorm:"not null;column:totalOwnHalfPasses"`
	AccurateOppositionHalfPasses uint16    `gorm:"not null;column:accurateOppositionHalfPasses"`
	TotalOppositionHalfPasses    uint16    `gorm:"not null;column:totalOppositionHalfPasses"`
	TotalCross                   uint16    `gorm:"not null;column:totalCross"`
	AccurateCross                uint16    `gorm:"not null;column:accurateCross"`
	AerialLost                   uint16    `gorm:"not null;column:aerialLost"`
	AerialWon                    uint16    `gorm:"not null;column:aerialWon"`
	DuelLost                     uint16    `gorm:"not null;column:duelLost"`
	DuelWon                      uint16    `gorm:"not null;column:duelWon"`
	ChallengeLost                uint16    `gorm:"not null;column:challengeLost"`
	TotalContest                 uint16    `gorm:"not null;column:totalContest"`
	WonContest                   uint16    `gorm:"not null;column:wonContest"`
	OnTargetScoringAttempt       uint16    `gorm:"not null;column:onTargetScoringAttempt"`
	Goals                        uint16    `gorm:"not null;column:goals"`
	TotalClearance               uint16    `gorm:"not null;column:totalClearance"`
	OutfielderBlock              uint16    `gorm:"not null;column:outfielderBlock"`
	BallRecovery                 uint16    `gorm:"not null;column:ballRecovery"`
	TotalTackle                  uint16    `gorm:"not null;column:totalTackle"`
	WonTackle                    uint16    `gorm:"not null;column:wonTackle"`
	ErrorLeadToAShot             uint16    `gorm:"not null;column:errorLeadToAShot"`
	OwnGoals                     uint16    `gorm:"not null;column:ownGoals"`
	UnsuccessfulTouch            uint16    `gorm:"not null;column:unsuccessfulTouch"`
	WasFouled                    uint16    `gorm:"not null;column:wasFouled"`
	Fouls                        uint16    `gorm:"not null;column:fouls"`
	TotalOffside                 uint16    `gorm:"not null;column:totalOffside"`
	MinutesPlayed                uint16    `gorm:"not null;column:minutesPlayed"`
	Touches                      uint16    `gorm:"not null;column:touches"`
	Rating                       float32   `gorm:"not null;column:rating"`
	PossessionLostCtrl           uint16    `gorm:"not null;column:possessionLostCtrl"`
	ExpectedGoals                float32   `gorm:"not null;column:expectedGoals"`
	ExpectedAssists              uint16    `gorm:"not null;column:expectedAssists"`
	TotalShots                   uint16    `gorm:"not null;column:totalShots"`
	InterceptionWon              uint16    `gorm:"not null;column:interceptionWon"`
	Saves                        uint16    `gorm:"not null;column:saves"`
	PenaltySave                  uint16    `gorm:"not null;column:penaltySave"`
	RedCard                      bool      `gorm:"not null;default:false;column:redCard"`
	YellowCard                   bool      `gorm:"not null;default:false;column:yellowCard"`
	CreatedAt                    time.Time `gorm:"autoCreateTime"`
	UpdatedAt                    time.Time `gorm:"autoUpdateTime"`
}

func (MatchStats) TableName() string {
	return "MatchStats"
}
