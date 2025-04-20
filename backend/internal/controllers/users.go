package controllers

import (
	"net/http"
	"strconv"

	"backend/internal/models"

	"github.com/gin-gonic/gin"
)

// GetAllUsers handles GET requests to fetch all users
func GetAllUsers(c *gin.Context) {
	users, err := models.FetchAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUserById handles GET requests to fetch a user by ID
func GetUserById(c *gin.Context) {
	id := c.Param("id")

	// Convert id to uint
	parsedID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := models.FetchUserById(uint(parsedID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id, "user": user})
}

// CreateUser handles POST requests to create a new user
func CreateUser(c *gin.Context) {
	var user models.User
	// Bind JSON input to user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	user, err := models.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created", "user": user})
}

// UpdateUser handles PUT requests to update an existing user
func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	// Convert id to uint
	parsedID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	var user models.User
	// Bind JSON input to user struct
	err = c.ShouldBindJSON(&user); 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	user.Id = uint(parsedID)
	user, err = models.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated", "user": user})
}

// DeleteUser handles DELETE requests to remove a user by ID
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	// Convert id to uint
	parsedID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	err = models.DeleteUser(uint(parsedID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}