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
	// TODO: Make the value a priority queue/heap holding the preferenceMapValues
	preference_map	map[IDType]preferenceMapValue;	//FacultyID --> {courseID, preferenceLevel}
}


// TODO: Get the input from the database. Communicates to the matching interface
func GetInput() (MatchingInput, error) {
	// TODO: Custom error type
}

func StartMatching(matchingIter MatchingIteration) (Matching, error) {
	preferences, err := GetInput();
	if err != nil {
		return nil, err; // TODO: Custom error type?
	}

	UpdateMatchingIterStatus(matchingIter);
	matching, err := runMatching(preferences);
}

func preprocessMatchingInput(mI MatchingInput) (pMatchingInput, error) {
	var preprocessInput pMatchingInput;

	for _, value := range mI.faculty {
		preprocessInput.faculty_ids 	= append(preprocessInput.faculty_ids, value);
		preprocessInput.fc_map[value] 	= nil;
	}
	for _, value := range mI.course_s {
		preprocessInput.course_ids 	= append(preprocessInput.course_ids, value);
		preprocessInput.c_map[value]= false;
	}
	for _, value := range mI.preferences {
		userId, courseId, level 				   := value.UserID, value.CourseID, value.PreferenceLevel;
		preprocessInput.preferences 				= append(preprocessInput.preferences, value);
		preprocessingInput.preference_map[userId] 	= preferenceMapValue{courseID: courseId, preferenceLevel: level};
	}

	return preprocessInput, nil
}

func matchingEngine(pI pMatchingInput) (Matching, error) {
	matchingId = generateMatchingID()	// TODO: Implement function to generate a random matching ID
	rand.Seed(time.Now().UnixNano());	// Randomize the faculty array
	rand.Shuffle(len(pI.faculty_ids), func(i, j, IDType) {
		pI.faculty_ids[i], pI.faculty_ids[j] =pI.faculty_ids[j], pI.faculty_ids[i]; 
	});

	for _, value := range pI.faculty_ids {
		for isEmpty(pI.preference_map, value) { //TODO: Implement the isEmpty helper.
			highestPrefVal := getHighestPreferenceValue(pI.preference_map, value); // TODO: Write this function OR make priority queue
			if !pI.c_map[highestPrefVal.courseID] { // If course is not assigned
				pI.fc_map[value] = highestPrefVal.courseID;
				break;
			}
		}
	}

	for f_id, c_id := range pI.fc_map {
		if c_id == nil {
			/* TODO: Assign an unassigned course in the `yellow` category. */
		}
	}

	var matching Matching;
	for key, value := range pI.fc_map {
		matchingElement := MatchingElement{
			MatchingID: generateMatchingID(),	// TODO: Helper function to generate a matching id
			MatchingIterationID: -1,			// TODO: Design question? Pass the iteration id or add in the interface? I vote passing
			UserID: key,
			CourseID: value.courseId,
			MatchingScore: 1.0					// TODO: Metric to assign score. Between 0 - 1
		}

		matching.matchings = append(matching.matchings, matchingElement);
	}

	return matching, nil
}

func runMatching(mI MatchingInput) (Matching, error) {
	preprocessInput, err := preprocessMatchingInput(mI)
	if err != nil {
		return nil, err	// TODO: Custom error type
	}

	matching, err = matchingEngine(preprocessInput)
	if err != nil {
		return nil, err	// TODO: Custom error type
	}

	return matching, nil
}
