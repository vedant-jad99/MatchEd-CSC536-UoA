package models

import (
	"errors"
)

// import "gorm.io/gorm"

type User struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	Name          string `json:"name" gorm:"type:varchar"`
	Email         string `json:"email" gorm:"type:varchar"`
	NumReqCourses int    `json:"num_req_courses" gorm:"column:num_req_courses"`
}

func (User) TableName() string {
	return "match_schema.users"
}

// update by name
// Add User (Faculty) by name
func AddUserByName(name string) (*User, error) {
	// Check if the user (faculty) already exists by name
	var existingUser User
	if err := db.Where("name = ?", name).First(&existingUser).Error; err == nil {
		return nil, errors.New("faculty already exists")
	}

	// Create new user (faculty) record
	user := User{Name: name}
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Remove User (Faculty) by name
func RemoveUserByName(name string) error {
	// Find the user by name
	var user User
	if err := db.Where("name = ?", name).First(&user).Error; err != nil {
		return err // User not found
	}

	// Delete the user (faculty)
	if err := db.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

// Edit User (Faculty) by name
func EditUserByName(oldName, newName string) (*User, error) {
	// Find the user by old name
	var user User
	if err := db.Where("name = ?", oldName).First(&user).Error; err != nil {
		return nil, err // User not found
	}

	// Update the user's name
	user.Name = newName
	if err := db.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// cannot assign these to handlers.
// gin works by updating the router context
var FetchAllUsers = func() ([]User, error) {
	var users []User
	txn := db.Find(&users)

	return users, txn.Error
}

var FetchUserById = func(id uint) (User, error) {
	var user User
	txn := db.First(&user, id)

	return user, txn.Error
}

var UpsertUser = func(user User) (User, error) {
	txn := db.Save(&user)

	return user, txn.Error
}
var DeleteUser = func(id uint) error {
	var user User
	txn := db.Delete(&user, id)

	return txn.Error
}
