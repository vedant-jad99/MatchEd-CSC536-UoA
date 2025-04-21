package routes

import (
	"net/http"
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

		courseSemesters := api.Group("/course_semester")
		{
			courseSemesters.GET("/fetch", controllers.HandleFetchCourseSemesters())
			courseSemesters.GET("/get", controllers.HandleFetchCourseSemester())
			courseSemesters.POST("/add", controllers.HandleAddCourseSemester())
			courseSemesters.DELETE("/remove", controllers.HandleRemoveCourseSemester())
			courseSemesters.PUT("/update", controllers.HandleUpdateCourseSemester())
		}

		faculty := api.Group("/faculty")
		{
			faculty.GET("/fetch", controllers.HandleFetchAllFaculty())
			faculty.DELETE("/remove", controllers.HandleRemoveFaculty())
			faculty.POST("/create", controllers.HandleCreateUser())
		}

		preferences := api.Group("/preferences")
		{
			preferences.GET("/fetch_all", controllers.HandleFetchAllPreferences())
			preferences.GET("/fetch_by_user", controllers.HandleFetchPreferences())
		}

		matchings := api.Group("/matchings")
		{
			matchings.GET("/latest", controllers.HandleFetchLatestMatchings())
			matchings.GET("/by_iteration", controllers.HandleFetchMatchingsByIterationID())
			matchings.GET("/pairs", models.GetAllMatchPairs)
		}
		/*
			api.OPTIONS("/*path", func(c *gin.Context) {
				c.Status(http.StatusOK)
			})
		*/
		/*
			r.OPTIONS("/*path", func(c *gin.Context) {
				c.Header("Access-Control-Allow-Origin", "*")
				c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
				c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
				c.Status(http.StatusOK)
			})
		*/
		/*
			// Course routes
			courses := api.Group("/courses")
			{
				courses.GET("/", models.GetAllCourses)
				courses.GET("/:id", models.GetCourseById)
				courses.POST("/", models.CreateCourse)
				courses.PUT("/:id", models.UpdateCourse)
				courses.DELETE("/:id", models.DeleteCourse)
			}
		*/

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
