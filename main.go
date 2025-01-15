package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rushi/Desktop/ecom/configs"
	"github.com/rushi/Desktop/ecom/middleware"
	"github.com/rushi/Desktop/ecom/routes"
)

func main() {

	configs.ConnectDatabase()
	defer configs.CloseDatabase()

	r := mux.NewRouter()

	// Apply middleware
	r.Use(middleware.LogRequest)   // Log all requests
	r.Use(middleware.ErrorHandler) // Handle errors globally

	routes.SetupRoutes(r)

	fmt.Printf("Server is running on http://localhost%s\n", ":8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
