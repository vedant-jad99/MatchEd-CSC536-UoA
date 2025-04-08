package models

// Role model represents user roles.
type Role struct {
	//gorm.Model        // Adds ID, CreatedAt, UpdatedAt, DeletedAt
	UserID uint   `gorm:"column:user_id; primaryKey"`
	Role   string `gorm:"type:varchar;not null"`
}

func (Role) TableName() string {
	return `"match_schema"."role"`
}
