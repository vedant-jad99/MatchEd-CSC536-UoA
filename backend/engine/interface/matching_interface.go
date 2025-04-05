package matching

import (
	"fmt"
)

type IDType int64

const (
	PreferenceLevelInvalid int64 = -1
	PreferenceLevelGreen   int64 = 1
	PreferenceLevelYellow  int64 = 2
	PreferenceLevelRed     int64 = 3
	PreferenceLevelEnd     int64 = 4
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
	PreferenceID    IDType
	UserID          IDType
	CourseID        IDType
	Semester        string
	PreferenceLevel int64
}

type CourseSems struct {
	CourseSemID    IDType
	CourseID       IDType
	Semester       string
	MandatoryLevel int    // TODO: Custom mandatory type
	TimeSlot       string // TODO: Custom time slot type or timestamp?
}

type Faculty struct {
	UserID IDType
}

type MatchingInput struct {
	faculty     []Faculty
	course_s    []CourseSems
	preferences []Preferences
}

type MatchingElement struct {
	MatchingID          IDType
	MatchingIterationID IDType
	UserID              IDType
	CourseID            IDType
	MatchingScore       float64
}

type Matching struct {
	matchings []MatchingElement
}

type MatchingQueue struct {
	m_IterQ []IDType
}

type MatchingInterface struct {
	m_Queue     MatchingQueue
	m_StatusMap map[IDType]string // TODO: Custom status type
}

func NewMatchingInterface() *MatchingInterface {
	return &MatchingInterface{m_Queue: MatchingQueue{m_IterQ: []IDType{}}, m_StatusMap: make(map[IDType]string)}
}

func (m_Interface *MatchingInterface) TriggerMatching(matchingIterID IDType) error {
	if matchingIterID == -1 {
		// TODO: Return custom error type
		return fmt.Errorf("Invalid matching iteration ID")
	}

	m_Interface.m_Queue.m_IterQ = append(m_Interface.m_Queue.m_IterQ, matchingIterID)
	m_Interface.m_StatusMap[matchingIterID] = MatchingIterationStatusInitialized
	return nil
}

// TODO: Function to fetch the input from the database to run the engine
func (m_Interface *MatchingInterface) GetInput() error {
	return nil
}

// TODO: Store the matching result to database
func (m_Interface *MatchingInterface) StoreMatchingResult(matching Matching) error {
	return nil
}

func (m_Interface *MatchingInterface) UpdateStatus(matchingIterID IDType) (string, error) {
	if matchingIterID == -1 {
		// TODO: Return custom error type
		return "Error", fmt.Errorf("Invalid matching iteration ID")
	}
	/*
		TODO: Some code goes here. Get the status
	*/

	currentStatus := m_Interface.m_StatusMap[matchingIterID]
	// update status in db

	return currentStatus, nil
}
