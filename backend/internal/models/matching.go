package models

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Matching struct {
	ID                  uint    `json:"id" gorm:"primaryKey"`
	MatchingIterationId uint    `json:"matching_iteration_id" gorm:"column:matching_iteration_id;not null"`
	UserID              uint    `json:"user_id" gorm:"column:user_id;not null"`
	CourseSemID         uint    `json:"course_sem_id" gorm:"column:course_sem_id;not null"`
	Score               float64 `json:"score" gorm:"not null"`
}

func (Matching) TableName() string {
	return "match_schema.matchings"
}

// GetAllCourses handles GET requests to fetch all courses
func GetAllMatchings(c *gin.Context) {
	var matchings []Matching

	// Fetch courses from the database
	result := DB.Find(&matchings)
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
	var matchings []Matching
	if err := DB.Find(&matchings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch matchings"})
		return
	}

	// build a response json to match the frontend interface
	var results []map[string]interface{}

	for _, m := range matchings {
		var role Role
		var user User
		var course_sem CourseSemester
		var course Course

		// role -> user
		if err := DB.First(&role, m.UserID).Error; err != nil {
			continue
		}
		if err := DB.First(&user, m.UserID).Error; err != nil {
			continue
		}

		// course sem -> course
		if err := DB.First(&course_sem, m.CourseSemID).Error; err != nil {
			continue
		}
		if err := DB.First(&course, course_sem.CourseID).Error; err != nil {
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
				"campus":    "Main Campus",
				"semesters": strings.Split(course_sem.Semester, ","), // assuming comma-separated in DB
			},
			"status":     "confirmed", // hardcoded or derive from Score
			"confidence": m.Score,
			"timestamp":  time.Now().Format(time.RFC3339), // real timestamp if you add one
		}

		results = append(results, result)
	}

	c.JSON(http.StatusOK, gin.H{"matchings": results})
}
