package models

import (
	"gorm.io/gorm"
)

// Role model represents user roles.
type Role struct {
	gorm.Model        // Adds ID, CreatedAt, UpdatedAt, DeletedAt
	UserID     uint   `gorm:"primaryKey"`
	Role       string `gorm:"type:varchar;not null"`
}
