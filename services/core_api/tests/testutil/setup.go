package testutil

import (
	"context"
	"fmt"
	"os"
	"pkg/database"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/gorm"
)

const USER_ID = "a7d64e79-d36d-4f3d-9e3a-9ae27adecef1"

func SetupTestDB(t *testing.T) (*gorm.DB, func()) {
	ctx := context.Background()

	// Start PostgreSQL container
	pgContainer, err := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithDatabase("test_db"),
		postgres.WithUsername("test_user"),
		postgres.WithPassword("test_password"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(60*time.Second),
		),
	)
	if err != nil {
		t.Fatalf("Failed to start postgres container: %v", err)
	}

	// Get connection string
	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to get connection string: %v", err)
	}

	fmt.Println(connStr)

	// Connect with GORM
	db, err := database.NewConnection(connStr)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations
	err = database.AutoMigrate(db)
	if err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}

	err = loadSeedData(db, "/Users/mac/Desktop/Fantasy-Project/Fantasy-Engine/pkg/database/seed.sql")
	if err != nil {
		// t.Fatalf("Failed to load seed data: %v", err)
		fmt.Println("Failed to load seed data: %v", err)
		fmt.Println("Continuing with test...")
	}

	// Return cleanup function that terminates the container
	cleanup := func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Logf("Failed to terminate container: %v", err)
		}
	}
	return db, cleanup
}

// loadSeedData reads and executes a SQL file
func loadSeedData(db *gorm.DB, filepath string) error {
	// Read the SQL file
	sqlBytes, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("failed to read seed file %s: %w", filepath, err)
	}

	// Execute the SQL
	sql := string(sqlBytes)
	if err := db.Exec(sql).Error; err != nil {
		return fmt.Errorf("failed to execute seed SQL: %w", err)
	}

	return nil
}