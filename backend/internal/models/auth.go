package models

type Auth struct {
	Id        uint   `json:"id" gorm:"primaryKey"`
	UserId    uint   `json:"user_id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	AuthToken string `json:"auth_token"`
}

func (Auth) TableName() string {
	return "match_schema.auth"
}
