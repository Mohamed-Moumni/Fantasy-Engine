package models

import "time"

type Squad struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type GameWeekSquadScore struct {
	ID         uint      `gorm:"primaryKey"`
	GameWeekID uint      `gorm:"not null"`
	GameWeek   GameWeek  `gorm:"foreignKey:GameWeekID"`
	SquadID    uint      `gorm:"not null"`
	Squad      Squad     `gorm:"foreignKey:SquadID"`
	Score      uint      `gorm:"not null"`
	LeagueID   uint      `gorm:"not null"`
	League     League    `gorm:"foreignKey:LeagueID"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

type SquadPlayer struct {
	ID            uint      `gorm:"primaryKey"`
	SquadID       uint      `gorm:"not null"`
	Squad         Squad     `gorm:"foreignKey:SquadID"`
	PlayerID      uint      `gorm:"not null"`
	Player        Player    `gorm:"foreignKey:PlayerID"`
	Captain       bool      `gorm:"not null;default:false"`
	ViceCaptain   bool      `gorm:"not null;default:false"`
	Injured       bool      `gorm:"not null;default:false"`
	Suspended     bool      `gorm:"not null;default:false"`
	Score         uint      `gorm:"not null"`
	Price         float32   `gorm:"not null"`
	IsStarting    bool      `gorm:"not null;default:false"`
	PositionOrder uint      `gorm:"not null"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}

func (GameWeekSquadScore) TableName() string {
	return "game_week_squad_scores"
}

func (Squad) TableName() string {
	return "squads"
}
