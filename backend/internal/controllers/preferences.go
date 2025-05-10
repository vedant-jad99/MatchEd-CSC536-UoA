package controllers

import (
	"backend/internal/models"
	"net/http"
	"sort"
	"strconv"

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

func DeletePreference(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	err = models.DeletePreference(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Preference"})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func HandleFetchAllPreferencesFormatted(c *gin.Context) {
	prefs, err := models.FetchAllPreferences()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var formattedPreferences []map[string]interface{}

	for _, pref := range prefs {
		user, err := models.FetchUserById(pref.UserID)
		if err != nil {
			continue // skip if user not found
		}

		courseSemester, err := models.FetchCourseSemester(pref.CourseSemesterID)
		if err != nil {
			continue // skip if course semester not found
		}

		course, err := models.FetchCourseById(courseSemester.CourseID)
		if err != nil {
			continue // skip if course not found
		}

		formattedPreference := map[string]interface{}{
			"id":         pref.ID,
			"course":     course.Number,
			"faculty":    user.Name,
			"preference": pref.PreferenceLevel,
			"weight":     pref.PreferenceWeight,
			"courseload": user.NumReqCourses,
		}

		formattedPreferences = append(formattedPreferences, formattedPreference)
	}

	sort.Slice(formattedPreferences, func(i, j int) bool {
		return formattedPreferences[i]["course"].(string) < formattedPreferences[j]["course"].(string)
	})

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

// fetch preference by id
func HandleFetchPreferenceById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := models.FetchPreferenceById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// Upsert preference
func HandleUpsertPreference(c *gin.Context) {
	var input models.Preferences
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	preference, err := models.UpsertPreference(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, preference)
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
