package controllers

import (
	"backend/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// fetches semesters with the context value {semester: }
func HandleFetchCourseSemesters(c *gin.Context) {
	results, err := models.FetchCourseSemesters()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, results)
}

// fetches semester by id {id: }
func HandleFetchCourseSemester(c *gin.Context) {
	var input struct {
		ID uint `json:"id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := models.FetchCourseSemester(input.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// creates a new CourseSemester with from context {course_id: , semester: }
func HandleAddCourseSemester(c *gin.Context) {
	var input struct {
		CourseID uint   `json:"course_id"`
		Semester string `json:"semester"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := models.AddCourseSemester(input.CourseID, input.Semester)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "created"})
}

// Removes CourseSemester with the given context {id: }
func HandleRemoveCourseSemester(c *gin.Context) {
	var input struct {
		ID uint `json:"id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := models.RemoveCourseSemester(input.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

// Updates a CourseSemester from context
// {id:i, course_id:ci, semester:s, mandatory_level:ml, timeslot:ts}
func HandleUpdateCourseSemester(c *gin.Context) {
	var cs models.CourseSemester
	if err := c.ShouldBindJSON(&cs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := models.UpdateCourseSemester(cs.ID, cs.CourseID, cs.Semester, cs.MandatoryLevel, cs.Timeslot)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}
