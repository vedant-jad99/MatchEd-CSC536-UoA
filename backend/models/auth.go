package models

import (
	"time"

	"gorm.io/gorm"
)

// Auth model tracks authentication sessions.
type Auth struct {
	gorm.Model           // Adds ID, CreatedAt, UpdatedAt, DeletedAt
	ID         uint      `gorm:"primaryKey"`
	UserID     uint      `gorm:"not null"`
	StartTime  time.Time `gorm:"not null"`
	EndTime    *time.Time
	AuthToken  string `gorm:"type:varchar;not null;unique"`
}
