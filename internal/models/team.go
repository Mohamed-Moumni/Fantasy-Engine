// internal/models/team.go
package models

import (
	"time"
)

type Team struct {
	ID           uint      `gorm:"primaryKey"`
	name         string    `gorm:"not null"`
	slug         string    `gorm:"not null"`
	shortName    string    `gorm:"not null"`
	countryID    uint      `gorm:"not null"`
	country      Country   `gorm:"foreignKey:countryID"`
	fullName     string    `gorm:"not null"`
	nameCode     string    `gorm:"not null"`
	teamColorsID uint      `gorm:"not null"`
	teamColors   TeamColor `gorm:"foreignKey:teamColorsID"`
	Players      []Player  `gorm:"foreignKey:TeamID"` // One-to-many relationship
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

type TeamColor struct {
	ID        uint      `gorm:"primaryKey"`
	primary   string    `gorm:"not null"`
	secondary string    `gorm:"not null"`
	textColor string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}