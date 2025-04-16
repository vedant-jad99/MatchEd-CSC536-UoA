package models

// TODO change time reprentation of days and times the course section is scheduled for
type Course struct {
	Id     uint   `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
	Number string `json:"number"`
	Campus string `json:"campus"`
}

type CourseSem struct {
	Id             uint   `json:"id" gorm:"primaryKey"`
	CourseId       uint   `json:"course_id"`
	Semester       string `json:"semester"`
	MandatoryLevel string `json:"mandatory_level"`
	Timeslot       string `json:"timeslot"`
}

func (Course) TableName() string {
	return "match_schema.courses"
}

func (CourseSem) TableName() string {
	return "match_schema.course_sem"
}