package matching

import (
	"math/rand"
	"slices"
	"sort"
	"time"
)

type pData struct {
	courseSID  IDType
	prefWeight int64
}

type pMatchingInput struct {
	faculty_ids    []IDType
	course_s_ids   []IDType
	preferences    []Preferences
	fc_map         map[IDType]*[4]IDType
	c_map          map[IDType]bool
	f2rc_map       map[IDType]int               //FacultyID --> number of required courses to teach
	preference_map map[IDType]map[int64][]pData //FacultyID --> preference (3,2,1) --> (courseSemID, weight)[]
}

func StartMatching(matchingIter IDType, input MatchingInput) (Matching, error) {
	matching, err := runMatching(input, matchingIter)
	if err != nil {
		return Matching{}, err
	}
	return matching, nil
}

func preprocessMatchingInput(mI MatchingInput) pMatchingInput {
	var preprocessInput pMatchingInput
	preprocessInput.fc_map = make(map[IDType]*[4]IDType)
	preprocessInput.c_map = make(map[IDType]bool)
	preprocessInput.f2rc_map = make(map[IDType]int)
	preprocessInput.preference_map = make(map[IDType]map[int64][]pData)

	for _, value := range mI.Faculty {
		preprocessInput.faculty_ids = append(preprocessInput.faculty_ids, value.UserID)
		preprocessInput.fc_map[value.UserID] = &([4]IDType{-1, -1, -1, -1})
		preprocessInput.f2rc_map[value.UserID] = value.NumReqCourses
	}
	for _, value := range mI.Course_s {
		preprocessInput.course_s_ids = append(preprocessInput.course_s_ids, value.CourseSemID)
		preprocessInput.c_map[value.CourseSemID] = false
	}
	for _, value := range mI.Preferences {
		userId, courseSemId, level, weight := value.UserID, value.CourseSemID, value.PreferenceLevel, value.PreferenceWeight
		preprocessInput.preferences = append(preprocessInput.preferences, value)
		data := pData{courseSemId, weight}

		_, exists := preprocessInput.preference_map[userId]
		if exists {
			preprocessInput.preference_map[userId][level] = append(preprocessInput.preference_map[userId][level], data)
			/* Store in sorted order */
			sort.Slice(preprocessInput.preference_map[userId][level], func(i, j int) bool {
				return preprocessInput.preference_map[userId][level][i].prefWeight >
					preprocessInput.preference_map[userId][level][j].prefWeight
			})
		} else {
			preprocessInput.preference_map[userId] = make(map[int64][]pData)
			preprocessInput.preference_map[userId][level] = append(preprocessInput.preference_map[userId][level], data)
		}
	}

	return preprocessInput
}

func matchingEngine(pI pMatchingInput, matchingIter IDType) (Matching, error) {
	rand.NewSource(time.Now().UnixNano())
	rand.Shuffle(len(pI.faculty_ids), func(i, j int) {
		pI.faculty_ids[i], pI.faculty_ids[j] = pI.faculty_ids[j], pI.faculty_ids[i]
	})
	for iter := 0; iter < 4; iter++ {
		for _, value := range pI.faculty_ids {
			if iter >= pI.f2rc_map[value] {
				continue
			}
			for i := PreferenceLevelGreen; i < PreferenceLevelEnd; i++ {
				_, exists := pI.preference_map[value][i]
				if !exists {
					continue
				}

				length, j, flag := len(pI.preference_map[value][i]), 0, false
				for j < length {
					courseSemId := pI.preference_map[value][i][j].courseSID
					j++

					// Check if the course exists in the course map
					_, ok := pI.c_map[courseSemId];
					if !ok {
						continue
					}
					if !pI.c_map[courseSemId] { // If course is not assigned
						for k := 0; k < 4; k++ {
							if pI.fc_map[value][k] == -1 {
								pI.fc_map[value][k] = courseSemId
								break
							}
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
	preprocessInput := preprocessMatchingInput(mI)
	matching, err := matchingEngine(preprocessInput, matchingIter)
	if err != nil {
		return Matching{}, err // TODO: Custom error type
	}

	return matching, nil
}
