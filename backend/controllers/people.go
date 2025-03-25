// controllers/people.go
package controllers

import (
	"net/http"

	"backend/models"

	"github.com/gin-gonic/gin"
)

// GetAllPeople handles GET requests to fetch all people
func GetAllPeople(c *gin.Context) {
	var people []models.Person

	// Fetch people from database
	// models.DB.Find(&people)

	c.JSON(http.StatusOK, people)
}

// GetPerson handles GET requests to fetch a single person
func GetPersonById(c *gin.Context) {
	// id := c.Param("id")
	var person models.Person
	// Fetch person from database
	// result := models.DB.First(&person, id)
	// if result.Error != nil {
	//    c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
	//    return
	// }

	c.JSON(http.StatusOK, person)
}

// CreatePerson handles POST requests to create a new person
func CreatePerson(c *gin.Context) {
	var newPerson models.Person

	if err := c.ShouldBindJSON(&newPerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save to database
	// models.DB.Create(&newPerson)

	c.JSON(http.StatusCreated, newPerson)
}

// UpdatePerson handles PUT requests to update an existing person
func UpdatePerson(c *gin.Context) {
	var person models.Person
	//id := c.Param("id") // Get the person ID from the request URL

	// Fetch the existing person from the database (mock example)
	// if err := models.DB.First(&person, id).Error; err != nil {
	//     c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
	//     return
	// }

	// Bind JSON data to the person object
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Save the updated person in the database
	// models.DB.Save(&person)

	c.JSON(http.StatusOK, gin.H{"message": "Person updated successfully", "person": person})
}

// DeletePerson handles DELETE requests to remove a person by ID
func DeletePerson(c *gin.Context) {
	//id := c.Param("id") // Get the person ID from the request URL

	// Fetch the existing person from the database (mock example)
	// var person models.Person
	// if err := models.DB.First(&person, id).Error; err != nil {
	//     c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
	//     return
	// }

	// Delete the person from the database
	// models.DB.Delete(&person)

	c.JSON(http.StatusOK, gin.H{"message": "Person deleted successfully"})
}
