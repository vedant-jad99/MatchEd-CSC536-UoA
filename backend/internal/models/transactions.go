package models

func FetchCourseSemesters(semester string) ([]CourseSemester, error) {
	var courseSemesters []CourseSemester
	err := db.Where("semester = ?", semester).Find(&courseSemesters).Error
	return courseSemesters, err
}

func FetchCourseSemester(id uint) (CourseSemester, error) {
	var cs CourseSemester
	err := db.First(&cs, id).Error
	return cs, err
}

func AddCourseSemester(courseID uint, semester string) error {
	cs := CourseSemester{CourseID: courseID, Semester: semester}
	return db.Create(&cs).Error
}

func RemoveCourseSemester(id uint) error {
	return db.Delete(&CourseSemester{}, id).Error
}

func UpdateCourseSemester(id uint, courseID uint, semester string, mandatoryLevel string, timeslot string) error {
	updates := map[string]interface{}{
		"course_id":       courseID,
		"semester":        semester,
		"mandatory_level": mandatoryLevel,
		"timeslot":        timeslot,
	}
	return db.Model(&CourseSemester{}).Where("id = ?", id).Updates(updates).Error
}

func CreateUser(user User) error {
	return db.Create(&user).Error
}

func FetchAllFaculty() ([]User, error) {
	var users []User
	err := db.Find(&users).Error
	return users, err
}

func RemoveFaculty(id uint) error {
	return db.Delete(&User{}, id).Error
}

func FetchAllPreferences(semester string) ([]Preferences, error) {
	var prefs []Preferences
	err := db.Where("semester = ?", semester).Find(&prefs).Error
	return prefs, err
}

func FetchPreferences(userID uint) ([]Preferences, error) {
	var prefs []Preferences
	err := db.Where("user_id = ?", userID).Find(&prefs).Error
	return prefs, err
}

func FetchMatchingsByIterationID(id uint) ([]Matching, error) {
	var matches []Matching
	err := db.Where("matching_iteration_id = ?", id).Find(&matches).Error
	return matches, err
}

func FetchLatestMatchings() ([]Matching, error) {
	var latestIteration MatchingIteration
	err := db.Order("updated_at DESC").First(&latestIteration).Error
	if err != nil {
		return nil, err
	}
	return FetchMatchingsByIterationID(latestIteration.ID)
}
