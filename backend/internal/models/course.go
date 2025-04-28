package models

import (
	"errors"
)

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

// Add CourseSemester by Course name
func AddCourseSemesterByName(courseName, semester, mandatoryLevel, timeslot string) (*CourseSemester, error) {
	// Check if the course exists
	var course Course
	if err := db.Where("name = ?", courseName).First(&course).Error; err != nil {
		// Course doesn't exist, create it
		course = Course{Name: courseName}
		if err := db.Create(&course).Error; err != nil {
			return nil, err
		}
	}

	// Create the CourseSemester
	courseSemester := CourseSemester{
		CourseID:       course.ID,
		Semester:       semester,
		MandatoryLevel: mandatoryLevel,
		Timeslot:       timeslot,
	}

	if err := db.Create(&courseSemester).Error; err != nil {
		return nil, err
	}
	return &courseSemester, nil
}

// Edit CourseSemester by Course name (update Course)
func EditCourseSemesterByName(oldCourseName, newCourseName, semester, mandatoryLevel, timeslot string) (*CourseSemester, error) {
	// Find CourseSemester by old Course name
	var courseSemester CourseSemester
	if err := db.Where("course_id IN (SELECT id FROM match_schema.courses WHERE name = ?)", oldCourseName).First(&courseSemester).Error; err != nil {
		return nil, errors.New("course semester not found")
	}

	// Find the new Course by name
	var newCourse Course
	if err := db.Where("name = ?", newCourseName).First(&newCourse).Error; err != nil {
		// If the course does not exist, create it
		newCourse = Course{Name: newCourseName}
		if err := db.Create(&newCourse).Error; err != nil {
			return nil, err
		}
	}

	// Update CourseSemester with the new Course ID
	courseSemester.CourseID = newCourse.ID
	courseSemester.Semester = semester
	courseSemester.MandatoryLevel = mandatoryLevel
	courseSemester.Timeslot = timeslot
	if err := db.Save(&courseSemester).Error; err != nil {
		return nil, err
	}
	return &courseSemester, nil
}

// Remove CourseSemester by Course name
func RemoveCourseSemesterByName(courseName string) error {
	// Find the CourseSemester by course name
	var courseSemester CourseSemester
	if err := db.Where("course_id IN (SELECT id FROM match_schema.courses WHERE name = ?)", courseName).First(&courseSemester).Error; err != nil {
		return errors.New("course semester not found")
	}

	// Remove the CourseSemester
	if err := db.Delete(&courseSemester).Error; err != nil {
		return err
	}
	return nil
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
