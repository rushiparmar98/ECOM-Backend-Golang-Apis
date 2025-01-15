package configs

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDatabase() {

	var err error

	connStr := "user=postgres dbname=ecom_db password=jamesbond007 host=localhost port=5432 sslmode=disable"

	DB, err = sql.Open("postgres", connStr)
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
