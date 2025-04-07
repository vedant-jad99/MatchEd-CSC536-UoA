package models

// Matching model represents the results of a matching iteration.
//id, matching_iteration_id, user_id, course_sem_id, score
type Matching struct {
	ID                  uint `gorm:"primaryKey"`
	MatchingIterationID uint `gorm:"column:matching_iteration_id;not null"`
	UserID              uint `gorm:"column:user_id;not null"`
	CourseID            uint `gorm:"column:course_id;not null"`
	CourseSemID         uint `gorm:"column:course_sem_id;not null"`
	Score               uint `gorm:"not null"`
}

func (Matching) TableName() string {
	return `"match_schema"."matchings"`
}
