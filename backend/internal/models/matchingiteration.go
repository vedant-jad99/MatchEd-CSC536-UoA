package models

import "time"

// MatchingIteration model represents a round of course matching.
// id | triggered_by | status | created_at | updated_at
type MatchingIteration struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	TriggeredBy uint      `json:"triggered_by" gorm:"column:triggered_by;not null"`
	Status      string    `json:"status" gorm:"column:status;type:varchar;not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at;not null"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at;not null"`
}

func (MatchingIteration) TableName() string {
	return `"match_schema"."matching_iterations"`
}
