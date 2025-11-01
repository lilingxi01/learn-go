package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	// ConnectionString for local PostgreSQL
	// Change these values based on your setup
	host     = "localhost"
	port     = 5432 // Use 54322 for Supabase local
	user     = "postgres"
	password = "postgres"
	dbname   = "tutorial"
)

func main() {
	fmt.Println("=== PostgreSQL Connection Test ===\n")

	// Build connection string
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Connect to PostgreSQL
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	defer db.Close()

	// Test connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	fmt.Println("✓ Successfully connected to PostgreSQL!")

	// Get PostgreSQL version
	var version string
	err = db.QueryRow("SELECT version()").Scan(&version)
	if err != nil {
		log.Fatal("Failed to get version:", err)
	}

	fmt.Printf("\nPostgreSQL Version:\n%s\n", version)

	// Create a test table
	fmt.Println("\nCreating test table...")
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS test_connection (
		id SERIAL PRIMARY KEY,
		message TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
	fmt.Println("✓ Table created successfully")

	// Insert test data
	fmt.Println("\nInserting test data...")
	_, err = db.Exec("INSERT INTO test_connection (message) VALUES ($1)", "Hello from Go!")
	if err != nil {
		log.Fatal("Failed to insert:", err)
	}
	fmt.Println("✓ Data inserted successfully")

	// Query data
	fmt.Println("\nQuerying data:")
	rows, err := db.Query("SELECT id, message, created_at FROM test_connection")
	if err != nil {
		log.Fatal("Failed to query:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var message string
		var createdAt string

		err = rows.Scan(&id, &message, &createdAt)
		if err != nil {
			log.Fatal("Failed to scan:", err)
		}

		fmt.Printf("  ID: %d, Message: %s, Created: %s\n", id, message, createdAt)
	}

	// Cleanup
	fmt.Println("\nCleaning up...")
	_, err = db.Exec("DROP TABLE test_connection")
	if err != nil {
		log.Fatal("Failed to drop table:", err)
	}
	fmt.Println("✓ Table dropped successfully")

	fmt.Println("\n=== Connection Test Complete! ===")
}
