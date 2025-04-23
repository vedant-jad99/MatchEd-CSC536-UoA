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

var RemoveCourseSemesters = func(courseID uint) error {
	return db.Where("course_id = ?", courseID).Delete(&CourseSemester{}).Error
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

var BulkUpsertCourseSemester = func(courseSemesters []CourseSemester) ([]CourseSemester, error) {
	err := db.Save(&courseSemesters).Error
	return courseSemesters, err
}

var BulkDeleteCourseSemesters = func(ids []uint) error {
	var courseSemesters []CourseSemester
	err := db.Where("id IN ?", ids).Delete(&courseSemesters).Error
	return err
}
var DeleteCourseSemestersByCourseID = func(courseID uint) error {
	var courseSemesters []CourseSemester
	err := db.Where("course_id = ?", courseID).Delete(&courseSemesters).Error
	return err
}

var FetchAllCourses = func() ([]Course, error) {
	var courses []Course
	err := db.Find(&courses).Error
	return courses, err
}
var FetchCourseById = func(id uint) (Course, error) {
	var course Course
	err := db.First(&course, id).Error
	return course, err
}
var UpsertCourse = func(course Course) (Course, error) {
	txn := db.Save(&course)
	return course, txn.Error
}
var DeleteCourse = func(id uint) error {
	var course Course
	if err := RemoveCourseSemesters(id); err != nil {
		return err
	}
	txn := db.Delete(&course, id)
	return txn.Error
}
