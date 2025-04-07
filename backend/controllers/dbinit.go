package controllers

import (
	"backend/models"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"gorm.io/gorm"
)

func InitializeTables(db *gorm.DB) {
	log.Printf("Initializing database tables from init-data/*.csv")
	// order matters here, because of database contstraints

	// user file
	csvPath := getFilePath("../models/init-data/users.csv")
	LoadCSVtoDatabase(db, csvPath, models.User{})

	// courses file
	csvPath = getFilePath("../models/init-data/courses.csv")
	LoadCSVtoDatabase(db, csvPath, models.Course{})

	// matchings file
	csvPath = getFilePath("../models/init-data/matching_iterations.csv")
	LoadCSVtoDatabase(db, csvPath, models.MatchingIteration{})

	// matchings file
	csvPath = getFilePath("../models/init-data/roles.csv")
	LoadCSVtoDatabase(db, csvPath, models.Matching{})

	// matchings file
	csvPath = getFilePath("../models/init-data/matchings.csv")
	LoadCSVtoDatabase(db, csvPath, models.Matching{})

}

func getFilePath(partialPath string) string {
	// Get the directory of the current script
	_, currentFile, _, _ := runtime.Caller(0)
	dir := filepath.Dir(currentFile)

	// Construct the path to the CSV file relative to the script
	csvPath := filepath.Join(dir, partialPath)

	// Check if the file exists
	if _, err := os.Stat(csvPath); err != nil {
		fmt.Println("Error:", err)
	}
	return csvPath
}

func LoadCSVtoDatabase(db *gorm.DB, filePath string, modelType interface{}) {
	jsonData, err := LoadCSVToJSON(filePath)
	if err != nil {
		log.Printf("Error loading CSV:%s", err)
		return
	}

	//log.Printf("loaded csv")
	//log.Printf("JSON Entry: %s", jsonData) // Debugging output

	err = AddFromJSON(db, jsonData, modelType)
	if err != nil {
		log.Printf("Error adding to relation:%s", err)
	} else {
		log.Printf("Users successfully added from%s", filePath)
	}
}
