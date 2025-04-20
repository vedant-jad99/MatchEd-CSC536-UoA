package models

type Preferences struct {
	Id               uint   `json:"id" gorm:"primaryKey"`
	Semester         string `json:"semester"`
	UserId           uint   `json:"user_id"`
	CourseSemId      uint   `json:"course_sem_id"`
	PreferenceLevel  string `json:"preference_level"`
	PreferenceWeight int    `json:"preference_weight"`
}

func (Preferences) TableName() string {
	return "match_schema.preferences"
}