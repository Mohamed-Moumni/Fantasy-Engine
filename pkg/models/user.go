package models

import "time"

type User struct {
	ID               string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Email            string    `gorm:"not null;unique"`
	Username         string    `gorm:"not null;unique"`
	Age              uint      `gorm:"not null"`
	FavoriteTeamID   uint      `gorm:"default:null"`
	FavoritePlayerID uint      `gorm:"default:null"`
	SquadID          *uint     `gorm:"default:null"`
	Squad            *Squad    `gorm:"foreignKey:SquadID"`
	CountryID        uint      `gorm:"default:null"`
	Gender           string    `gorm:"not null;type:varchar(255);check:gender IN ('M', 'F')"`
	PhoneNumber      string    `gorm:"size:20;not null;uniqueIndex"`
	PasswordHash     string    `gorm:"default:null"`
	PasswordActived  bool      `gorm:"not null;default:false"`
	CreatedAt        time.Time `gorm:"autoCreateTime"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime"`
}

func (User) TableName() string {
	return "users"
}
