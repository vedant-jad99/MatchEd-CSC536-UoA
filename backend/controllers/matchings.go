package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"backend/models"

	"github.com/gin-gonic/gin"
)

// GetAllCourses handles GET requests to fetch all courses
func GetAllMatchings(c *gin.Context) {
	var matchings []models.Matching

	// Fetch courses from the database
	result := models.DB.Find(&matchings)
	if result.Error != nil {
		log.Println("Error fetching matchings:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch matchings"})
		return
	}

	// Log the response
	log.Printf("Fetched %d Matchings\n", len(matchings))

	// Return JSON response (wrapper transforms it from Matchings[] to matching: Matchings[])
	c.JSON(http.StatusOK, gin.H{
		"matchings": matchings,
	})
}

func GetAllMatchPairs(c *gin.Context) {
	var matchings []models.Matching
	if err := models.DB.Find(&matchings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch matchings"})
		return
	}

	var results []map[string]interface{}
	for _, m := range matchings {
		var user models.User
		var course models.Course

		if err := models.DB.First(&user, m.UserID).Error; err != nil {
			continue
		}
		if err := models.DB.First(&course, m.CourseID).Error; err != nil {
			continue
		}

		result := map[string]interface{}{
			"id": m.ID,
			"user": map[string]interface{}{
				"ID":    fmt.Sprint(user.ID),
				"name":  user.Name,
				"email": user.Email,
			},
			"course": map[string]interface{}{
				"ID":        fmt.Sprint(course.ID),
				"number":    course.Number,
				"name":      course.Name,
				"campus":    course.Campus,
				"semesters": strings.Split(course.Semesters, ","), // assuming comma-separated in DB
			},
			"status":     "confirmed", // hardcoded or derive from Score
			"confidence": m.Score,
			"timestamp":  time.Now().Format(time.RFC3339), // real timestamp if you add one
		}

		results = append(results, result)
	}

	c.JSON(http.StatusOK, gin.H{"matchings": results})
}
