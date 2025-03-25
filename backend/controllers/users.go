package controllers

import (
	"net/http"

	"backend/models"

	"github.com/gin-gonic/gin"
)

// GetAllUsers handles GET requests to fetch all users
func GetAllUsers(c *gin.Context) {
	var users []models.User

	// Mock fetching users from a database
	// models.DB.Find(&users)

	c.JSON(http.StatusOK, users)
}

// GetUserById handles GET requests to fetch a user by ID
func GetUserById(c *gin.Context) {
	id := c.Param("id")

	// Mock fetching a user by ID
	// var user models.User
	// if err := models.DB.First(&user, id).Error; err != nil {
	//     c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	//     return
	// }

	c.JSON(http.StatusOK, gin.H{"id": id, "message": "User details fetched"})
}

// CreateUser handles POST requests to create a new user
func CreateUser(c *gin.Context) {
	var user models.User

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
	var user models.User

	// Bind JSON input to user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Mock updating user in database
	// models.DB.Model(&user).Where("id = ?", id).Updates(user)

	c.JSON(http.StatusOK, gin.H{"message": "User updated", "id": id, "user": user})
}

// DeleteUser handles DELETE requests to remove a user by ID
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	// Mock deleting user from database
	// models.DB.Delete(&models.User{}, id)

	c.JSON(http.StatusOK, gin.H{"message": "User deleted", "id": id})
}
