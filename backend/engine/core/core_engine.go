package matching

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"time"
)

type pMatchingInput struct {
	faculty_ids    []IDType
	course_s_ids   []IDType
	preferences    []Preferences
	fc_map         map[IDType]*[2]IDType
	c_map          map[IDType]bool
	preference_map map[IDType]map[int64][]IDType //FacultyID --> preference (3,2,1) --> courseSemID[]
}

// TODO: Get the input from the database. Communicates to the matching interface
func GetInput(file string) (MatchingInput, error) {
	// Open the JSON file
	data, err := os.ReadFile(file)
	if err != nil {
		return MatchingInput{}, fmt.Errorf("failed to read input file: %w", err)
	}

	// Define a struct to match the JSON structure
	type rawInput struct {
		User         map[string]string `json:"user"`
		Faculty      []int64           `json:"faculty"`
		CourseSemIDs []int64           `json:"course_sem_ids"`
		Preferences  []struct {
			ID          int64  `json:"id"`
			UserID      int64  `json:"user_id"`
			CourseSemID int64  `json:"course_sem_id"`
			Sem         string `json:"sem"`
			Level       string `json:"level"`
		} `json:"preferences"`
	}

	// Parse the JSON data
	var raw rawInput
	if err := json.Unmarshal(data, &raw); err != nil {
		return MatchingInput{}, fmt.Errorf("failed to parse input JSON: %w", err)
	}

	// Convert raw input to MatchingInput
	var matchingInput MatchingInput
	for _, facultyID := range raw.Faculty {
		matchingInput.faculty = append(matchingInput.faculty, Faculty{UserID: IDType(facultyID)})
	}
	for _, courseSemID := range raw.CourseSemIDs {
		matchingInput.course_s = append(matchingInput.course_s, CourseSems{CourseSemID: IDType(courseSemID)})
	}
	for _, pref := range raw.Preferences {
		level := PreferenceLevelInvalid
		switch pref.Level {
		case "r":
			level = PreferenceLevelRed
		case "y":
			level = PreferenceLevelYellow
		case "g":
			level = PreferenceLevelGreen
		}
		matchingInput.preferences = append(matchingInput.preferences, Preferences{
			PreferenceID:    IDType(pref.ID),
			UserID:          IDType(pref.UserID),
			CourseSemID:     IDType(pref.CourseSemID),
			Semester:        pref.Sem,
			PreferenceLevel: level,
		})
	}

	return matchingInput, nil
}

func StartMatching(matchingIter IDType) (Matching, error) {
	preferences, err := GetInput("sample/preferences_clean.json")
	if err != nil {
		return Matching{}, err // TODO: Custom error type?
	}
	fmt.Println("Preferences: ", preferences)
	//UpdateMatchingIterStatus(matchingIter);
	matching, err := runMatching(preferences, matchingIter)
	if err != nil {
		return Matching{}, err
	}
	return matching, nil
}

func preprocessMatchingInput(mI MatchingInput) (pMatchingInput, error) {
	var preprocessInput pMatchingInput
	preprocessInput.fc_map = make(map[IDType]*[2]IDType)
	preprocessInput.c_map = make(map[IDType]bool)
	preprocessInput.preference_map = make(map[IDType]map[int64][]IDType)

	for _, value := range mI.faculty {
		preprocessInput.faculty_ids = append(preprocessInput.faculty_ids, value.UserID)
		preprocessInput.fc_map[value.UserID] = &([2]IDType{-1, -1})
	}
	for _, value := range mI.course_s {
		preprocessInput.course_s_ids = append(preprocessInput.course_s_ids, value.CourseSemID)
		preprocessInput.c_map[value.CourseSemID] = false
	}
	for _, value := range mI.preferences {
		userId, courseSemId, level := value.UserID, value.CourseSemID, value.PreferenceLevel
		preprocessInput.preferences = append(preprocessInput.preferences, value)

		_, exists := preprocessInput.preference_map[userId]
		if exists {
			preprocessInput.preference_map[userId][level] = append(preprocessInput.preference_map[userId][level], courseSemId)
		} else {
			preprocessInput.preference_map[userId] = make(map[int64][]IDType)
			preprocessInput.preference_map[userId][level] = append(preprocessInput.preference_map[userId][level], courseSemId)
		}
	}

	return preprocessInput, nil
}

func matchingEngine(pI pMatchingInput, matchingIter IDType) (Matching, error) {
	rand.NewSource(time.Now().UnixNano())
	rand.Shuffle(len(pI.faculty_ids), func(i, j int) {
		pI.faculty_ids[i], pI.faculty_ids[j] = pI.faculty_ids[j], pI.faculty_ids[i]
	})
	for iter := 0; iter < 2; iter++ {
		for _, value := range pI.faculty_ids {
			for i := PreferenceLevelGreen; i < PreferenceLevelEnd; i++ {
				_, exists := pI.preference_map[value][i]
				if !exists {
					continue
				}

				length, j, flag := len(pI.preference_map[value][i]), 0, false
				for j < length {
					courseSemId := pI.preference_map[value][i][j]
					j++
					if !pI.c_map[courseSemId] { // If course is not assigned
						if pI.fc_map[value][0] == -1 {
							pI.fc_map[value][0] = courseSemId
						} else {
							pI.fc_map[value][1] = courseSemId
						}
						pI.c_map[courseSemId] = true
						flag = true
						break
					}
				}

				pI.preference_map[value][i] = slices.Delete(pI.preference_map[value][i], 0, j)
				if flag {
					break
				}
			}
		}
	}

	// TODO: Big todo! Can make or break the algo
	// for f_id, c_id := range pI.fc_map {
	// }

	var matching Matching
	for key, value := range pI.fc_map {
		matchingElement := MatchingElement{
			MatchingID:          -1,
			MatchingIterationID: matchingIter,
			UserID:              key,
			CourseSemID:         value[0],
			MatchingScore:       1.0,
		}

		matching.Matchings = append(matching.Matchings, matchingElement)

		if value[1] != -1 {
			matchingElement2 := MatchingElement{
				MatchingID:          -1,
				MatchingIterationID: matchingIter,
				UserID:              key,
				CourseSemID:         value[1],
				MatchingScore:       1.0,
			}

			matching.Matchings = append(matching.Matchings, matchingElement2)
		}
	}

	return matching, nil
}

func runMatching(mI MatchingInput, matchingIter IDType) (Matching, error) {
	preprocessInput, err := preprocessMatchingInput(mI)
	if err != nil {
		return Matching{}, err // TODO: Custom error type
	}

	matching, err := matchingEngine(preprocessInput, matchingIter)
	if err != nil {
		return Matching{}, err // TODO: Custom error type
	}

	return matching, nil
}
