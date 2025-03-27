package controllers

import (
	"backend/models"
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"time"

	"gorm.io/gorm"
)

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

// Functions to add and save models to the database
func AddUser(db *gorm.DB, name string, email string) (models.User, error) {
	user := models.User{Name: name, Email: email}
	err := db.Create(&user).Error
	return user, err
}

func AddCourse(db *gorm.DB, number string, name string, campus string, semesters string) (models.Course, error) {
	course := models.Course{Number: number, Name: name, Campus: campus, Semesters: semesters}
	err := db.Create(&course).Error
	return course, err
}

func AddCourseSemester(db *gorm.DB, courseID uint, semester string) (models.CourseSemester, error) {
	cs := models.CourseSemester{CourseID: courseID, Semester: semester}
	err := db.Create(&cs).Error
	return cs, err
}

func AddPreferences(db *gorm.DB, userID, courseID uint, priority int) (models.Preferences, error) {
	pref := models.Preferences{UserID: userID, CourseID: courseID, Priority: priority}
	err := db.Create(&pref).Error
	return pref, err
}

func AddMatchingIteration(db *gorm.DB) (models.MatchingIteration, error) {
	mi := models.MatchingIteration{CreatedAt: time.Now()}
	err := db.Create(&mi).Error
	return mi, err
}

func AddMatching(db *gorm.DB, matchingIterationID, userID, courseID uint) (models.Matching, error) {
	match := models.Matching{MatchingIterationID: matchingIterationID, UserID: userID, CourseID: courseID}
	err := db.Create(&match).Error
	return match, err
}
