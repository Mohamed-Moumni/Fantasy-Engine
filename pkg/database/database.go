package database

import (
	"errors"
	"fmt"
	"log"
	"pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewConnection(databaseURL string) (*gorm.DB, error) {
	if databaseURL == "" {
		return nil, errors.New("database URL is required")
	}
	
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("Failed to get SqlDB: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("Failed to ping database: %w", err)
	}

	log.Println("Successfully connected to the database")
	return db, nil
}

// AutoMigrate runs database migrations using GORM
// This creates tables based on your model definitions
func AutoMigrate(db *gorm.DB) error {
	log.Println("🔄 Running database migrations...")

	err := db.AutoMigrate(
		&models.Country{},
		&models.TeamColor{},
		&models.Team{},
		&models.Player{},
		&models.GameWeek{},
		&models.Match{},
		&models.MatchStats{},
	)

	if err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	log.Println("✅ Database migrations completed successfully")
	return nil
}
