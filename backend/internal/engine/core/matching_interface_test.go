package matching

import (
	"sort"
	"reflect"
	"testing"
)

func TestNewMatchingInterface(t *testing.T) {
	expected :=  MatchingInterface{m_StatusMap: make(map[IDType]string)}
	result := NewMatchingInterface()

	if !reflect.DeepEqual(*result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestTriggerMatching(t *testing.T) {
	matchingInterface := NewMatchingInterface()
	if matchingInterface == nil {
		t.Errorf("Failed to create matching interface");
	}

	matchingIter := IDType(1)
	input := MatchingInput{
		faculty: []Faculty{
			{UserID: 1, NumReqCourses: 2},
			{UserID: 2, NumReqCourses: 2},
		},
		course_s: []CourseSems{
			{CourseSemID: 101},
			{CourseSemID: 102},
		},
		preferences: []Preferences{
			{UserID: 1, CourseSemID: 101, PreferenceLevel: 1, PreferenceWeight: 2},
			{UserID: 1, CourseSemID: 103, PreferenceLevel: 1, PreferenceWeight: 4},
			{UserID: 1, CourseSemID: 102, PreferenceLevel: 1, PreferenceWeight: 1},
			{UserID: 2, CourseSemID: 102, PreferenceLevel: 2, PreferenceWeight: 3},
		},
	}

	expected := Matching{
		Matchings: []MatchingElement{
			{MatchingID: -1, MatchingIterationID: matchingIter, UserID: 1, CourseSemID: 103, MatchingScore: 1.0},
			{MatchingID: -1, MatchingIterationID: matchingIter, UserID: 1, CourseSemID: 101, MatchingScore: 1.0},
			{MatchingID: -1, MatchingIterationID: matchingIter, UserID: 2, CourseSemID: 102, MatchingScore: 1.0},
		},
	}

	matching, err := matchingInterface.TriggerMatching(matchingIter, input)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(matching.Matchings) != len(expected.Matchings)  {
		t.Errorf("Expected %d matchings, got %d", len(expected.Matchings), len(matching.Matchings))
	}

	for _, match := range matching.Matchings {
		if match.MatchingIterationID != matchingIter {
			t.Errorf("Expected MatchingIterationID %d, got %d", matchingIter, match.MatchingIterationID)
		}
		if match.MatchingScore != 1.0 {
			t.Errorf("Expected MatchingScore 1.0, got %f", match.MatchingScore)
		}
	}
	// sort the result and expected's matchings using userid
	sort.Slice(matching.Matchings, func(i, j int) bool {
		return matching.Matchings[i].UserID < matching.Matchings[j].UserID
	})
	sort.Slice(expected.Matchings, func(i, j int) bool {
		return expected.Matchings[i].UserID < expected.Matchings[j].UserID
	})
	if !reflect.DeepEqual(matching, expected) {
		t.Errorf("Expected %v, got %v", expected, matching)
	}
}
