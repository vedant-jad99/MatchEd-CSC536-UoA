package models

// import "gorm.io/gorm"

type User struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	Name          string `json:"name" gorm:"type:varchar"`
	Email         string `json:"email" gorm:"type:varchar"`
	NumReqCourses int    `json:"num_req_courses" gorm:"column:num_req_courses"`
}

func (User) TableName() string {
	return "match_schema.users"
}

//cannot assign these to handlers.
// gin works by updating the router context
var FetchAllUsers = func () ([]User, error) {
	var users []User
	txn := db.Find(&users)

	return users, txn.Error
}

var FetchUserById = func (id uint) (User, error) {
	var user User
	txn := db.First(&user, id)

	return user, txn.Error
}

var UpsertUser = func (user User) (User, error) {
	txn := db.Save(&user)

	return user, txn.Error
}
var DeleteUser = func (id uint) error {
	var user User
	txn := db.Delete(&user, id)

	return txn.Error
}
