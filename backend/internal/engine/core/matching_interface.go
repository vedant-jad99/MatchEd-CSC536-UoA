package matching

import (
	"fmt"
)

type IDType int64

const (
	PreferenceLevelInvalid int64 = -1
	PreferenceLevelGreen   int64 = 1
	PreferenceLevelYellow  int64 = 2
	PreferenceLevelEnd     int64 = 3
	PreferenceLevelRed     int64 = 4
)

const (
	MatchingIterationStatusInitialized string = "INITIALIZED"
	MatchingIterationStatusRunning     string = "RUNNING"
	MatchingIterationStatusCompleted   string = "COMPLETED"
	MatchingIterationStatusError       string = "ERROR"
	MatchingIterationStatusCancelled   string = "CANCELLED"
	MatchingIterationStatusPaused      string = "PAUSED"
)

type Preferences struct {
	PreferenceID    	IDType
	UserID          	IDType
	CourseSemID     	IDType
	Semester        	string
	PreferenceLevel 	int64
	PreferenceWeight	int64
}

type CourseSems struct {
	CourseSemID	IDType
	CourseID	IDType
	Semester	string
	Section		string
	MandatoryLevel	int    // TODO: Custom mandatory type
	TimeSlot	string // TODO: Custom time slot type or timestamp?
}

type Faculty struct {
	UserID 		IDType
	NumReqCourses	int
}

type MatchingInput struct {
	faculty     []Faculty
	course_s    []CourseSems
	preferences []Preferences
}

type MatchingElement struct {
	MatchingID          IDType  `json:"matching_id"`
	MatchingIterationID IDType  `json:"matching_iteration_id"`
	UserID              IDType  `json:"user_id"`
	CourseSemID         IDType  `json:"course_sem_id"`
	MatchingScore       float64 `json:"matching_score"`
}

type Matching struct {
	Matchings []MatchingElement
}

type MatchingInterface struct {
	m_StatusMap map[IDType]string
}

func NewMatchingInterface() *MatchingInterface {
	return &MatchingInterface{m_StatusMap: make(map[IDType]string)}
}

func (m_Interface *MatchingInterface) TriggerMatching(matchingIterID IDType, input MatchingInput) (Matching, error) {
	if matchingIterID == -1 {
		// TODO: Return custom error type
		return Matching{}, fmt.Errorf("invalid matching iteration ID")
	}

	m_Interface.m_StatusMap[matchingIterID] = MatchingIterationStatusInitialized
	matching, err := StartMatching(matchingIterID, input);
	if err != nil {
		m_Interface.m_StatusMap[matchingIterID] = MatchingIterationStatusError;
		return Matching{}, err
	}

	m_Interface.m_StatusMap[matchingIterID] = MatchingIterationStatusCompleted; 
	return matching, nil
}
