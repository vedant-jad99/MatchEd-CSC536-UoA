package matching

import (
	"math/rand"
	"time"
)

type pMatchingInput struct {
	faculty_ids    []IDType
	course_s_ids   []IDType
	preferences    []Preferences
	fc_map         map[IDType]IDType
	c_map          map[IDType]bool
	preference_map map[IDType]map[int64]IDType //FacultyID --> preference (3,2,1) --> courseSemID
}

// TODO: Get the input from the database. Communicates to the matching interface
func GetInput() (MatchingInput, error) {
	// TODO: Custom error type
	return MatchingInput{}, nil
}

func StartMatching(matchingIter IDType) (Matching, error) {
	preferences, err := GetInput()
	if err != nil {
		return Matching{}, err // TODO: Custom error type?
	}

	//UpdateMatchingIterStatus(matchingIter);
	matching, err := runMatching(preferences, matchingIter)
	if err != nil {
		return Matching{}, err
	}
	return matching, nil
}

func preprocessMatchingInput(mI MatchingInput) (pMatchingInput, error) {
	var preprocessInput pMatchingInput
	preprocessInput.fc_map = make(map[IDType]IDType)
	preprocessInput.c_map = make(map[IDType]bool)
	preprocessInput.preference_map = make(map[IDType]map[int64]IDType)

	for _, value := range mI.faculty {
		preprocessInput.faculty_ids = append(preprocessInput.faculty_ids, value.UserID)
		preprocessInput.fc_map[value.UserID] = -1
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
			preprocessInput.preference_map[userId][level] = courseSemId
		} else {
			preprocessInput.preference_map[userId] = make(map[int64]IDType)
			preprocessInput.preference_map[userId][level] = courseSemId
		}
	}

	return preprocessInput, nil
}

func matchingEngine(pI pMatchingInput, matchingIter IDType) (Matching, error) {
	rand.NewSource(time.Now().UnixNano())
	rand.Shuffle(len(pI.faculty_ids), func(i, j int) {
		pI.faculty_ids[i], pI.faculty_ids[j] = pI.faculty_ids[j], pI.faculty_ids[i]
	})
	for _, value := range pI.faculty_ids {
		for i := PreferenceLevelGreen; i < PreferenceLevelEnd; i++ {
			courseSemId, exists := pI.preference_map[value][i]
			if !exists {
				continue
			}
			if !pI.c_map[courseSemId] { // If course is not assigned
				pI.fc_map[value] = courseSemId
				pI.c_map[courseSemId] = true
				break
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
			CourseSemID:         value,
			MatchingScore:       1.0,
		}

		matching.matchings = append(matching.matchings, matchingElement)
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
