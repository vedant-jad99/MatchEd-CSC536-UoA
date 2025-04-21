package controllers

import (
	"backend/internal/engine/core"
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
		"bou.ke/monkey"
	)
	
	type mockMatchingInterface struct {
		TriggerMatchingFunc func(iterationID matching.IDType, input matching.MatchingInput) (matching.Matching, error)
	}
	
	func (m *mockMatchingInterface) TriggerMatching(iterationID matching.IDType, input matching.MatchingInput) (matching.Matching, error) {
		return m.TriggerMatchingFunc(iterationID, input)
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
	// mockMatchingOutput := matching.Matching{}

	// Monkey patching
	monkey.Patch(models.FetchAllPreferences, func() ([]models.Preferences, error) {
		return mockPreferences, nil
	})
	monkey.Patch(models.FetchAllUsers, func() ([]models.User, error) {
		return mockUsers, nil
	})
	monkey.Patch(models.FetchAllCourseSemesters, func() ([]models.CourseSemester, error) {
		return mockCourseSemesters, nil
	})
	monkey.Patch(models.CreateMatchingIteration, func(iteration models.MatchingIteration) (models.MatchingIteration, error) {
		return mockIteration, nil
	})
	monkey.Patch(models.UpdateMatchingIteration, func(iteration models.MatchingIteration) (models.MatchingIteration, error) {
		return mockIteration, nil
	})
	// monkey.Patch(matching.NewMatchingInterface, func() matching.MatchingInterface {
	// 	return &mockMatchingInterface{
	// 		TriggerMatchingFunc: func(iterationID matching.IDType, input matching.MatchingInput) (matching.Matching, error) {
	// 			return mockMatchingOutput, nil
	// 		},
	// 	}
	// })

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
	gin.SetMode(gin.TestMode)

	// Monkey patching
	monkey.Patch(models.FetchAllPreferences, func() ([]models.Preferences, error) {
		return nil, errors.New("failed to fetch preferences")
	})

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
