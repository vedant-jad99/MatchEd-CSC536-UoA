package models

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO change time reprentation of days and times the course section is scheduled for
type Course struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Name   string `json:"name" gorm:"type:varchar;not null"`
	Number string `json:"number" gorm:"type:varchar;not null"`
	Campus string `json:"campus" gorm:"type:varchar;not null"`
}

func (Course) TableName() string {
	return "match_schema.courses"
}

// GetAllCourses handles GET requests to fetch all courses
func GetAllCourses(c *gin.Context) {
	var courses []Course

	// Fetch courses from the database
	result := DB.Find(&courses)
	if result.Error != nil {
		log.Println("Error fetching courses:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch courses"})
		return
	}

	// Log the response
	log.Printf("Fetched %d courses\n", len(courses))

	// Return JSON response (wrapper transforms it from course[] to courses:course[])
	c.JSON(http.StatusOK, gin.H{
		"courses": courses,
	})
}

// GetcourseById handles GET requests to fetch a course by ID
func GetCourseById(c *gin.Context) {
	id := c.Param("id")

	// Mock fetching a course by ID
	// var course Course
	// if err := DB.First(&course, id).Error; err != nil {
	//     c.JSON(http.StatusNotFound, gin.H{"error": "course not found"})
	//     return
	// }

	c.JSON(http.StatusOK, gin.H{"id": id, "message": "course details fetched"})
}

// CreateCourse handles POST requests to create a new course
func CreateCourse(c *gin.Context) {
	var course Course

	// Bind JSON input to course struct
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Mock saving course to database
	// db.Create(&course)

	c.JSON(http.StatusCreated, gin.H{"message": "course created", "course": course})
}

// Updatecourse handles PUT requests to update an existing course
func UpdateCourse(c *gin.Context) {
	id := c.Param("id")
	var course Course

	// Bind JSON input to course struct
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Mock updating course in database
	// DB.Model(&course).Where("id = ?", id).Updates(course)

	c.JSON(http.StatusOK, gin.H{"message": "course updated", "id": id, "course": course})
}

// Deletecourse handles DELETE requests to remove a course by ID
func DeleteCourse(c *gin.Context) {
	id := c.Param("id")

	// Mock deleting course from database
	// DB.Delete(&Course{}, id)

	c.JSON(http.StatusOK, gin.H{"message": "course deleted", "id": id})
}
