package models

// import "gorm.io/gorm"

type User struct {
	// gorm.Model        // Adds ID, CreatedAt, UpdatedAt, DeletedAt
	ID         uint   `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	Email      string `json:"email"`
}

func (User) TableName() string {
	return "match_schema.users";
}