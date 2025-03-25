package models

import (
	"gorm.io/gorm"
)

// TODO change time reprentation of days and times the course section is scheduled for
type Course struct {
	gorm.Model             // Adds ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string      `json:"name"`
	People     []*Person   `json:"people"`
	Schedule   map[int]int `json:"schedule"` //day -> minute? ex monday@9 = 1 -> 9*60
}
