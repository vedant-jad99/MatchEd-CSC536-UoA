package models

type Role struct {
	UserId uint   `json:"user_id" gorm:"column:user_id; primaryKey"`
	Role   string `json:"role" gorm:"type:varchar;not null"`
}

func (Role) TableName() string {
	return "match_schema.roles"
}
