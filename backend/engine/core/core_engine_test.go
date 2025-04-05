package matching

import (
	"errors"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPreprocessMatchingInput(t *testing.T) {
	matchingInput := MatchingInput{
		faculty:    []IDType{"f1", "f2"},
		course_s:   []IDType{"c1", "c2"},
		preferences: []Preferences{
			{UserID: "f1", CourseID: "c1", PreferenceLevel: 1},
			{UserID: "f2", CourseID: "c2", PreferenceLevel: 2},
		},
	}

	preprocessedInput, err := preprocessMatchingInput(matchingInput)
	assert.NoError(t, err)
	assert.Equal(t, []IDType{"f1", "f2"}, preprocessedInput.faculty_ids)
	assert.Equal(t, []IDType{"c1", "c2"}, preprocessedInput.course_ids)
	assert.Equal(t, 2, len(preprocessedInput.preferences))
	assert.Equal(t, preferenceMapValue{courseID: "c1", preferenceLevel: 1}, preprocessedInput.preference_map["f1"])
	assert.Equal(t, preferenceMapValue{courseID: "c2", preferenceLevel: 2}, preprocessedInput.preference_map["f2"])
}

func TestMatchingEngine(t *testing.T) {
	preprocessedInput := pMatchingInput{
		faculty_ids: []IDType{"f1", "f2"},
		course_ids:  []IDType{"c1", "c2"},
		preferences: []Preferences{
			{UserID: "f1", CourseID: "c1", PreferenceLevel: 1},
			{UserID: "f2", CourseID: "c2", PreferenceLevel: 2},
		},
		fc_map: map[IDType]any{
			"f1": nil,
			"f2": nil,
		},
		c_map: map[IDType]bool{
			"c1": false,
			"c2": false,
		},
		preference_map: map[IDType]preferenceMapValue{
			"f1": {courseID: "c1", preferenceLevel: 1},
			"f2": {courseID: "c2", preferenceLevel: 2},
		},
	}

	rand.Seed(time.Now().UnixNano())
	matching, err := matchingEngine(preprocessedInput)
	assert.NoError(t, err)
	assert.NotNil(t, matching)
	assert.Equal(t, 2, len(matching.matchings))
}

func TestRunMatching(t *testing.T) {
	matchingInput := MatchingInput{
		faculty:    []IDType{"f1", "f2"},
		course_s:   []IDType{"c1", "c2"},
		preferences: []Preferences{
			{UserID: "f1", CourseID: "c1", PreferenceLevel: 1},
			{UserID: "f2", CourseID: "c2", PreferenceLevel: 2},
		},
	}

	matching, err := runMatching(matchingInput)
	assert.NoError(t, err)
	assert.NotNil(t, matching)
	assert.Equal(t, 2, len(matching.matchings))
}

func TestGetInput(t *testing.T) {
	_, err := GetInput()
	assert.Error(t, err) // Since the function is not implemented, it should return an error
}

func TestStartMatching(t *testing.T) {
	matchingIter := MatchingIteration{}
	_, err := StartMatching(matchingIter)
	assert.Error(t, err) // Since GetInput is not implemented, this should return an error
}