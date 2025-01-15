package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rushi/Desktop/ecom/configs"
	"github.com/rushi/Desktop/ecom/middleware"
	"github.com/rushi/Desktop/ecom/routes"
)

func main() {
	// Get the port from environment variables or default to 8080
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the port from environment variables or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	// Get the database URL from environment variables
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	// Connect to the database using the database URL
	configs.ConnectDatabase() // This should now work as the function is defined correctly
	defer configs.CloseDatabase()

	// Create a new router
	r := mux.NewRouter()

	// Apply middleware
	r.Use(middleware.LogRequest)   // Log all requests
	r.Use(middleware.ErrorHandler) // Handle errors globally

	// Setup routes
	routes.SetupRoutes(r)

	// Start the server
	serverAddr := fmt.Sprintf(":%s", port)
	fmt.Printf("Server is running on http://localhost%s\n", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, r))
}
