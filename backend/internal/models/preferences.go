package models

import (
	"errors"
)

type Preferences struct {
	ID               uint   `json:"id" gorm:"primaryKey"`
	Semester         string `json:"semester" gorm:"varchar;not null"`
	UserID           uint   `json:"user_id" gorm:"column:user_id; not null"`
	CourseSemesterID uint   `json:"course_sem_id" gorm:"column:course_sem_id; not null"`
	PreferenceLevel  string `json:"preference_level" gorm:"column:preference_level;not null"`
	PreferenceWeight int    `json:"preference_weight" gorm:"column:preference_weight;not null"`
}

func (Preferences) TableName() string {
	return "match_schema.preferences"
}

// Edit Preference Color by Course Name and Faculty (User) Name
func EditPreferenceColorByCourseAndUser(courseName, userName, newPreferenceColor string) (*Preferences, error) {
	// Find the course by name
	var course Course
	if err := db.Where("name = ?", courseName).First(&course).Error; err != nil {
		return nil, errors.New("course not found")
	}

	// Find the user (faculty) by name
	var user User
	if err := db.Where("name = ?", userName).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	// todo, more than one CourseSemester per course?
	// Find the CourseSemester by course name
	var courseSemester CourseSemester
	if err := db.Where("course_id IN (SELECT id FROM match_schema.courses WHERE name = ?)", courseName).First(&courseSemester).Error; err != nil {
		return nil, errors.New("course semester not found")
	}

	// Find the preference by course ID and user ID
	var preference Preferences
	if err := db.Where("course_sem_id = ? AND user_id = ?", courseSemester.ID, user.ID).First(&preference).Error; err != nil {
		return nil, errors.New("preference not found")
	}

	// Update the preference color
	preference.PreferenceLevel = newPreferenceColor
	if err := db.Save(&preference).Error; err != nil {
		return nil, err
	}

	return &preference, nil
}

var FetchAllPreferences = func() ([]Preferences, error) {
	var prefs []Preferences
	err := db.Find(&prefs).Error
	return prefs, err
}

var FetchPreferences = func(userID uint) ([]Preferences, error) {
	var prefs []Preferences
	err := db.Where("user_id = ?", userID).Find(&prefs).Error
	return prefs, err
}
var FetchPreferencesBySemester = func(semester string) ([]Preferences, error) {
	var prefs []Preferences
	err := db.Where("semester = ?", semester).Find(&prefs).Error
	return prefs, err
}
