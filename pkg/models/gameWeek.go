// internal/models/gameWeek.go
package models

import (
	"time"
)

type GameWeek struct {
	ID        uint    `gorm:"primaryKey"`
	Round     uint    `gorm:"not null"`
	Matches   []Match `gorm:"foreignKey:GameWeekID"` // One-to-many relationship
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (GameWeek) TableName() string {
	return "GameWeek"
}