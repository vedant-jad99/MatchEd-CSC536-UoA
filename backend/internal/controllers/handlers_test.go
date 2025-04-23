package controllers

import (
	"backend/internal/models"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// mock setup
func init() {
	gin.SetMode(gin.TestMode)
}

// mock DB responses (you can overwrite models.* with stubs)
func TestHandleFetchCourseSemesters(t *testing.T) {
	original := models.FetchCourseSemesters
	models.FetchCourseSemesters = func() ([]models.CourseSemester, error) {
		return []models.CourseSemester{}, nil
	}
	defer func() { models.FetchCourseSemesters = original }()

	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	HandleFetchCourseSemesters(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestHandleFetchCourseSemester_BadRequest(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer([]byte("bad json")))
	c.Request.Header.Set("Content-Type", "application/json")

	HandleFetchCourseSemester(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestHandleAddCourseSemester(t *testing.T) {
	original := models.AddCourseSemester
	models.AddCourseSemester = func(courseID uint, semester string) error {
		return nil
	}
	defer func() { models.AddCourseSemester = original }()

	payload := map[string]interface{}{"course_id": 1, "semester": "Fall"}
	b, _ := json.Marshal(payload)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(b))
	c.Request.Header.Set("Content-Type", "application/json")

	HandleAddCourseSemester(c)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestHandleRemoveCourseSemester_Error(t *testing.T) {
	original := models.RemoveCourseSemester
	models.RemoveCourseSemester = func(id uint) error {
		return errors.New("fail")
	}
	defer func() { models.RemoveCourseSemester = original }()

	payload := map[string]interface{}{"id": 1}
	b, _ := json.Marshal(payload)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(b))
	c.Request.Header.Set("Content-Type", "application/json")

	HandleRemoveCourseSemester(c)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestHandleUpdateCourseSemester(t *testing.T) {
	original := models.UpdateCourseSemester
	models.UpdateCourseSemester = func(id, courseID uint, semester, level, slot string) error {
		return nil
	}
	defer func() { models.UpdateCourseSemester = original }()

	cs := models.CourseSemester{ID: 1, CourseID: 2, Semester: "Spring", MandatoryLevel: "High", Timeslot: "MWF"}
	b, _ := json.Marshal(cs)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(b))
	c.Request.Header.Set("Content-Type", "application/json")

	HandleUpdateCourseSemester(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestHandleCreateUser(t *testing.T) {
	original := models.UpsertUser
	models.UpsertUser = func(u models.User) (models.User, error) { return u, nil }
	defer func() { models.UpsertUser = original }()

	u := models.User{Name: "John", Email: "john@example.com", NumReqCourses: 2}
	b, _ := json.Marshal(u)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(b))
	c.Request.Header.Set("Content-Type", "application/json")

	UpsertUser(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestHandleFetchAllFaculty(t *testing.T) {
	original := models.FetchAllUsers
	models.FetchAllUsers = func() ([]models.User, error) {
		return []models.User{{Name: "Alice"}}, nil
	}
	defer func() { models.FetchAllUsers = original }()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/", nil)

	GetAllUsers(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestHandleRemoveFaculty(t *testing.T) {
	original := models.DeleteUser
	models.DeleteUser = func(id uint) error { return nil }
	defer func() { models.DeleteUser = original }()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Request = httptest.NewRequest(http.MethodDelete, "/1", nil)

	DeleteUser(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestHandleFetchAllPreferences(t *testing.T) {
	original := models.FetchAllPreferences
	models.FetchAllPreferences = func() ([]models.Preferences, error) {
		return []models.Preferences{}, nil
	}
	defer func() { models.FetchAllPreferences = original }()

	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	HandleFetchAllPreferences(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestHandleFetchPreferences(t *testing.T) {
	original := models.FetchPreferences
	models.FetchPreferences = func(userID uint) ([]models.Preferences, error) {
		return []models.Preferences{{UserID: userID}}, nil
	}
	defer func() { models.FetchPreferences = original }()

	b, _ := json.Marshal(map[string]uint{"user_id": 1})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(b))
	c.Request.Header.Set("Content-Type", "application/json")

	HandleFetchPreferences(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestHandleFetchMatchingsByIterationID(t *testing.T) {
	original := models.FetchMatchingsByIterationID
	models.FetchMatchingsByIterationID = func(id uint) ([]models.Matching, error) {
		return []models.Matching{{MatchingIterationID: id}}, nil
	}
	defer func() { models.FetchMatchingsByIterationID = original }()

	b, _ := json.Marshal(map[string]uint{"id": 5})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(b))
	c.Request.Header.Set("Content-Type", "application/json")

	HandleFetchMatchingsByIterationID(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestHandleFetchLatestMatchings(t *testing.T) {
	original := models.FetchLatestMatchings
	models.FetchLatestMatchings = func() ([]models.Matching, error) {
		return []models.Matching{{MatchingIterationID: 42}}, nil
	}
	defer func() { models.FetchLatestMatchings = original }()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/", nil)

	HandleFetchLatestMatchings(c)
	assert.Equal(t, http.StatusOK, w.Code)
}
