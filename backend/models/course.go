package models

// Course model represents a course entry.
// TODO change time reprentation of days and times the course section is scheduled for

// TODO add the following fields to the schema:
// number, location, semesters
type Course struct {
	//gorm.Model        // Adds ID, CreatedAt, UpdatedAt, DeletedAt
	ID        uint   `gorm:"primaryKey"`
	Number    string `gorm:"type:varchar;not null"` // occasionally contains chars (199H)
	Name      string `gorm:"type:varchar;not null"`
	Campus    string `gorm:"type:varchar;not null"`
	Semesters string `gorm:"type:varchar;not null"`
}

func (Course) TableName() string {
	return `"match_schema"."courses"`
}
