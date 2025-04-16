package routes

import (
	"os"

	"backend/internal/controllers" // Import your controllers

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

	// "/api" routes
	setApiHandlers(r)

	// Serve static files from the React app build directory in production
	// r.Use(static.Serve("/", static.LocalFile("../frontend/dist", false)))

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
