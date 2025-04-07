// controllers/people.go
package controllers

import (
	"net/http"

	"backend/models"

	"github.com/gin-gonic/gin"
)

// GetPreferences handles GET requests and determines what to return from the database to the client 
func GetPreferences(c *gin.Context) {
	var preferences []models.Preferences 
	if err := models.DB.Find(&preferences).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return 
	}
	c.JSON(http.StatusOK, preferences) 
}