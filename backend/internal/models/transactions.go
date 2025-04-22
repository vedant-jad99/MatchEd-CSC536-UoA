package models


var FetchAllFaculty = func() ([]User, error) {
	var users []User
	err := db.Find(&users).Error
	return users, err
}

var RemoveFaculty = func(id uint) error {
	return db.Delete(&User{}, id).Error
}

