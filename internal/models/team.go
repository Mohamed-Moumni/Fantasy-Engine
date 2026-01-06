// internal/models/team.go
package models

import (
	"time"
)

type Team struct {
	ID           uint      `gorm:"primaryKey"`
	Name         string    `gorm:"not null"`
	Slug         string    `gorm:"not null"`
	ShortName    string    `gorm:"not null"`
	Gender       string    `gorm:"not null"`
	NameCode     string    `gorm:"not null"`
	CountryID    uint      `gorm:"not null"`
	Country      Country   `gorm:"foreignKey:countryID"`
	TeamColorsID uint      `gorm:"not null"`
	TeamColors   TeamColor `gorm:"foreignKey:teamColorsID"`
	Players      []Player  `gorm:"foreignKey:TeamID"` // One-to-many relationship
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

type TeamColor struct {
	ID        uint      `gorm:"primaryKey"`
	Primary   string    `gorm:"not null"`
	Secondary string    `gorm:"not null"`
	TextColor string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}