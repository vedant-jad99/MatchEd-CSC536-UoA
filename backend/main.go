package main

import (
	"log"
	"os"

	// "fmt"
	"backend/internal/controllers"
	"backend/internal/models"
	"backend/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// init gorm
	err = models.InitDB()
	if err != nil {
		log.Fatalf("Error connecting to database %v\n", err)
	}

	// assign handlers to API routes
	router := routes.SetupRouter()

<<<<<<< HEAD
	// initialize the tables
	//controllers.InitializeTables(models.DB)
=======
	controllers.InitializeTables()
>>>>>>> b85926a (added column tags, testing, and a few api calls)

	// Get port from environment variable or default to 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}

// Handler function placeholders - implement these in separate controller files
func loginHandler(c *gin.Context) {
	// Implementation
}

func registerHandler(c *gin.Context) {
	// Implementation
}

// ... implement other handler functions
