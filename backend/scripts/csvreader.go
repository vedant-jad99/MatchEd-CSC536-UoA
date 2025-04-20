// controllers/people.go
package controllers

import (
	"encoding/csv"
	"os"
)

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
