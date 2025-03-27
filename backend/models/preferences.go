package models

// Preferences model represents user preferences for courses.
type Preferences struct {
	//gorm.Model      // Adds ID, CreatedAt, UpdatedAt, DeletedAt
	ID       uint `gorm:"primaryKey"`
	UserID   uint `gorm:"not null"`
	CourseID uint `gorm:"not null"`
	Priority int  `gorm:"not null"`
}
