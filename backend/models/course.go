package models

// Course model represents a course entry.
// TODO change time reprentation of days and times the course section is scheduled for

// TODO add the following fields to the schema:
// number, location, semesters
type Course struct {
	//gorm.Model        // Adds ID, CreatedAt, UpdatedAt, DeletedAt
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar;not null"`
	Type string `gorm:"type:varchar;not null"`
}

func (Course) TableName() string {
	return "match_schema.courses"
}
