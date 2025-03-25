package models

import "gorm.io/gorm"

// TODO add time preference, courseload, time availability
type Person struct {
	gorm.Model           // Adds ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Courses    []*Course `json:"courses"`

	// courseID -> preference level
	Preferences map[uint]int `json:"preferences"`
}
