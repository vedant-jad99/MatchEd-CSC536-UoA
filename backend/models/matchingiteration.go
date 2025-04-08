package models

import "time"

// MatchingIteration model represents a round of course matching.
// id | triggered_by | status | created_at | updated_at
type MatchingIteration struct {
	ID          uint      `gorm:"primaryKey"`
	TriggeredBy uint      `gorm:"column:triggered_by;not null"`
	Status      string    `gorm:"column:status;type:varchar;not null"`
	CreatedAt   time.Time `gorm:"column:created_at;not null"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null"`
}

func (MatchingIteration) TableName() string {
	return `"match_schema"."matching_iteration"`
}
