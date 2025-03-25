# Go/Gin Backend Setup Instructions

This document provides guidance on setting up a Go/Gin backend that will work with the matchEd React frontend.

## Prerequisites

1. Install Go (version 1.16 or later recommended)
2. Basic knowledge of Go programming language
3. Understanding of RESTful API design

## Project Structure

Create a new directory for your Go backend, separate from your React frontend:

```
backend/
├── main.go              # Entry point for your Gin application
├── controllers/         # Handler functions for routes
├── models/              # Data models
├── middleware/          # Middleware (auth, logging, etc.)
├── routes/              # Route definitions
├── config/              # Configuration
├── utils/               # Utility functions
└── static/              # Will store the built React frontend for production
```

## Setting Up the Go/Gin Server

1. Initialize a new Go module:

```bash
mkdir -p backend
cd backend
go mod init github.com/yourusername/matched-backend
```

2. Install Gin and other dependencies:

```bash
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm             # If you need an ORM
go get -u gorm.io/driver/postgres  # Or other database driver
go get -u github.com/dgrijalva/jwt-go  # For JWT authentication
```

3. Create the main.go file:

```go
package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	// Set the mode (debug/release)
	gin.SetMode(gin.ReleaseMode)
	if os.Getenv("DEBUG") == "true" {
		gin.SetMode(gin.DebugMode)
	}

	// Initialize the router
	r := gin.Default()

	// Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// API routes
	api := r.Group("/api")
	{
		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/login", loginHandler)
			auth.POST("/register", registerHandler)
			auth.POST("/logout", logoutHandler)
		}

		// User routes
		users := api.Group("/users")
		{
			users.GET("/", getAllUsersHandler)
			users.GET("/:id", getUserByIdHandler)
			users.POST("/", createUserHandler)
			users.PUT("/:id", updateUserHandler)
			users.DELETE("/:id", deleteUserHandler)
		}

		// Course routes
		courses := api.Group("/courses")
		{
			courses.GET("/", getAllCoursesHandler)
			courses.GET("/:id", getCourseByIdHandler)
			courses.POST("/", createCourseHandler)
			courses.PUT("/:id", updateCourseHandler)
			courses.DELETE("/:id", deleteCourseHandler)
		}

		// People routes
		people := api.Group("/people")
		{
			people.GET("/", getAllPeopleHandler)
			people.GET("/:id", getPersonByIdHandler)
			people.POST("/", createPersonHandler)
			people.PUT("/:id", updatePersonHandler)
			people.DELETE("/:id", deletePersonHandler)
		}

		// Roles routes
		roles := api.Group("/roles")
		{
			roles.GET("/", getAllRolesHandler)
			roles.POST("/assign", assignPersonToRoleHandler)
			roles.POST("/unassign", unassignPersonFromRoleHandler)
		}
	}

	// Serve static files from the React app build directory in production
	r.Use(static.Serve("/", static.LocalFile("./static", false)))

	// Handle all routes for SPA (forward to index.html)
	// This should be after all API routes
	r.NoRoute(func(c *gin.Context) {
		// Check if the request path is an API route
		if c.Request.URL.Path[:4] == "/api" {
			c.JSON(http.StatusNotFound, gin.H{"error": "API endpoint not found"})
			return
		}
		
		// Otherwise, serve the SPA
		c.File("./static/index.html")
	})

	// Get port from environment variable or default to 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Start the server
	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
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
```

## Implementing Controllers

Create controller files for each resource. Example for people controller:

```go
// controllers/people.go
package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"your-project/models"
)

// GetAllPeople handles GET requests to fetch all people
func GetAllPeople(c *gin.Context) {
	var people []models.Person
	
	// Fetch people from database
	// db.Find(&people)
	
	c.JSON(http.StatusOK, people)
}

// GetPerson handles GET requests to fetch a single person
func GetPerson(c *gin.Context) {
	id := c.Param("id")
	var person models.Person
	
	// Fetch person from database
	// result := db.First(&person, id)
	// if result.Error != nil {
	//    c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
	//    return
	// }
	
	c.JSON(http.StatusOK, person)
}

// CreatePerson handles POST requests to create a new person
func CreatePerson(c *gin.Context) {
	var newPerson models.Person
	
	if err := c.ShouldBindJSON(&newPerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// Save to database
	// db.Create(&newPerson)
	
	c.JSON(http.StatusCreated, newPerson)
}

// ... implement other CRUD operations
```

## Building and Deploying

1. Build your React frontend:

```bash
cd /path/to/frontend
npm run build
```

2. Copy the build output to your Go project's static directory:

```bash
cp -r dist/* /path/to/backend/static/
```

3. Build and run your Go application:

```bash
cd /path/to/backend
go build -o app
./app
```

## Testing the API

You can test your API endpoints using curl or Postman before deployment:

```bash
# Example: Get all people
curl http://localhost:3000/api/people

# Example: Create a new person
curl -X POST http://localhost:3000/api/people \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","details":"New faculty member","type":"human","status":"active"}'
```

## Deployment Considerations

1. Use environment variables for configuration in production
2. Set up proper database connections
3. Implement authentication and authorization
4. Consider containerization with Docker
5. Set up proper CORS handling if needed
