package controllers

import (
	matching "backend/internal/engine/core"
	"backend/internal/models"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func getMatchingInput(matchingInput *matching.MatchingInput) error {
	// Fetch all preferences
	preferences, err := models.FetchAllPreferences()
	if err != nil {
		return err
	}
	// Fetch all users
	users, err := models.FetchAllUsers()
	if err != nil {
		return err
	}
	// Fetch all course semesters
	courseSemesters, err := models.FetchAllCourseSemesters()
	if err != nil {
		return err
	}

	for _, user := range users {
		matchingInput.Faculty = append(matchingInput.Faculty, matching.Faculty{
			UserID:        matching.IDType(user.ID),
			NumReqCourses: user.NumReqCourses,
		})
	}
	for _, courseSemester := range courseSemesters {
		matchingInput.Course_s = append(matchingInput.Course_s, matching.CourseSems{
			CourseSemID: matching.IDType(courseSemester.ID),
			CourseID:    matching.IDType(courseSemester.CourseID),
			Semester:    courseSemester.Semester,
			TimeSlot:    courseSemester.Timeslot,
		})
	}

	for _, preference := range preferences {
		matchingInput.Preferences = append(matchingInput.Preferences, matching.Preferences{
			PreferenceID:     matching.IDType(preference.ID),
			UserID:           matching.IDType(preference.UserID),
			CourseSemID:      matching.IDType(preference.CourseSemesterID),
			Semester:         preference.Semester,
			PreferenceLevel:  matching.ConvertPreferenceLevel(preference.PreferenceLevel),
			PreferenceWeight: int64(preference.PreferenceWeight),
		})
	}
	return nil
}

func TriggerMatchingEngine(c *gin.Context) {
	// Call the matching engine function
	var err error
	// err := models
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to trigger matching engine"})
	// 	return
	// }

	mIteration := models.MatchingIteration{
		TriggeredBy: uint(rand.Int31()), // Ensure the value fits within PostgreSQL's int4 range
		Status:      matching.MatchingIterationStatusInitialized,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	mIteration, err = models.CreateMatchingIteration(mIteration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create matching iteration"})
		return
	}
	mInput := matching.MatchingInput{
		Faculty:     []matching.Faculty{},
		Course_s:    []matching.CourseSems{},
		Preferences: []matching.Preferences{},
	}
	err = getMatchingInput(&mInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch matching input"})
		return
	}
	mInterface := matching.NewMatchingInterface()
	var matchingOutput matching.Matching
	matchingOutput, err = mInterface.TriggerMatching(matching.IDType(mIteration.ID), mInput)

	if err != nil {
		mIteration.Status = matching.MatchingIterationStatusError
		mIteration.UpdatedAt = time.Now()
		_, err = models.UpdateMatchingIteration(mIteration)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Matching engine failed. Failed to update matching iteration status."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Matching engine failed"})
		return
	}
	// Save the matching output to the database
	var matchings []models.Matching
	for _, match := range matchingOutput.Matchings {
		matchings = append(matchings, models.Matching{
			MatchingIterationID: uint(mIteration.ID),
			UserID:              uint(match.UserID),
			CourseSemID:         uint(match.CourseSemID),
			Score:               match.MatchingScore,
		})
	}
	if len(matchings) == 0 {
		mIteration.Status = matching.MatchingIterationStatusCompleted
		mIteration.UpdatedAt = time.Now()
		_, err = models.UpdateMatchingIteration(mIteration)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Matching engine failed. Failed to update matching iteration status."})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Matching engine completed. But No matchings found."})
		return
	}
	// Store the matchings in the database
	err = models.StoreMatchings(matchings)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Matching engine completed. Failed to store matchings"})
		return
	}
	mIteration.Status = matching.MatchingIterationStatusCompleted
	mIteration.UpdatedAt = time.Now()
	_, err = models.UpdateMatchingIteration(mIteration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Matching engine completed. Failed to update matching iteration status."})
		return
	}
	// Save the matching output to the database

	c.JSON(http.StatusOK, gin.H{"message": "Matching engine successful", "data": matchings})
}
