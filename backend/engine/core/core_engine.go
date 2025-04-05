package matching

import (
	"time"
)

type preferenceMapValue	struct {
	courseID		IDType
	preferenceLevel	int64
}

type pMatchingInput struct {
	faculty_ids 	[]IDType
	course_ids		[]IDType
	preferences		[]Preferences
	fc_map 			map[IDType]any;
	c_map			map[IDType]bool;
	preference_map	map[IDType]map[int64]IDType;	//FacultyID --> preference (3,2,1) --> courseID
}


// TODO: Get the input from the database. Communicates to the matching interface
func GetInput() (MatchingInput, error) {
	// TODO: Custom error type
}

func StartMatching(matchingIter IDType) (Matching, error) {
	preferences, err := GetInput();
	if err != nil {
		return nil, err; // TODO: Custom error type?
	}

	//UpdateMatchingIterStatus(matchingIter);
	matching, err := runMatching(preferences, matchingIter);
}

func preprocessMatchingInput(mI MatchingInput) (pMatchingInput, error) {
	var preprocessInput pMatchingInput;

	for _, value := range mI.faculty {
		preprocessInput.faculty_ids 	= append(preprocessInput.faculty_ids, value);
		preprocessInput.fc_map[value] 	= nil;
	}
	for _, value := range mI.course_s {
		preprocessInput.course_ids 		= append(preprocessInput.course_ids, value);
		preprocessInput.c_map[value]	= false;
	}
	for _, value := range mI.preferences {
		userId, courseId, level 	:= value.UserID, value.CourseID, value.PreferenceLevel;
		preprocessInput.preferences  = append(preprocessInput.preferences, value);

		_, exists := preprocessingInput.preference_map[userId];
		if exists {
			preprocessingInput.preference_map[userId][level] = courseId; 
		} else {
			preprocessingInput.preference_map[userId] 		 = make(map[int64]IDType);
			preprocessingInput.preference_map[userId][level] = courseId;
		}
	}

	return preprocessInput, nil
}

func matchingEngine(pI pMatchingInput, matchingIter IDType) (Matching, error) {
	matchingId = -1;
	rand.Seed(time.Now().UnixNano());	// Randomize the faculty array
	rand.Shuffle(len(pI.faculty_ids), func(i, j, IDType) {
		pI.faculty_ids[i], pI.faculty_ids[j] =pI.faculty_ids[j], pI.faculty_ids[i]; 
	});

	for _, value := range pI.faculty_ids {
		for i := PreferenceLevelGreen; i < PreferenceLevelEnd; i++ {
			courseId, exists := pI.preference_map[facultyID][i];
			if !exists {
				continue;
			}
			if !pI.c_map[courseId] { // If course is not assigned
				pI.fc_map[value] = courseId;
				pI.c_map[courseId] = true;
				break;
			}
		}
	}

	// TODO: Big todo! Can make or break the algo
	for f_id, c_id := range pI.fc_map {
	}

	var matching Matching;
	for key, value := range pI.fc_map {
		matchingElement := MatchingElement{
			MatchingID: -1,
			MatchingIterationID: matchingIter,
			UserID: key,
			CourseID: value,
			MatchingScore: 1.0					// TODO: Metric to assign score. Between 0 - 1
		}

		matching.matchings = append(matching.matchings, matchingElement);
	}

	return matching, nil
}

func runMatching(mI MatchingInput, matchingIter IDType) (Matching, error) {
	preprocessInput, err := preprocessMatchingInput(mI)
	if err != nil {
		return nil, err	// TODO: Custom error type
	}

	matching, err = matchingEngine(preprocessInput, matchingIter)
	if err != nil {
		return nil, err	// TODO: Custom error type
	}

	return matching, nil
}
