package controllers

import (
	"log"

	"gorm.io/gorm"
)

func InitializeTables(db *gorm.DB)
{
	log.Printf("Initializing database tables from init-data/*.csv")
	LoadCSVtoDatabase(db, "../models/init-data/users.csv", models.User{})
}

func LoadCSVtoDatabase(db *gorm.DB, filePath string, modelType interface{}) {
	jsonData, err := LoadCSVToJSON(filePath)
	if err != nil {
		log.Printf("Error loading CSV:%s", err)
		return
	}

	log.Printf("loaded csv")
	log.Printf("JSON Entry: %s", jsonData) // Debugging output

	err = AddFromJSON(db, jsonData, modelType)
	if err != nil {
		log.Printf("Error adding users:%s", err)
	} else {
		log.Printf("Users successfully added from%s", filePath)
	}
}
