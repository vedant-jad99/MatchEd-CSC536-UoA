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

func FetchAllPreferences() ([]Preferences, error) {
	var preferences []Preferences
	txn := DB.Find(&preferences)

	return preferences, txn.Error
}