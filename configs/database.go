package configs

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

// ConnectDatabase connects to the PostgreSQL database using the connection string
// from the environment variable or a hardcoded local connection string.
func ConnectDatabase() {

	var err error

	// Get the database URL from the environment variable
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		// If DATABASE_URL is not set, use the hardcoded connection string (local development)
		dbURL = "user=postgres dbname=ecom_db password=jamesbond007 host=localhost port=5432 sslmode=disable"
	}

	// Open a connection to the database
	DB, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Test the database connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database", err)
	}
	fmt.Println("Connected to the database successfully!")
}

// CloseDatabase closes the database connection.
func CloseDatabase() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			log.Printf("Error closing the database connection: %v", err)
		} else {
			fmt.Println("Database connection closed")
		}
	}
}
