package matching

import (
	"time"
)

type IDType	int64

type Preferences struct {
	PreferenceID	IDType
	UserID			IDType
	CourseID 		IDType
	Semester		string
	PreferenceLevel	int64
}

type CourseSems struct {
	CourseSemID 	IDType
	CourseID		IDType
	Semester		string
	MandatoryLevel	int		// TODO: Custom mandatory type
	TimeSlot		string	// TODO: Custom time slot type or timestamp?
}

type Faculty struct {
	UserID			IDType
}

type MatchingInput struct {
	faculty 	[]Faculty
	course_s	[]CourseSems
	preferences	[]Preferences
}

type MatchingElement struct {
	MatchingID			IDType
	MatchingIterationID	IDType
	UserID				IDType
	CourseID			IDType
	MatchingScore		float64
}

type Matching struct {
	matchings []MatchingElement
}

type MatchingQueue struct {
	m_IterQ []MatchingIteration
}

type MatchingInterface struct {
	m_Queue 	MatchingQueue
	m_StatusMap	map[MatchingIteration]string // TODO: Custom status type
}

func NewMatchingInterface() *MatchingInterface {
	return &MatchingInterface{m_Queue: MatchingQueue{m_IterQ: []int64}};
}

func (m_Interface *MatchingInterface) TriggerMatching(matchingIterID MatchingIteration) error {
	if matchingIterID == -1 {
		// TODO: Return custom error type
		return nil
	}

	m_Interface.m_Queue.m_IterQ = append(m_Interface.m_Queue.m_IterQ, matchingIterID);
	m_Interface.m_StatusMap[matchingIterID] = "Initialized" // TODO: Custom status type
	return nil
}

// TODO: Function to fetch the input from the database to run the engine
func (m_Interface *MatchingInterface) GetInput() error {

}

// TODO: Store the matching result to database
func (m_Interface *MatchingInterface) StoreMatchingResult() (Matching, error) {

}

func (m_Interface *MatchingInterface) UpdateStatus(matchingIterID MatchingIteration) (string, error) {
	if matchingIterID == -1 {
		// TODO: Return custom error type
		return "Error", nil
	}
	/*
		TODO: Some code goes here. Get the status
	*/

	m_Interace.m_StatusMap[matchingIterID] = status; // TODO: Status got above
}
