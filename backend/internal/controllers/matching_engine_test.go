package controllers

import (
	matching "backend/internal/engine/core"
	"backend/internal/models"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

// mock setup
func init() {
	gin.SetMode(gin.TestMode)
}

func TestTriggerMatchingEngine_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Mock data
	mockPreferences := []models.Preferences{
		{ID: 1, UserID: 1, CourseSemesterID: 1, Semester: "Fall", PreferenceLevel: "High", PreferenceWeight: 10},
	}
	mockUsers := []models.User{
		{ID: 1, NumReqCourses: 2},
	}
	mockCourseSemesters := []models.CourseSemester{
		{ID: 1, CourseID: 1, Semester: "Fall", Timeslot: "Morning"},
	}
	mockIteration := models.MatchingIteration{
		ID:          1,
		TriggeredBy: 1,
		Status:      matching.MatchingIterationStatusInitialized,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	originalFetchAllPreferences := models.FetchAllPreferences
	defer func() { models.FetchAllPreferences = originalFetchAllPreferences }()
	originalFetchAllUsers := models.FetchAllUsers
	defer func() { models.FetchAllUsers = originalFetchAllUsers }()
	originalFetchAllCourseSemesters := models.FetchAllCourseSemesters
	defer func() { models.FetchAllCourseSemesters = originalFetchAllCourseSemesters }()
	originalCreateMatchingIteration := models.CreateMatchingIteration
	defer func() { models.CreateMatchingIteration = originalCreateMatchingIteration }()
	originalUpdateMatchingIteration := models.UpdateMatchingIteration
	defer func() { models.UpdateMatchingIteration = originalUpdateMatchingIteration }()

	// Mocking the functions
	models.FetchAllPreferences = func() ([]models.Preferences, error) {
		return mockPreferences, nil
	}
	models.FetchAllUsers = func() ([]models.User, error) {
		return mockUsers, nil
	}
	models.FetchAllCourseSemesters = func() ([]models.CourseSemester, error) {
		return mockCourseSemesters, nil
	}
	models.CreateMatchingIteration = func(iteration models.MatchingIteration) (models.MatchingIteration, error) {
		return mockIteration, nil
	}
	models.UpdateMatchingIteration = func(iteration models.MatchingIteration) (models.MatchingIteration, error) {
		iteration.Status = matching.MatchingIterationStatusCompleted
		return iteration, nil
	}
	// Create a test request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	reqBody := map[string]interface{}{}
	jsonBody, _ := json.Marshal(reqBody)
	c.Request = httptest.NewRequest(http.MethodPost, "/trigger-matching", bytes.NewBuffer(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	// Call the function
	TriggerMatchingEngine(c)

	// Assertions
	require.Equal(t, http.StatusOK, w.Code)
}

func TestTriggerMatchingEngine_FetchPreferencesError(t *testing.T) {
	mockIteration := models.MatchingIteration{
		ID:          1,
		TriggeredBy: 1,
		Status:      matching.MatchingIterationStatusInitialized,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	originalCreateMatchingIteration := models.CreateMatchingIteration
	defer func() { models.CreateMatchingIteration = originalCreateMatchingIteration }()
	originalFetchAllPreferences := models.FetchAllPreferences
	defer func() { models.FetchAllPreferences = originalFetchAllPreferences }()

	models.FetchAllPreferences = func() ([]models.Preferences, error) {
		return nil, errors.New("failed to fetch preferences")
	}
	models.CreateMatchingIteration = func(iteration models.MatchingIteration) (models.MatchingIteration, error) {
		return mockIteration, nil
	}
	// Create a test request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	reqBody := map[string]interface{}{}
	jsonBody, _ := json.Marshal(reqBody)
	c.Request = httptest.NewRequest(http.MethodPost, "/trigger-matching", bytes.NewBuffer(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	// Call the function
	TriggerMatchingEngine(c)

	// Assertions
	require.Equal(t, http.StatusInternalServerError, w.Code)
}
