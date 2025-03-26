package models

import (
	"gorm.io/gorm"
)

// Course model represents a course entry.
// TODO change time reprentation of days and times the course section is scheduled for
type Course struct {
	gorm.Model        // Adds ID, CreatedAt, UpdatedAt, DeletedAt
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"type:varchar;not null"`
	Type       string `gorm:"type:varchar;not null"`
}

// instantiate new user to database
func NewUser(db *gorm.DB, name, email string) (User, error) {
	user := User{Name: name, Email: email}
	err := db.Create(&user).Error
	return user, err
}
