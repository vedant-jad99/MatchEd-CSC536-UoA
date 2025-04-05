package matching

import (
	"reflect"
	"sort"
	"testing"
)

func TestPreprocessMatchingInput(t *testing.T) {
	input := MatchingInput{
		faculty: []Faculty{
			{UserID: 1},
			{UserID: 2},
		},
		course_s: []CourseSems{
			{CourseSemID: 101},
			{CourseSemID: 102},
		},
		preferences: []Preferences{
			{UserID: 1, CourseSemID: 101, PreferenceLevel: 3},
			{UserID: 2, CourseSemID: 102, PreferenceLevel: 2},
		},
	}

	expected := pMatchingInput{
		faculty_ids:  []IDType{1, 2},
		course_s_ids: []IDType{101, 102},
		preferences: []Preferences{
			{UserID: 1, CourseSemID: 101, PreferenceLevel: 3},
			{UserID: 2, CourseSemID: 102, PreferenceLevel: 2},
		},
		fc_map: map[IDType]IDType{
			1: -1,
			2: -1,
		},
		c_map: map[IDType]bool{
			101: false,
			102: false,
		},
		preference_map: map[IDType]map[int64]IDType{
			1: {3: 101},
			2: {2: 102},
		},
	}

	result, err := preprocessMatchingInput(input)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestMatchingEngine(t *testing.T) {
	input := pMatchingInput{
		faculty_ids:  []IDType{1, 2},
		course_s_ids: []IDType{101, 102},
		preferences: []Preferences{
			{UserID: 1, CourseSemID: 101, PreferenceLevel: 3},
			{UserID: 2, CourseSemID: 102, PreferenceLevel: 2},
		},
		fc_map: map[IDType]IDType{
			1: -1,
			2: -1,
		},
		c_map: map[IDType]bool{
			101: false,
			102: false,
		},
		preference_map: map[IDType]map[int64]IDType{
			1: {3: 101},
			2: {2: 102},
		},
	}

	matchingIter := IDType(1)

	expected := Matching{
		matchings: []MatchingElement{
			{MatchingID: -1, MatchingIterationID: matchingIter, UserID: 1, CourseSemID: 101, MatchingScore: 1.0},
			{MatchingID: -1, MatchingIterationID: matchingIter, UserID: 2, CourseSemID: 102, MatchingScore: 1.0},
		},
	}
	// Simulate the matching engine logic

	result, err := matchingEngine(input, matchingIter)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(result.matchings) != 2 {
		t.Errorf("Expected 2 matchings, got %d", len(result.matchings))
	}

	for _, match := range result.matchings {
		if match.MatchingIterationID != matchingIter {
			t.Errorf("Expected MatchingIterationID %d, got %d", matchingIter, match.MatchingIterationID)
		}
		if match.MatchingScore != 1.0 {
			t.Errorf("Expected MatchingScore 1.0, got %f", match.MatchingScore)
		}
	}
	// sort the result and expected's matchings using userid
	sort.Slice(result.matchings, func(i, j int) bool {
		return result.matchings[i].UserID < result.matchings[j].UserID
	})
	sort.Slice(expected.matchings, func(i, j int) bool {
		return expected.matchings[i].UserID < expected.matchings[j].UserID
	})
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestRunMatching(t *testing.T) {
	input := MatchingInput{
		faculty: []Faculty{
			{UserID: 1},
			{UserID: 2},
		},
		course_s: []CourseSems{
			{CourseSemID: 101},
			{CourseSemID: 102},
		},
		preferences: []Preferences{
			{UserID: 1, CourseSemID: 101, PreferenceLevel: 3},
			{UserID: 2, CourseSemID: 102, PreferenceLevel: 2},
		},
	}

	matchingIter := IDType(1)

	expected := Matching{
		matchings: []MatchingElement{
			{MatchingID: -1, MatchingIterationID: matchingIter, UserID: 1, CourseSemID: 101, MatchingScore: 1.0},
			{MatchingID: -1, MatchingIterationID: matchingIter, UserID: 2, CourseSemID: 102, MatchingScore: 1.0},
		},
	}
	// Simulate the runMatching logic

	result, err := runMatching(input, matchingIter)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(result.matchings) != 2 {
		t.Errorf("Expected 2 matchings, got %d", len(result.matchings))
	}

	for _, match := range result.matchings {
		if match.MatchingIterationID != matchingIter {
			t.Errorf("Expected MatchingIterationID %d, got %d", matchingIter, match.MatchingIterationID)
		}
		if match.MatchingScore != 1.0 {
			t.Errorf("Expected MatchingScore 1.0, got %f", match.MatchingScore)
		}
	}
	// sort the result and expected's matchings using userid
	sort.Slice(result.matchings, func(i, j int) bool {
		return result.matchings[i].UserID < result.matchings[j].UserID
	})
	sort.Slice(expected.matchings, func(i, j int) bool {
		return expected.matchings[i].UserID < expected.matchings[j].UserID
	})
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
