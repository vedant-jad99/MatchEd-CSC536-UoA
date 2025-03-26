package models

import "gorm.io/gorm"

// Matching model represents the results of a matching iteration.
type Matching struct {
	gorm.Model               // Adds ID, CreatedAt, UpdatedAt, DeletedAt
	ID                  uint `gorm:"primaryKey"`
	MatchingIterationID uint `gorm:"not null"`
	UserID              uint `gorm:"not null"`
	CourseID            uint `gorm:"not null"`
}
