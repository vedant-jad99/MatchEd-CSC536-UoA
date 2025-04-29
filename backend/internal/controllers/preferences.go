package controllers

import (
	"backend/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// gets all preferences for the given semester {semester:semester}
func HandleFetchAllPreferences(c *gin.Context) {
	prefs, err := models.FetchAllPreferences()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, prefs)
}

func HandleFetchAllPreferencesFormatted(c *gin.Context) {
	// Fetch all preferences from the database
	prefs, err := models.FetchAllPreferences()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create a slice to hold the formatted preferences
	var formattedPreferences []map[string]interface{}

	// Loop through preferences and get corresponding course and faculty
	for _, pref := range prefs {
		// Fetch user (faculty) by ID
		user, err := models.FetchUserById(pref.UserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
			continue
		}

		// Fetch course by ID (assuming `FetchCourseById` exists)
		course, err := models.FetchCourseById(pref.CourseSemesterID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Course not found"})
			continue
		}

		// Format the preference data as expected by the frontend
		formattedPreference := map[string]interface{}{
			"course":     course.Number,
			"faculty":    user.Name,
			"preference": pref.PreferenceLevel, // Assuming you want to return PreferenceLevel
		}

		// Append the formatted preference to the slice
		formattedPreferences = append(formattedPreferences, formattedPreference)
	}

	// Return the formatted preferences as a JSON response
	c.JSON(http.StatusOK, formattedPreferences)
}

// gets all preferences by the user_id {user_id:id}
func HandleFetchPreferences(c *gin.Context) {
	var input struct {
		UserID uint `json:"user_id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	prefs, err := models.FetchPreferences(input.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, prefs)
}

func HandleBulkUpsertPreferences(c *gin.Context) {
	var input struct {
		Preferences []models.Preferences `json:"preferences"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := models.BulkUpsertPreferences(input.Preferences)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated", "preferences": result})
}

func HandleBulkDeletePreferences(c *gin.Context) {
	var input struct {
		IDs []uint `json:"ids"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := models.BulkDeletePreferences(input.IDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
