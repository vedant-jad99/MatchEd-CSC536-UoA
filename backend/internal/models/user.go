package models

// import "gorm.io/gorm"

type User struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	Name          string `json:"name" gorm:"type:varchar"`
	Email         string `json:"email" gorm:"type:varchar"`
	NumReqCourses int    `json:"num_req_courses" gorm:"column:num_req_courses;not null"`
}

func (User) TableName() string {
	return "match_schema.users"
}

//cannot assign these to handlers.
// gin works by updating the router context
func FetchAllUsers() ([]User, error) {
	var users []User
	txn := DB.Find(&users)

	return users, txn.Error
}

func FetchUserById(id uint) (User, error) {
	var user User
	txn := DB.First(&user, id)

	return user, txn.Error
}

func CreateUser(user User) (User, error) {
	txn := DB.Create(&user)

	return user, txn.Error
}
func UpdateUser(user User) (User, error) {
	txn := DB.Save(&user)

	return user, txn.Error
}
func DeleteUser(id uint) error {
	var user User
	txn := DB.Delete(&user, id)

	return txn.Error
}
