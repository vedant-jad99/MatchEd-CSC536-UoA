package models

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