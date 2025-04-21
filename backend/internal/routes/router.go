package routes

import (
	// "net/http"
	"os"

	// Import your controllers

	// "github.com/gin-contrib/static"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	// local models
	"backend/internal/controllers"
	"backend/internal/models"
)

func SetupRouter() *gin.Engine {
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
	// alternative frontend
	r.Use(static.Serve("/", static.LocalFile("../new-frontend/", false)))

	return r
}

// maps the handlers responsible for
// GET, POST, PUT, DELETE
// for the api group "/api"
func setApiHandlers(r *gin.Engine) {
	// need to give the slash at the end for CORS to work because
	// otherwise 301 redirect response will be sent without CORS headers.
	api := r.Group("/api")
	{
		// Matching engine routes
		api.POST("/trigger-matching", controllers.TriggerMatchingEngine)
		// Course routes
		courses := api.Group("/courses")
		{
			courses.GET("/", models.GetAllCourses)
			courses.GET("/:id", models.GetCourseById)
			courses.POST("/", models.CreateCourse)
			courses.PUT("/:id", models.UpdateCourse)
			courses.DELETE("/:id", models.DeleteCourse)
		}

		// Matching routes
		matchings := api.Group("/matchings")
		{
			matchings.GET("/", models.GetAllMatchPairs)
		}
		/*
			// User routes
			users := api.Group("/users")
			{
				users.GET("/", models.GetAllUsers)
				users.GET("/:id", models.GetUserById)
				users.POST("/", models.CreateUser)
				users.PUT("/:id", models.UpdateUser)
				users.DELETE("/:id", models.DeleteUser)
			}
		*/
		// ping the server
		// TODO, this should have a handler defined in controllers
		// if you want to keep it
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong"})
		})
	}
}
