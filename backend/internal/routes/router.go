package routes

import (
	"net/http"
	"os"
	"time"

	// Import router dependencies
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"backend/internal/controllers"
)

func SetupRouter() *gin.Engine {
	// Set the mode (debug/release)
	gin.SetMode(gin.ReleaseMode)
	if os.Getenv("DEBUG") == "true" {
		gin.SetMode(gin.DebugMode)
	}

	// Initialize the router
	r := gin.Default()

	// cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"}, // or "*" for all origins
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// Configure CORS middleware
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, Accept, User-Agent, Cache-Control, Pragma, Sec-Fetch-Dest, Sec-Fetch-Mode, Sec-Fetch-Site, Accept-Encoding, Accept-Language, Content-Length")
		c.Header("Access-Control-Expose-Headers", "Content-Length")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Max-Age", "43200")

		// Handle OPTIONS preflight
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
	// "/api" routes
	setApiHandlers(r)

	// Serve static files from the React app build directory in production
	// r.Use(static.Serve("/", static.LocalFile("../frontend/dist", false)))

	// alternative frontend
	r.Use(static.Serve("/", static.LocalFile("../match-view/dist", false)))

	// Handle all routes for SPA (forward to index.html)
	// This should be after all API routes
	r.NoRoute(func(c *gin.Context) {
		// Check if the request path is an API route
		if c.Request.URL.Path[:4] == "/api" {
			c.JSON(http.StatusNotFound, gin.H{"error": "API endpoint not found"})
			return
		}

		// Otherwise, serve the SPA
		// c.File("./static/index.html")
	})

	return r
}

// maps the handlers responsible for
// GET, POST, PUT, DELETE
// for the api group "/api"
func setApiHandlers(r *gin.Engine) {
	api := r.Group("/api")
	{
		// User routes
		users := api.Group("/users")
		{
			users.GET("/", controllers.GetAllUsers)
			users.GET("/:id", controllers.GetUserById)
			users.POST("/", controllers.CreateUser)
			users.PUT("/:id", controllers.UpdateUser)
			users.DELETE("/:id", controllers.DeleteUser)
		}

		// ping the server
		// TODO, this should have a handler defined in controllers
		// if you want to keep it
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong"})
		})

		/*
			// Auth routes
			auth := api.Group("/auth")
			{
				auth.POST("/login", login)
				auth.POST("/register", register)
				auth.POST("/logout", logout)
			}




		*/
	}
}
