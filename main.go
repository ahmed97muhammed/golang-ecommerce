package main

import (
	"golang_auth/database"
	"golang_auth/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize the database connection
	db.InitDB()

	// Initialize Gin router
	r := gin.Default()

	// Register routes
	routes.RegisterRoutes(r)

	// Start the server
	port := os.Getenv("PORT") // Get port from environment or default to 8080
	if port == "" {
		port = "8080" // Default port
	}
	r.Run(":" + port) // Start the server
}
