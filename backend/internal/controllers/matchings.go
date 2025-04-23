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
