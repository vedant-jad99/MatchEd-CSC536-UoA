package controllers

import (
	"backend/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllCourses handles GET requests to fetch all courses
func GetAllMatchings(c *gin.Context) {

	// Fetch courses from the database
	matchings, err := models.FetchLatestMatchings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch matchings"})
		return
	}
	// Return JSON response (wrapper transforms it from Matchings[] to matching: Matchings[])
	c.JSON(http.StatusOK, gin.H{
		"matchings": matchings,
	})
}

// Get matchings with MatchIterationID equal to id,{id:id}
func HandleFetchMatchingsByIterationID(c *gin.Context) {
	var input struct {
		ID uint `json:"id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	matches, err := models.FetchMatchingsByIterationID(input.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, matches)
}

// get matchings with the most recent updated_at field
func HandleFetchLatestMatchings(c *gin.Context) {
	matches, err := models.FetchLatestMatchings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, matches)
}

func HandleFetchLatestMatchingsFormatted(c *gin.Context) {
	matches, err := models.FetchLatestMatchings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create a slice to hold the formatted preferences
	var prettyMatches []map[string]interface{}

	// Loop through preferences and get corresponding course and faculty
	for _, match := range matches {
		// Fetch user (faculty) by ID
		user, err := models.FetchUserById(match.UserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
			continue
		}

		// Fetch course semester by ID
		courseSemester, err := models.FetchCourseSemester(match.CourseSemID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Course not found"})
			continue
		}

		// Fetch course semester by ID
		course, err := models.FetchCourseById(courseSemester.CourseID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Course not found"})
			continue
		}

		// Format the preference data as expected by the frontend
		formattedMatch := map[string]interface{}{
			"id":      match.ID,
			"course":  course.Number,
			"faculty": user.Name,
			"score":   match.Score, // Assuming you want to return PreferenceLevel
		}

		// Append the formatted preference to the slice
		prettyMatches = append(prettyMatches, formattedMatch)
	}
	c.JSON(http.StatusOK, prettyMatches)
}
