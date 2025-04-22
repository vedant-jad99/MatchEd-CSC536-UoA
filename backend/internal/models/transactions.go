package models

var FetchCourseSemesters = func(semester string) ([]CourseSemester, error) {
	var courseSemesters []CourseSemester
	err := db.Where("semester = ?", semester).Find(&courseSemesters).Error
	return courseSemesters, err
}

var FetchCourseSemester = func(id uint) (CourseSemester, error) {
	var cs CourseSemester
	err := db.First(&cs, id).Error
	return cs, err
}

var AddCourseSemester = func(courseID uint, semester string) error {
	cs := CourseSemester{CourseID: courseID, Semester: semester}
	return db.Create(&cs).Error
}

var RemoveCourseSemester = func(id uint) error {
	return db.Delete(&CourseSemester{}, id).Error
}

var UpdateCourseSemester = func(id uint, courseID uint, semester string, mandatoryLevel string, timeslot string) error {
	updates := map[string]interface{}{
		"course_id":       courseID,
		"semester":        semester,
		"mandatory_level": mandatoryLevel,
		"timeslot":        timeslot,
	}
	return db.Model(&CourseSemester{}).Where("id = ?", id).Updates(updates).Error
}

var CreateUser = func(user User) error {
	return db.Create(&user).Error
}

var FetchAllFaculty = func() ([]User, error) {
	var users []User
	err := db.Find(&users).Error
	return users, err
}

var RemoveFaculty = func(id uint) error {
	return db.Delete(&User{}, id).Error
}

var FetchAllPreferences = func(semester string) ([]Preferences, error) {
	var prefs []Preferences
	err := db.Where("semester = ?", semester).Find(&prefs).Error
	return prefs, err
}

var FetchPreferences = func(userID uint) ([]Preferences, error) {
	var prefs []Preferences
	err := db.Where("user_id = ?", userID).Find(&prefs).Error
	return prefs, err
}

var FetchMatchingsByIterationID = func(id uint) ([]Matching, error) {
	var matches []Matching
	err := db.Where("matching_iteration_id = ?", id).Find(&matches).Error
	return matches, err
}

var FetchLatestMatchings = func() ([]Matching, error) {
	var latestIteration MatchingIteration
	err := db.Order("updated_at DESC").First(&latestIteration).Error
	if err != nil {
		return nil, err
	}
	return FetchMatchingsByIterationID(latestIteration.ID)
}
