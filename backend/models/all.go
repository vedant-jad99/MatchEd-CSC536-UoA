package models

import "gorm.io/gorm"

type Preferences struct {
	gorm.Model            
	Semester      string  `json:"semester"`       
	UserID        uint    `json:"user_id"`        
	CourseSemID   uint    `json:"course_sem_id"`  
	PreferenceLevel int    `json:"preference_level"` 
}