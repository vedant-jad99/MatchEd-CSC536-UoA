package models

import (
	"time"
)

// Auth model tracks authentication sessions.
type Auth struct {
	//gorm.Model           // Adds ID, CreatedAt, UpdatedAt, DeletedAt
	ID        uint       `gorm:"primaryKey"`
	UserID    uint       `gorm:"column:user_id;not null"`
	StartTime time.Time  `gorm:"column:start_time;not null"`
	EndTime   *time.Time `gorm:"column:end_time"`
	AuthToken string     `gorm:"column:auth_token;type:varchar;not null;unique"`
}

func (Auth) TableName() string {
	return `"match_schema"."auth"`
}
