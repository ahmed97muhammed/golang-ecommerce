package db

import (
	"fmt"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
	"os"
	"golang_auth/models"
)

// Global variable for the DB connection
var DB *gorm.DB

// Initialize the database connection
func InitDB() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Retrieve the MySQL connection details from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbCharset := os.Getenv("DB_CHARSET")

	// Format the DSN string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbCharset)

	// Establish the connection
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// List of models to migrate
	modelsList := []interface{}{
        &models.Product{},
        &models.User{},      
        &models.Order{},      
    }

    // Migrate each model
    for _, model := range modelsList {
        err := DB.AutoMigrate(model)
        if err != nil {
            fmt.Printf("Error migrating %T: %v\n", model, err)
        } else {
            fmt.Printf("Successfully migrated %T\n", model)
        }
    }
	
	fmt.Println("Database connection established successfully")
}
