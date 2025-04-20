package models

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"

	"gorm.io/gorm"
)

func InitializeTables(db *gorm.DB) {
	log.Printf("Initializing database tables from init-data/*.csv")
	// ORDER MATTERS HERE, because of database contstraints

	// auth file
	csvPath := getFilePath("../models/init-data/auth.csv")
	LoadCSVtoDatabase(db, csvPath, Auth{})

	// roles file
	csvPath = getFilePath("../models/init-data/roles.csv")
	LoadCSVtoDatabase(db, csvPath, Role{})

	// matching iterations file
	csvPath = getFilePath("../models/init-data/matching_iterations.csv")
	LoadCSVtoDatabase(db, csvPath, MatchingIteration{})

	// courses file
	csvPath = getFilePath("../models/init-data/courses.csv")
	LoadCSVtoDatabase(db, csvPath, Course{})

	// course semesters file
	csvPath = getFilePath("../models/init-data/course_semester.csv")
	LoadCSVtoDatabase(db, csvPath, CourseSemester{})

	// user file
	csvPath = getFilePath("../models/init-data/users.csv")
	LoadCSVtoDatabase(db, csvPath, User{})

	// matchings file
	csvPath = getFilePath("../models/init-data/matchings.csv")
	LoadCSVtoDatabase(db, csvPath, Matching{})

	// preferences file
	csvPath = getFilePath("../models/init-data/preferences.csv")
	LoadCSVtoDatabase(db, csvPath, Preferences{})

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

// Generic function to load CSV data and return JSON
func LoadCSVToJSON(filePath string) ([]map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var data []map[string]string
	headers := records[0]
	for _, record := range records[1:] {
		entry := make(map[string]string)
		for i, value := range record {
			entry[headers[i]] = value
		}
		data = append(data, entry)
	}
	return data, nil
}

// Function to process JSON data and insert into the database
// Assumes that the JSON keys EXACTLY MATCH the modeltype keys
func AddFromJSON(db *gorm.DB, jsonData []map[string]string, modelType interface{}) error {
	modelTypeValue := reflect.TypeOf(modelType)
	if modelTypeValue.Kind() != reflect.Struct {
		return errors.New("modelType must be a struct")
	}

	for _, entry := range jsonData {
		processedEntry := make(map[string]interface{})

		for key, value := range entry {
			field, found := modelTypeValue.FieldByName(key)
			if !found {
				processedEntry[key] = value
				continue
			}

			switch field.Type.Kind() {
			case reflect.Uint, reflect.Uint32, reflect.Uint64:
				num, err := strconv.ParseUint(value, 10, 64)
				if err != nil {
					return errors.New("invalid number format for field: " + key)
				}
				processedEntry[key] = uint(num)
			case reflect.Int, reflect.Int32, reflect.Int64:
				num, err := strconv.Atoi(value)
				if err != nil {
					return errors.New("invalid number format for field: " + key)
				}
				processedEntry[key] = num
			default:
				processedEntry[key] = value
			}
		}

		jsonBytes, err := json.Marshal(processedEntry)
		if err != nil {
			return err
		}

		newModel := reflect.New(modelTypeValue).Interface()
		err = json.Unmarshal(jsonBytes, newModel)
		if err != nil {
			return err
		}

		if err := db.Create(newModel).Error; err != nil {
			return err
		}
	}
	return nil
}
