package models

import "gorm.io/gorm"

// CourseSemester model represents the relationship between courses and semesters.
type CourseSemester struct {
	gorm.Model        // Adds ID, CreatedAt, UpdatedAt, DeletedAt
	ID         uint   `gorm:"primaryKey"`
	CourseID   uint   `gorm:"not null"`
	Semester   string `gorm:"type:varchar;not null"`
}
