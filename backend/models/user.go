package models

import "gorm.io/gorm"

// User model represents application users.
type User struct {
	gorm.Model        // Adds ID, CreatedAt, UpdatedAt, DeletedAt
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"type:varchar"`
	Email      string `gorm:"type:varchar"`
}
