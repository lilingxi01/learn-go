package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// User model
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"size:100;uniqueIndex;not null"`
	Age       int    `gorm:"check:age >= 0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {
	fmt.Println("=== GORM Introduction ===\n")

	// Connect to PostgreSQL
	dsn := "host=localhost user=postgres password=postgres dbname=tutorial port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}

	fmt.Println("✓ Connected to PostgreSQL with GORM\n")

	// Auto-migrate
	fmt.Println("Running auto-migration...")
	db.AutoMigrate(&User{})
	fmt.Println("✓ Migration complete\n")

	// Create
	fmt.Println("Creating users...")
	users := []User{
		{Name: "Alice", Email: "alice@example.com", Age: 30},
		{Name: "Bob", Email: "bob@example.com", Age: 25},
	}

	result := db.Create(&users)
	fmt.Printf("✓ Created %d users\n\n", result.RowsAffected)

	// Read
	fmt.Println("Reading users:")
	var allUsers []User
	db.Find(&allUsers)
	for _, u := range allUsers {
		fmt.Printf("  %d: %s (%s) - Age: %d\n", u.ID, u.Name, u.Email, u.Age)
	}

	// Cleanup
	fmt.Println("\nCleaning up...")
	db.Exec("DROP TABLE users")
	fmt.Println("✓ Complete")
}
