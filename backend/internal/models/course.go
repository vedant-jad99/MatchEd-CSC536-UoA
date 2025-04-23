package models

// TODO change time reprentation of days and times the course section is scheduled for
type Course struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Name   string `json:"name" gorm:"type:varchar;not null"`
	Number string `json:"number" gorm:"type:varchar;not null"`
	Campus string `json:"campus" gorm:"type:varchar;not null"`
}

func (Course) TableName() string {
	return "match_schema.courses"
}

// CourseSemester model represents an instance of a course
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

var FetchCourseSemesters = func() ([]CourseSemester, error) {
	var courseSemesters []CourseSemester
	err := db.Find(&courseSemesters).Error
	return courseSemesters, err
}
var FetchAllCourseSemesters = func() ([]CourseSemester, error) {
	var courseSemesters []CourseSemester
	err := db.Find(&courseSemesters).Error
	return courseSemesters, err
}

var FetchCourseSemester = func(id uint) (CourseSemester, error) {
	var cs CourseSemester
	err := db.First(&cs, id).Error
	return cs, err
}

var AddCourseSemester = func(courseID uint, semester string) error {
	cs := CourseSemester{CourseID: courseID, Semester: semester}
	return db.Create(&cs).Error
}

var RemoveCourseSemester = func(id uint) error {
	return db.Delete(&CourseSemester{}, id).Error
}

var UpdateCourseSemester = func(id uint, courseID uint, semester string, mandatoryLevel string, timeslot string) error {
	updates := map[string]interface{}{
		"course_id":       courseID,
		"semester":        semester,
		"mandatory_level": mandatoryLevel,
		"timeslot":        timeslot,
	}
	return db.Model(&CourseSemester{}).Where("id = ?", id).Updates(updates).Error
}
