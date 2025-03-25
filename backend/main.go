package main

import (
	"log"
	"os"

	"backend/models"
	"backend/routes"

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
	models.InitDB()
	// assign handlers to API routes
	router := routes.SetupRouter()

	// Get port from environment variable or default to 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
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
