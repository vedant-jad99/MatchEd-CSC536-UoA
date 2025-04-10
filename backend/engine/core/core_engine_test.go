package matching

import (
	"encoding/json"
	"fmt"
	"os"
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

	arr1 := [2]IDType{-1, -1}
	arr2 := [2]IDType{-1, -1}
	expected := pMatchingInput{
		faculty_ids:  []IDType{1, 2},
		course_s_ids: []IDType{101, 102},
		preferences: []Preferences{
			{UserID: 1, CourseSemID: 101, PreferenceLevel: 3},
			{UserID: 2, CourseSemID: 102, PreferenceLevel: 2},
		},
		fc_map: map[IDType]*[2]IDType{
			1: &arr1,
			2: &arr2,
		},
		c_map: map[IDType]bool{
			101: false,
			102: false,
		},
		preference_map: map[IDType]map[int64][]IDType{
			1: {3: []IDType{101}},
			2: {2: []IDType{102}},
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
		fc_map: map[IDType]*[2]IDType{
			1: &[2]IDType{-1,-1},
			2: &[2]IDType{-1,-1},
		},
		c_map: map[IDType]bool{
			101: false,
			102: false,
		},
		preference_map: map[IDType]map[int64][]IDType{
			1: {3: []IDType{101}},
			2: {2: []IDType{102}},
		},
	}

	matchingIter := IDType(1)

	expected := Matching{
		Matchings: []MatchingElement{
			{MatchingID: -1, MatchingIterationID: matchingIter, UserID: 1, CourseSemID: 101, MatchingScore: 1.0},
			{MatchingID: -1, MatchingIterationID: matchingIter, UserID: 2, CourseSemID: 102, MatchingScore: 1.0},
		},
	}
	// Simulate the matching engine logic

	result, err := matchingEngine(input, matchingIter)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(result.Matchings) != 2 {
		t.Errorf("Expected 2 matchings, got %d", len(result.Matchings))
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
		Matchings: []MatchingElement{
			{MatchingID: -1, MatchingIterationID: matchingIter, UserID: 1, CourseSemID: 101, MatchingScore: 1.0},
			{MatchingID: -1, MatchingIterationID: matchingIter, UserID: 2, CourseSemID: 102, MatchingScore: 1.0},
		},
	}
	// Simulate the runMatching logic

	result, err := runMatching(input, matchingIter)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(result.Matchings) != 2 {
		t.Errorf("Expected 2 matchings, got %d", len(result.Matchings))
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
	matching, err := StartMatching(matchingIter)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	fmt.Println("Matching started successfully:")

	// Output matching data to a JSON file
	filePath := "matching_output.json"
	file, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("Failed to create output file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(matching); err != nil {
		t.Fatalf("Failed to write matching data to JSON file: %v", err)
	}

	fmt.Printf("Matching data successfully written to %s\n", filePath)
}
