package controllers

import (
	"backend/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// fetches semesters with the context value {semester: }
func HandleFetchAllCourseSemesters(c *gin.Context) {
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

// fetches all courses
func HandleFetchAllCourses(c *gin.Context) {
	results, err := models.FetchAllCourses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, results)
}

// fetches course by id {id: }
func HandleFetchCourseById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}
	result, err := models.FetchCourseById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// Upsert Course
func HandleUpsertCourse(c *gin.Context) {
	var input models.Course
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	course, err := models.UpsertCourse(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "created", "course": course})
}

// deletes a course with the given id {id: }
func HandleDeleteCourse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}
	err = models.DeleteCourse(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
