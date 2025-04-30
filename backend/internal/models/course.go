package models

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
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
	Section        string `json:"section" gorm:"type:varchar;"`
	Timeslot       string `json:"timeslot" gorm:"type:varchar;"`
}

func (CourseSemester) TableName() string {
	return `"match_schema"."course_sem"`
}

func CountSemestersForCourse(courseID uint) (int64, error) {
	var count int64
	err := db.Model(&CourseSemester{}).Where("course_id = ?", courseID).Count(&count).Error
	return count, err
}

// adds a new course semester with the course that matches the name
// or updates the coursesemester with the course that matches the name
// creates a new course if no course matches
func UpsertCourseAndCourseSemester(courseSemesterID uint, courseName, courseNumber string) (*CourseSemester, error) {
	var course Course
	var err error

	// Step 1: Find or create the Course based on Number
	err = db.Where("number = ?", courseNumber).First(&course).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		course = Course{
			Name:   courseName,
			Number: courseNumber,
			Campus: "TBD",
		}
		if err = db.Create(&course).Error; err != nil {
			return nil, fmt.Errorf("failed to create course: %w", err)
		}
	} else if err != nil {
		return nil, fmt.Errorf("failed to find course: %w", err)
	}

	// Step 2: Prepare the CourseSemester struct
	courseSem := CourseSemester{
		CourseID: course.ID,
		Semester: "TBD",
		Section:  "1",
	}

	// Step 3: Try to find existing CourseSemester by ID
	var existingCourseSem CourseSemester
	err = db.Where("id = ?", courseSemesterID).First(&existingCourseSem).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Doesn't exist — insert new
		if err = db.Create(&courseSem).Error; err != nil {
			return nil, fmt.Errorf("failed to create course_semester: %w", err)
		}
		return &courseSem, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to find course_semester: %w", err)
	}

	// Exists — update fields
	if err = db.Model(&existingCourseSem).Updates(courseSem).Error; err != nil {
		return nil, fmt.Errorf("failed to update course_semester: %w", err)
	}

	return &existingCourseSem, nil
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

/*
var RemoveCourseSemester = func(id uint) error {
	return db.Delete(&CourseSemester{}, id).Error
}*/

var RemoveCourseSemester = func(id uint) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// Delete related preferences
		if err := tx.Where("course_sem_id = ?", id).Delete(&Preferences{}).Error; err != nil {
			return err
		}
		// Delete the course semester
		if err := tx.Delete(&CourseSemester{}, id).Error; err != nil {
			return err
		}
		return nil
	})
}

var RemoveCourseSemesters = func(courseID uint) error {
	return db.Where("course_id = ?", courseID).Delete(&CourseSemester{}).Error
}

var UpdateCourseSemester = func(id uint, courseID uint, semester string, mandatoryLevel string, section string, timeslot string) error {
	updates := map[string]interface{}{
		"course_id":       courseID,
		"semester":        semester,
		"mandatory_level": mandatoryLevel,
		"section":         section,
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
