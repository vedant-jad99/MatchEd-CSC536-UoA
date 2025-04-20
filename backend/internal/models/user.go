package models

// import "gorm.io/gorm"

type User struct {
	Id            uint   `json:"id" gorm:"primaryKey"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	NumReqCourses int    `json:"num_req_courses"`
}

func (User) TableName() string {
	return "match_schema.users"
}

func FetchAllUsers() ([]User, error) {
	var users []User
	txn := db.Find(&users);
	
	return users, txn.Error
}

func FetchUserById(id uint) (User, error) {
	var user User
	txn := db.First(&user, id);
	
	return user, txn.Error
}

func CreateUser(user User) (User, error) {
	txn := db.Create(&user);
	
	return user, txn.Error
}
func UpdateUser(user User) (User, error) {
	txn := db.Save(&user);
	
	return user, txn.Error
}
func DeleteUser(id uint) error {
	var user User
	txn := db.Delete(&user, id);
	
	return txn.Error
}