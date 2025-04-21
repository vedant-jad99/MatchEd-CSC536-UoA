package controllers

import (
	"backend/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// fetches semesters with the context value {semester: }
func HandleFetchCourseSemesters() gin.HandlerFunc {
	return func(c *gin.Context) {
		semester := c.Query("semester")
		results, err := models.FetchCourseSemesters(semester)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, results)
	}
}

// fetches semester by id {id: }
func HandleFetchCourseSemester() gin.HandlerFunc {
	return func(c *gin.Context) {
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
}

// creates a new CourseSemester with from context {course_id: , semester: }
func HandleAddCourseSemester() gin.HandlerFunc {
	return func(c *gin.Context) {
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
}

// Removes CourseSemester with the given context {id: }
func HandleRemoveCourseSemester() gin.HandlerFunc {
	return func(c *gin.Context) {
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
}

// Updates a CourseSemester from context
// {id:i, course_id:ci, semester:s, mandatory_level:ml, timeslot:ts}
func HandleUpdateCourseSemester() gin.HandlerFunc {
	return func(c *gin.Context) {
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
}

// creates a new user from context:
// "{name:n, email:e, num_req_courses:nrq}"
func HandleCreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := models.CreateUser(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"status": "created"})
	}
}

// fetches a list of all users
func HandleFetchAllFaculty() gin.HandlerFunc {
	return func(c *gin.Context) {
		faculty, err := models.FetchAllFaculty()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, faculty)
	}
}

// handles removal of user by id {id:id}
func HandleRemoveFaculty() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := models.RemoveFaculty(user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "deleted"})
	}
}

// gets all preferences for the given semester {semester:semester}
func HandleFetchAllPreferences() gin.HandlerFunc {
	return func(c *gin.Context) {
		semester := c.Query("semester")
		prefs, err := models.FetchAllPreferences(semester)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, prefs)
	}
}

// gets all preferences by the user_id {user_id:id}
func HandleFetchPreferences() gin.HandlerFunc {
	return func(c *gin.Context) {
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
}

// Get matchings with MatchIterationID equal to id,{id:id}
func HandleFetchMatchingsByIterationID() gin.HandlerFunc {
	return func(c *gin.Context) {
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
}

// get matchings with the most recent updated_at field
func HandleFetchLatestMatchings() gin.HandlerFunc {
	return func(c *gin.Context) {
		matches, err := models.FetchLatestMatchings()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, matches)
	}
}
