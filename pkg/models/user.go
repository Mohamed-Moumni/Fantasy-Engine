package models

import "time"

type User struct {
	ID               uint      `gorm:"primaryKey"`
	Email            string    `gorm:"not null;unique"`
	Username         string    `gorm:"not null;unique"`
	Age              uint      `gorm:"not null"`
	FavoriteTeamID   uint      `gorm:"not null"`
	FavoritePlayerID uint      `gorm:"not null"`
	SquadID          uint      `gorm:"not null"`
	Squad            Squad     `gorm:"foreignKey:SquadID"`
	CountryID        uint      `gorm:"not null"`
	Gender           string    `gorm:"not null;type:varchar(255);check:gender IN ('M', 'F')"`
	Password         string    `gorm:"not null"`
	CreatedAt        time.Time `gorm:"autoCreateTime"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime"`
}

func (User) TableName() string {
	return "users"
}
