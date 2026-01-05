// internal/models/player.go
package models

import (
	"time"
)

type Player struct {
	ID                      uint        `gorm:"primaryKey"`
	TeamID                  uint        `gorm:"not null;column:teamId"`
	Team                    Team        `gorm:"foreignKey:TeamID"`
	Name                    string      `gorm:"not null;type:varchar(255)"`
	Slug                    string      `gorm:"not null;type:varchar(255)"`
	ShortName               string      `gorm:"not null;type:varchar(255);default:'G'"`
	Position                string      `gorm:"not null;type:varchar(255);check:position IN ('G', 'D', 'M', 'F')"`
	JerseyNumber            string      `gorm:"not null;type:varchar(255);column:jerseyNumber"`
	Height                  float32     `gorm:"not null"`
	Gender                  string      `gorm:"not null;type:varchar(255);check:gender IN ('M', 'F')"`
	CountryID               uint        `gorm:"not null;column:country"`
	Country                 Country     `gorm:"foreignKey:CountryID"`
	DateOfBirthTimestamp    time.Time   `gorm:"not null;type:timestamp;column:dateOfBirthTimestamp"`
	ProposedMarketValueRaw  uint        `gorm:"not null;column:proposedMarketValueRaw"`
	ShirtNumber             uint        `gorm:"not null;column:shirtNumber"`
	MatchStats              []MatchStats `gorm:"foreignKey:PlayerID"` // One-to-many relationship
	CreatedAt               time.Time   `gorm:"autoCreateTime"`
	UpdatedAt               time.Time   `gorm:"autoUpdateTime"`
}

func (Player) TableName() string {
	return "Player"
}