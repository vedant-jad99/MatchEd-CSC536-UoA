package models

type Role struct {
	UserId uint   `json:"user_id"`
	Role   string `json:"role"`
}

func (Role) TableName() string {
	return "match_schema.roles"
}
