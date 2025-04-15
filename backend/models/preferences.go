package models

// Preferences model represents user preferences for courses.
type Preferences struct {
	//gorm.Model      // Adds ID, CreatedAt, UpdatedAt, DeletedAt
	ID               uint   `gorm:"primaryKey"`
	Semester         string `gorm:"varchar;not null"`
	UserID           uint   `gorm:"column:user_id; not null"`
	CourseSemesterID uint   `gorm:"column:course_sem_id; not null"`
	PreferenceLevel  uint   `gorm:"column:preference_level;not null"`
}

func (Preferences) TableName() string {
	return `"match_schema"."preferences"`
}
