package models

import "time"

type League struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type GameWeekLeagueScore struct {
	ID           uint      `gorm:"primaryKey"`
	GameWeekID   uint      `gorm:"not null"`
	GameWeek     GameWeek  `gorm:"foreignKey:GameWeekID"`
	LeagueID     uint      `gorm:"not null"`
	League       League    `gorm:"foreignKey:LeagueID"`
	averageScore uint      `gorm:"not null"`
	highestScore uint      `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

type LeagueLeaderboard struct {
	ID        uint      `gorm:"primaryKey"`
	LeagueID  uint      `gorm:"not null"`
	League    League    `gorm:"foreignKey:LeagueID"`
	UserID    uint      `gorm:"not null"`
	User      User      `gorm:"foreignKey:UserID"`
	Rank      uint      `gorm:"not null"`
	Score     uint      `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (League) TableName() string {
	return "leagues"
}
