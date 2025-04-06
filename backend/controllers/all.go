// controllers/people.go
package controllers

import (
	"net/http"

	"backend/models"

	"github.com/gin-gonic/gin"
)

// GetPreferences handles GET requests to fetch preferences (or any data from the database)
func GetPreferences(c *gin.Context) {
	var preferences []models.Preferences 
	c.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	if err := models.DB.Find(&preferences).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return 
	}
	c.JSON(http.StatusOK, preferences) 
}