package models

import (
	// "gorm.io/gorm"
	// "net/http"
	// "github.com/gin-gonic/gin"
)

type Preferences struct {
	// gorm.Model            
	Semester      string  `json:"semester"`       
	UserID        uint    `json:"user_id"`        
	CourseSemID   uint    `json:"course_sem_id"`  
	PreferenceLevel int    `json:"preference_level"` 
}

func (Preferences) TableName() string {
	return `"match_schema"."preferences"`
}
