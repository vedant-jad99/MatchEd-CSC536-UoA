package models

import "time"

type Auth struct {
	Id        uint       `json:"id" gorm:"primaryKey"`
	UserID    uint       `json:"user_id" gorm:"column:user_id;not null"`
	StartTime time.Time  `json:"start_time" gorm:"column:start_time;not null"`
	EndTime   *time.Time `json:"end_time" gorm:"column:end_time"`
	AuthToken string     `json:"auth_token" gorm:"column:auth_token;type:varchar;not null;unique"`
}

func (Auth) TableName() string {
	return "match_schema.auth"
}
