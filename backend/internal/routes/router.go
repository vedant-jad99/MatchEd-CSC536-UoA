package routes

import (
	// "net/http"
	"os"

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

		batch := api.Group("/batch")
		{
			batch.POST("/", controllers.HandleBatchPush)
		}

		// Course routes
		courses := api.Group("/course")
		{
			courses.GET("/", controllers.HandleFetchAllCourses)
			courses.GET("/:id", controllers.HandleFetchCourseById)
			courses.POST("/", controllers.HandleUpsertCourse)
			courses.DELETE("/:id", controllers.HandleDeleteCourse)
		}
		// User routes
		users := api.Group("/users")
		{
			users.GET("/", controllers.GetAllUsers)
			//users.GET("/fetch_formatted", controllers.GetAllUsersFormatted)
			users.GET("/:id", controllers.GetUserById)
			users.POST("/", controllers.UpsertUser)
			users.DELETE("/:id", controllers.DeleteUser)
		}

		courseSemesters := api.Group("/course_semester")
		{
			courseSemesters.GET("/", controllers.HandleFetchAllCourseSemesters)
			courseSemesters.GET("/get", controllers.HandleFetchCourseSemester)
			courseSemesters.POST("/add", controllers.HandleAddCourseSemester)
			courseSemesters.DELETE("/remove", controllers.HandleRemoveCourseSemester)
			courseSemesters.PUT("/update", controllers.HandleUpdateCourseSemester)
		}

		preferences := api.Group("/preferences")
		{
			preferences.GET("/fetch_all", controllers.HandleFetchAllPreferences)
			preferences.GET("/fetch_formatted", controllers.HandleFetchAllPreferencesFormatted)
			preferences.GET("/fetch_by_user", controllers.HandleFetchPreferences)
		}

		matchings := api.Group("/matchings")
		{
			matchings.GET("/latest", controllers.HandleFetchLatestMatchings)
			matchings.GET("/by_iteration", controllers.HandleFetchMatchingsByIterationID)
		}
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong"})
		})
	}
}
