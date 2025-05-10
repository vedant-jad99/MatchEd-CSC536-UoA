package controllers

import (
	"backend/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Struct to capture the incoming batch changes
type BatchPushRequest struct {
	FacultyChanges    []Change `json:"facultyChanges"`
	CourseChanges     []Change `json:"courseChanges"`
	PreferenceChanges []Change `json:"preferenceChanges"`
}

// Struct for each change (add, remove, edit)
type Change struct {
	Type        string `json:"type"` // add, remove, edit
	Name        string `json:"name"`
	NewName     string `json:"newName,omitempty"`
	CourseName  string `json:"courseName,omitempty"`
	FacultyName string `json:"facultyName,omitempty"`
	NewColor    string `json:"newColor,omitempty"`
}

// Handler for batch updates
func HandleBatchPush(c *gin.Context) {
	var request BatchPushRequest

	// Bind the incoming JSON to the request struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Process faculty changes
	for _, change := range request.FacultyChanges {
		switch change.Type {
		case "add":
			// Add new faculty
			_, err := models.AddUserByName(change.Name)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		case "remove":
			// Remove faculty by name
			err := models.RemoveUserByName(change.Name)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		case "edit":
			// Edit faculty name
			_, err := models.EditUserByName(change.Name, change.NewName)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid faculty change type"})
			return
		}
	}

	// Process course changes
	for _, change := range request.CourseChanges {
		switch change.Type {
		case "add":
			// Add CourseSemester (with Course)
			_, err := models.AddCourseSemesterByName(change.Name, "", "", "")
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		case "remove":
			// Remove CourseSemester by name
			err := models.RemoveCourseSemesterByName(change.Name)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		case "edit":
			// Edit CourseSemester by name
			_, err := models.EditCourseSemesterByName(change.Name, change.NewName, "", "", "")
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course change type"})
			return
		}
	}

	// Process preference changes
	for _, change := range request.PreferenceChanges {
		switch change.Type {
		case "edit":
			// Edit Preference Color by Course Name and Faculty Name
			_, err := models.EditPreferenceColorByCourseAndUser(change.CourseName, change.FacultyName, change.NewColor)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid preference change type"})
			return
		}
	}

	// Return success message if all changes are applied
	c.JSON(http.StatusOK, gin.H{"status": "Batch update successful"})
}
