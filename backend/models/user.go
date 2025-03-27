package models

// User model represents application users.
type User struct {
	//gorm.Model        // Adds ID, CreatedAt, UpdatedAt, DeletedAt
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"type:varchar"`
	Email string `gorm:"type:varchar"`
}

// user is a restricted name in postgres
// grom converts tablenames to lowercase and pluralizes them
// here is the workaround
func (User) TableName() string {
	return `"match_schema"."user"`
}
