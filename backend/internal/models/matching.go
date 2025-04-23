package models

import (
	"time"
)

type Matching struct {
	ID                  uint    `json:"id" gorm:"primaryKey"`
	MatchingIterationID uint    `json:"matching_iteration_id" gorm:"column:matching_iteration_id;not null"`
	UserID              uint    `json:"user_id" gorm:"column:user_id;not null"`
	CourseSemID         uint    `json:"course_sem_id" gorm:"column:course_sem_id;not null"`
	Score               float64 `json:"score" gorm:"not null"`
}

func (Matching) TableName() string {
	return "match_schema.matchings"
}

type MatchingIteration struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	TriggeredBy uint      `json:"triggered_by" gorm:"column:triggered_by;not null"`
	Status      string    `json:"status" gorm:"column:status;type:varchar;not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at;not null"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at;not null"`
}

func (MatchingIteration) TableName() string {
	return `"match_schema"."matching_iterations"`
}

var CreateMatchingIteration = func(matchingIteration MatchingIteration) (MatchingIteration, error) {
	txn := db.Create(&matchingIteration)

	return matchingIteration, txn.Error
}

var UpdateMatchingIteration = func(matchingIteration MatchingIteration) (MatchingIteration, error) {
	txn := db.Save(&matchingIteration)

	return matchingIteration, txn.Error
}

var FetchMatchingsByIterationID = func(id uint) ([]Matching, error) {
	var matches []Matching
	err := db.Where("matching_iteration_id = ?", id).Find(&matches).Error
	return matches, err
}

var FetchLatestMatchings = func() ([]Matching, error) {
	var latestIteration MatchingIteration
	err := db.Order("updated_at DESC").First(&latestIteration).Error
	if err != nil {
		return nil, err
	}
	return FetchMatchingsByIterationID(latestIteration.ID)
}
