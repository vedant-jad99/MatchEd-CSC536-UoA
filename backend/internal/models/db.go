package models

import (
	"fmt"
	// "log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// the name of the database.
// call it from other packages with models.DB
var db *gorm.DB

// TODO initialize gorm with postgres using real dbname
func InitDB() error {
	dsn := getDSN()

	var err error
	// try to connect to database with gorm
	db, err = gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{})
	return err;
}

func getDSN() string {
	// Get database connection details from environment variables
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))
	return dsn
}
