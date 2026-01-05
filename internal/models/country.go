package models

import "time"

type Country struct {
	ID        uint      `gorm:"primaryKey"`
	alpha2    string    `gorm:"not null"`
	alhpa3    string    `gorm:"not null"`
	name      string    `gorm:"not null"`
	slug      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
