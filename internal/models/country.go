package models

import "time"

type Country struct {
	ID        uint      `gorm:"primaryKey"`
	Alpha2    string    `gorm:"not null"`
	Alpha3    string    `gorm:"not null"`
	Name      string    `gorm:"not null;unique"`
	Slug      string    `gorm:"not null;unique"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
