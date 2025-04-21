package models

// CourseSemester model represents the relationship between courses and semesters.
type CourseSemester struct {
	//gorm.Model        // Adds ID, CreatedAt, UpdatedAt, DeletedAt
	ID             uint   `json:"id" gorm:"primaryKey"`
	CourseID       uint   `json:"course_id" gorm:"column:course_id;not null"`
	Semester       string `json:"semester" gorm:"type:varchar;not null"`
	MandatoryLevel string `json:"mandatory_level" gorm:"column:mandatory_level;type:varchar;"`
	Timeslot       string `json:"timeslot" gorm:"type:varchar;"`
}

func (CourseSemester) TableName() string {
	return `"match_schema"."course_sem"`
}

func FetchAllCourseSemesters() ([]CourseSemester, error) {
	var courseSemesters []CourseSemester
	txn := DB.Find(&courseSemesters)

	return courseSemesters, txn.Error
}