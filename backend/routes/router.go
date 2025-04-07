package routes

import (
	"net/http"
	"os"

	"backend/controllers" // Import your controllers

	// "github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
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
	// r.Use(static.Serve("/", static.LocalFile("../frontend/dist", false)))

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

		// People routes
		people := api.Group("/people")
		{
			people.GET("/", controllers.GetAllPeople)
			people.GET("/:id", controllers.GetPersonById)
			people.POST("/", controllers.CreatePerson)
			people.PUT("/:id", controllers.UpdatePerson)
			people.DELETE("/:id", controllers.DeletePerson)
		}

		// User routes
		users := api.Group("/users")
		{
			users.GET("/", controllers.GetAllUsers)
			users.GET("/:id", controllers.GetUserById)
			users.POST("/", controllers.CreateUser)
			users.PUT("/:id", controllers.UpdateUser)
			users.DELETE("/:id", controllers.DeleteUser)
		}

		// "All" routes
		all := api.Group("/all")
		{
			all.GET("/", controllers.GetPreferences)
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



			// Course routes
			courses := api.Group("/courses")
			{
				courses.GET("/", getAllCourses)
				courses.GET("/:id", getCourseById)
				courses.POST("/", createCourse)
				courses.PUT("/:id", updateCourse)
				courses.DELETE("/:id", deleteCourse)
			}
		*/
	}
}
