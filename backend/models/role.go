package models

// Role model represents user roles.
type Role struct {
	//gorm.Model        // Adds ID, CreatedAt, UpdatedAt, DeletedAt
	ID     uint   `gorm:"primaryKey"`
	UserID uint   `gorm:"column:user_id;"`
	Role   string `gorm:"type:varchar;not null"`
}
