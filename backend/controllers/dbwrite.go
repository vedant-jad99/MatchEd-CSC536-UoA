package controllers

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"

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
