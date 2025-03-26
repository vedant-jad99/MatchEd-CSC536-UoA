package controllers

import (
	"backend/models"
	"log"

	"gorm.io/gorm"
)

func LoadUsers(db *gorm.DB, filePath string) {
	jsonData, err := LoadCSVToJSON(filePath)
	if err != nil {
		log.Printf("Error loading CSV:%s", err)
		return
	}

	log.Printf("loaded csv")
	log.Printf("JSON Entry: %s", jsonData) // Debugging output

	err = AddFromJSON(db, jsonData, models.User{})
	if err != nil {
		log.Printf("Error adding users:%s", err)
	} else {
		log.Printf("Users successfully added from%s", filePath)
	}
}
