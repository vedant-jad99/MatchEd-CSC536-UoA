package models

type MatchingIteration struct {
	Id          uint   `json:"id" gorm:"primaryKey"`
	TriggeredBy uint   `json:"triggered_by"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func (MatchingIteration) TableName() string {
	return "match_schema.matching_iterations"
}

type Matching struct {
	Id                  uint    `json:"id" gorm:"primaryKey"`
	MatchingIterationId uint    `json:"matching_iteration_id"`
	UserId              uint    `json:"user_id"`
	CourseSemId         uint    `json:"course_sem_id"`
	Score               float64 `json:"score"`
}

func (Matching) TableName() string {
	return "match_schema.matchings"
}
