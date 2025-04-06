package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func main() {
	dsn := "host=localhost user=match_user password=swifty dbname=match_db port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate to appropriate  models
	err = db.AutoMigrate(&User{}, &Preferences{}, &Course{})
	if err != nil {
		panic("failed to migrate models")
	}

	var users []User
	db.Find(&users)
	fmt.Println(users)
}
