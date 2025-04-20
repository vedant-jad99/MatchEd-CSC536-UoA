package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// import "gorm.io/gorm"

type User struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	Name          string `json:"name" gorm:"type:varchar"`
	Email         string `json:"email" gorm:"type:varchar"`
	NumReqCourses int    `json:"num_req_courses" gorm:"column:num_req_courses;not null"`
}

func (User) TableName() string {
	return "match_schema.users"
}

// Functions to add and save models to the database
func AddUser(db *gorm.DB, name string, email string) (User, error) {
	user := User{Name: name, Email: email}
	err := db.Create(&user).Error
	return user, err
}

// GetAllUsers handles GET requests to fetch all users
func GetAllUsers(c *gin.Context) {
	var users []User

	// Mock fetching users from a database
	// DB.Find(&users)

	c.JSON(http.StatusOK, users)
}

// GetUserById handles GET requests to fetch a user by ID
func GetUserById(c *gin.Context) {
	id := c.Param("id")

	// Mock fetching a user by ID
	// var user User
	// if err := DB.First(&user, id).Error; err != nil {
	//     c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	//     return
	// }

	c.JSON(http.StatusOK, gin.H{"id": id, "message": "User details fetched"})
}

// CreateUser handles POST requests to create a new user
func CreateUser(c *gin.Context) {
	var user User

	// Bind JSON input to user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Mock saving user to database
	// db.Create(&user)

	c.JSON(http.StatusCreated, gin.H{"message": "User created", "user": user})
}

// UpdateUser handles PUT requests to update an existing user
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user User

	// Bind JSON input to user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Mock updating user in database
	// DB.Model(&user).Where("id = ?", id).Updates(user)

	c.JSON(http.StatusOK, gin.H{"message": "User updated", "id": id, "user": user})
}

// DeleteUser handles DELETE requests to remove a user by ID
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	// Mock deleting user from database
	// DB.Delete(&User{}, id)

	c.JSON(http.StatusOK, gin.H{"message": "User deleted", "id": id})
}

/*
//cannot assign these to handlers.
// gin works by updating the router context
func FetchAllUsers() ([]User, error) {
	var users []User
	txn := DB.Find(&users)

	return users, txn.Error
}

func FetchUserById(id uint) (User, error) {
	var user User
	txn := DB.First(&user, id)

	return user, txn.Error
}

func CreateUser(user User) (User, error) {
	txn := DB.Create(&user)

	return user, txn.Error
}
func UpdateUser(user User) (User, error) {
	txn := DB.Save(&user)

	return user, txn.Error
}
func DeleteUser(id uint) error {
	var user User
	txn := DB.Delete(&user, id)

	return txn.Error
}
*/
