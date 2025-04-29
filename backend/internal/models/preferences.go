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

	// Fetch all preferences for the user
	var preferences []Preferences
	if err := db.Where("user_id = ?", user.ID).Find(&preferences).Error; err != nil {
		return nil, errors.New("preferences not found")
	}

	// Iterate over all preferences to find the one matching the course
	var matchingPreference *Preferences
	for _, pref := range preferences {
		// Get the CourseSemester for this preference
		var courseSemester CourseSemester
		if err := db.Where("id = ?", pref.CourseSemesterID).First(&courseSemester).Error; err != nil {
			continue
		}

		// If the course matches, update the preference
		if courseSemester.CourseID == course.ID {
			matchingPreference = &pref
			break
		}
	}

	if matchingPreference == nil {
		return nil, errors.New("preference for the course not found")
	}

	// Update the preference color
	matchingPreference.PreferenceLevel = newPreferenceColor
	if err := db.Save(matchingPreference).Error; err != nil {
		return nil, err
	}

	return matchingPreference, nil
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

var BulkUpsertPreferences = func(prefs []Preferences) ([]Preferences, error) {
	err := db.Save(&prefs).Error
	return prefs, err
}

var BulkDeletePreferences = func(ids []uint) error {
	var prefs []Preferences
	err := db.Where("id IN ?", ids).Delete(&prefs).Error
	return err
}
var DeletePreferences = func(id uint) error {
	var prefs Preferences
	err := db.Delete(&prefs, id).Error
	return err
}
var DeletePreferencesByUserID = func(userID uint) error {
	var prefs Preferences
	err := db.Where("user_id = ?", userID).Delete(&prefs).Error
	return err
}

