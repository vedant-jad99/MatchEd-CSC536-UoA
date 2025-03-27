package models

import (
	"time"
)

// MatchingIteration model represents a round of course matching.
type MatchingIteration struct {
	//gorm.Model           // Adds ID, CreatedAt, UpdatedAt, DeletedAt
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"not null"`
}
