package matching

import (
	"reflect"
	"sort"
	"testing"
)

func TestPreprocessMatchingInput(t *testing.T) {
	input := MatchingInput{
		Faculty: []Faculty{
			{UserID: 1, NumReqCourses: 1},
			{UserID: 2, NumReqCourses: 2},
		},
		Course_s: []CourseSems{
			{CourseSemID: 101},
			{CourseSemID: 102},
			{CourseSemID: 103},
		},
		Preferences: []Preferences{
			{UserID: 1, CourseSemID: 101, PreferenceLevel: 1, PreferenceWeight: 2},
			{UserID: 1, CourseSemID: 103, PreferenceLevel: 1, PreferenceWeight: 4},
			{UserID: 1, CourseSemID: 102, PreferenceLevel: 1, PreferenceWeight: 1},
			{UserID: 2, CourseSemID: 102, PreferenceLevel: 2, PreferenceWeight: 3},
		},
	}

	arr1 := [2]IDType{-1, -1}
	arr2 := [2]IDType{-1, -1}
	expected := pMatchingInput{
		faculty_ids:  []IDType{1, 2},
		course_s_ids: []IDType{101, 102, 103},
		preferences: []Preferences{
			{UserID: 1, CourseSemID: 101, PreferenceLevel: 1, PreferenceWeight: 2},
			{UserID: 1, CourseSemID: 103, PreferenceLevel: 1, PreferenceWeight: 4},
			{UserID: 1, CourseSemID: 102, PreferenceLevel: 1, PreferenceWeight: 1},
			{UserID: 2, CourseSemID: 102, PreferenceLevel: 2, PreferenceWeight: 3},
		},
		fc_map: map[IDType]*[2]IDType{
			1: &arr1,
			2: &arr2,
		},
		c_map: map[IDType]bool{
			101: false,
			102: false,
			103: false,
		},
		f2rc_map: map[IDType]int{
			1: 1,
			2: 2,
		},
		preference_map: map[IDType]map[int64][]pData{
			1: {1: []pData{{103, 4}, {101, 2}, {102, 1}}},
			2: {2: []pData{{102, 3}}},
		},
	}

	result := preprocessMatchingInput(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestMatchingEngine(t *testing.T) {
	input := pMatchingInput{
		faculty_ids:  []IDType{1, 2},
		course_s_ids: []IDType{101, 102, 103},
		preferences: []Preferences{
			{UserID: 1, CourseSemID: 101, PreferenceLevel: 1, PreferenceWeight: 2},
			{UserID: 1, CourseSemID: 103, PreferenceLevel: 1, PreferenceWeight: 4},
			{UserID: 1, CourseSemID: 102, PreferenceLevel: 1, PreferenceWeight: 1},
			{UserID: 2, CourseSemID: 102, PreferenceLevel: 2, PreferenceWeight: 3},
		},
		fc_map: map[IDType]*[2]IDType{
			1: &[2]IDType{-1,-1},
			2: &[2]IDType{-1,-1},
		},
		c_map: map[IDType]bool{
			101: false,
			102: false,
			103: false,
		},
		f2rc_map: map[IDType]int{
			1: 1,
			2: 2,
		},
		preference_map: map[IDType]map[int64][]pData{
			1: {1: []pData{{103, 4}, {101, 2}, {102, 1}}},
			2: {2: []pData{{102, 3}}},
		},
	}

	matchingIter := IDType(1)

	expected := Matching{
		Matchings: []MatchingElement{
			{MatchingID: -1, MatchingIterationID: matchingIter, UserID: 1, CourseSemID: 103, MatchingScore: 1.0},
			{MatchingID: -1, MatchingIterationID: matchingIter, UserID: 2, CourseSemID: 102, MatchingScore: 1.0},
		},
	}
	// Simulate the matching engine logic

	result, err := matchingEngine(input, matchingIter)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(result.Matchings) != len(expected.Matchings) {
		t.Errorf("Expected %d matchings, got %d", len(expected.Matchings), len(result.Matchings))
	}

	for _, match := range result.Matchings {
		if match.MatchingIterationID != matchingIter {
			t.Errorf("Expected MatchingIterationID %d, got %d", matchingIter, match.MatchingIterationID)
		}
		if match.MatchingScore != 1.0 {
			t.Errorf("Expected MatchingScore 1.0, got %f", match.MatchingScore)
		}
	}
	// sort the result and expected's matchings using userid
	sort.Slice(result.Matchings, func(i, j int) bool {
		return result.Matchings[i].UserID < result.Matchings[j].UserID
	})
	sort.Slice(expected.Matchings, func(i, j int) bool {
		return expected.Matchings[i].UserID < expected.Matchings[j].UserID
	})
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestRunMatching(t *testing.T) {
	input := MatchingInput{
		Faculty: []Faculty{
			{UserID: 1, NumReqCourses: 2},
			{UserID: 2, NumReqCourses: 2},
		},
		Course_s: []CourseSems{
			{CourseSemID: 101},
			{CourseSemID: 102},
		},
		Preferences: []Preferences{
			{UserID: 1, CourseSemID: 101, PreferenceLevel: 1, PreferenceWeight: 2},
			{UserID: 1, CourseSemID: 103, PreferenceLevel: 1, PreferenceWeight: 4},
			{UserID: 1, CourseSemID: 102, PreferenceLevel: 1, PreferenceWeight: 1},
			{UserID: 2, CourseSemID: 102, PreferenceLevel: 2, PreferenceWeight: 3},
		},
	}

	matchingIter := IDType(1)

	expected := Matching{
		Matchings: []MatchingElement{
			{MatchingID: -1, MatchingIterationID: matchingIter, UserID: 1, CourseSemID: 103, MatchingScore: 1.0},
			{MatchingID: -1, MatchingIterationID: matchingIter, UserID: 1, CourseSemID: 101, MatchingScore: 1.0},
			{MatchingID: -1, MatchingIterationID: matchingIter, UserID: 2, CourseSemID: 102, MatchingScore: 1.0},
		},
	}
	// Simulate the runMatching logic

	result, err := runMatching(input, matchingIter)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(result.Matchings) != len(expected.Matchings) {
		t.Errorf("Expected %d matchings, got %d", len(expected.Matchings), len(result.Matchings))
	}

	for _, match := range result.Matchings {
		if match.MatchingIterationID != matchingIter {
			t.Errorf("Expected MatchingIterationID %d, got %d", matchingIter, match.MatchingIterationID)
		}
		if match.MatchingScore != 1.0 {
			t.Errorf("Expected MatchingScore 1.0, got %f", match.MatchingScore)
		}
	}
	// sort the result and expected's matchings using userid
	sort.Slice(result.Matchings, func(i, j int) bool {
		return result.Matchings[i].UserID < result.Matchings[j].UserID
	})
	sort.Slice(expected.Matchings, func(i, j int) bool {
		return expected.Matchings[i].UserID < expected.Matchings[j].UserID
	})
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestCoreStartMatching(t *testing.T) {
	matchingIter := IDType(1)
	input := MatchingInput{
		Faculty: []Faculty{
			{UserID: 1, NumReqCourses: 1},
			{UserID: 2, NumReqCourses: 1},
		},
		Course_s: []CourseSems{
			{CourseSemID: 101},
			{CourseSemID: 102},
		},
		Preferences: []Preferences{
			{UserID: 1, CourseSemID: 101, PreferenceLevel: 1, PreferenceWeight: 2},
			{UserID: 1, CourseSemID: 103, PreferenceLevel: 1, PreferenceWeight: 4},
			{UserID: 1, CourseSemID: 102, PreferenceLevel: 1, PreferenceWeight: 1},
			{UserID: 2, CourseSemID: 102, PreferenceLevel: 2, PreferenceWeight: 3},
		},
	}

	expected := Matching{
		Matchings: []MatchingElement{
			{MatchingID: -1, MatchingIterationID: matchingIter, UserID: 1, CourseSemID: 103, MatchingScore: 1.0},
			{MatchingID: -1, MatchingIterationID: matchingIter, UserID: 2, CourseSemID: 102, MatchingScore: 1.0},
		},
	}

	matching, err := StartMatching(matchingIter, input)
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
