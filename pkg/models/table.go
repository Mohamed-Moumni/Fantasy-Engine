package models

import "time"

type Table struct {
	ID                 uint      `gorm:"primaryKey"`
	TeamID             uint      `gorm:"not null;column:team_id"`
	Team               Team      `gorm:"foreinKey:TeamID;contraint:OnDelete:CASCADE" json:"team"`
	Position           int       `gorm:"not null" json:"position"`
	Played             int       `gorm:"not null;default:0" json:"played"`
	Wins               int       `gorm:"not null;default:0" json:"wins"`
	Losses             int       `gorm:"not null;default:0" json:"losses"`
	ScoresFor          int       `gorm:"not null;default:0" json:"scoresFor"`
	ScoresAgainst      int       `gorm:"not null;default:0" json:"scoresAgainst"`
	Draws              int       `gorm:"not null;default:0" json:"draws"`
	Points             int       `gorm:"not null;default:0" json:"points"`
	ScoreDiffFormatted int       `gorm:"not null;default:0" json:"scoreDiffFormatted"`
	CreatedAt          time.Time `gorm:"autoCreateTime"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime"`
}

func (Table) TableName() string {
	return "tables"
}
