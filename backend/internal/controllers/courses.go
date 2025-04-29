package controllers

import (
	"backend/internal/models"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// fetches courseSemesters, then fetches the course data for those semesters
// return includes course_sem_id
func HandleFetchAllCoursesFormatted(c *gin.Context) {
	results, err := models.FetchAllCourseSemesters()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create a slice to hold the formatted preferences
	var prettyCourses []map[string]interface{}

	for _, semester := range results {
		// Fetch course by courseSemester courseID
		course, err := models.FetchCourseById(semester.CourseID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				continue // Skip missing course
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Course not found"})
			continue
		}

		// calculate number of sections for a course name
		//count, err := models.CountSemestersForCourse(semester.CourseID)

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				continue // Skip missing course
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Course not found"})
			continue
		}

		// Format the data as expected by the frontend
		formattedCourse := map[string]interface{}{
			"id":       semester.ID,
			"name":     course.Number,
			"sections": semester.Section,
		}
		prettyCourses = append(prettyCourses, formattedCourse)
	}

	c.JSON(http.StatusOK, prettyCourses)
}

func HandleUpsertCourseAndSemester(c *gin.Context) {
	fmt.Println("Hit upsert handler 0")
	type request struct {
		ID     uint   `json:"id"`
		Name   string `json:"name"`
		Number string `json:"number"`
	}

	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	courseSem, err := models.UpsertCourseAndCourseSemester(req.ID, req.Name, req.Number)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upsert course and semester"})
		return
	}

	c.JSON(http.StatusOK, courseSem)
}

// fetches semesters with the context value {semester: }
func HandleFetchAllCourseSemesters(c *gin.Context) {
	results, err := models.FetchAllCourseSemesters()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, results)
}

// fetches semester by id {id: }
func HandleFetchCourseSemester(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := models.FetchCourseSemester(uint(id))
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
	c.JSON(http.StatusOK, gin.H{"status": "created"})
}

// Removes CourseSemester with the given context {id: }
func HandleRemoveCourseSemester(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = models.RemoveCourseSemester(uint(id))
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
	err := models.UpdateCourseSemester(cs.ID, cs.CourseID, cs.Semester, cs.MandatoryLevel, cs.Section, cs.Timeslot)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func HandleBulkUpsertCourseSemesters(c *gin.Context) {
	var input struct {
		CourseSemesters []models.CourseSemester `json:"course_semesters"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	results, err := models.BulkUpsertCourseSemester(input.CourseSemesters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated", "course_semesters": results})
}

func HandleBulkDeleteCourseSemesters(c *gin.Context) {
	var input struct {
		IDs []uint `json:"ids"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := models.BulkDeleteCourseSemesters(input.IDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
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
