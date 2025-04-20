package models

type Role struct {
	UserID uint   `json:"UserId" gorm:"column:user_id; primaryKey"`
	Role   string `json:"Role" gorm:"type:varchar;not null"`
}

func (Role) TableName() string {
	return "match_schema.roles"
}
