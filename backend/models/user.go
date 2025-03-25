package models

import "gorm.io/gorm"

type User struct {
	gorm.Model        // Adds ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string `json:"name"`
	Email      string `json:"email"`
}
