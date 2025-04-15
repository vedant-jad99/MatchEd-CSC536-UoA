package models

// CourseSemester model represents the relationship between courses and semesters.
type CourseSemester struct {
	//gorm.Model        // Adds ID, CreatedAt, UpdatedAt, DeletedAt
	ID             uint   `gorm:"primaryKey"`
	CourseID       uint   `gorm:"column:course_id;not null"`
	Semester       string `gorm:"type:varchar;not null"`
	MandatoryLevel string `gorm:"column:mandatory_level;type:varchar;"`
	Timeslot       string `gorm:"type:varchar;"`
}

func (CourseSemester) TableName() string {
	return `"match_schema"."course_sem"`
}
