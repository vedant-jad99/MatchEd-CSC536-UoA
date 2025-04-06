package models

import (
	"fmt"
	"log"
	// "os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// the name of the database.
// call it from other packages with models.DB
var DB *gorm.DB

// TODO initialize gorm with postgres using real dbname
func InitDB() error {

	// get dsn
	dsn := getDSN()
	// dsn := "host=localhost user=match_user password=swifty dbname=match_db port=5432"

	//TODO remove this print statement
	//fmt.Println(dsn)

	var err error
	// try to connect to database with gorm
	DB, err = gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{})
	return err;
}

// TODO: there is some bug here with a nil pointer
// isDatabaseAvailable checks if the database is available and returns true if connected
// takes the dsn string
func isDatabaseAvailable() bool {
	if DB == nil {
		log.Println("DB connection is not initialized")
		return false
	}

	sqlDB, err := DB.DB() // Get the underlying *sql.DB object
	if err != nil {
		log.Println("Error getting underlying sql.DB object:", err)
		return false
	}

	if err := sqlDB.Ping(); err != nil {
		log.Println("Error pinging the database:", err)
		return false
	}

	fmt.Println("Successfully connected to the database!")
	return true
}

func getDSN() string {
	// Get database connection details from environment variables
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"localhost",
		"postgres",
		"postgres",
		"postgres", 
		"5432")
	return dsn
}