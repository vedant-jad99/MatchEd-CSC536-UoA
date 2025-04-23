package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) {
	dsn := "host=localhost user=match_user password=swifty dbname=match_db_test port=5432 sslmode=disable search_path=match_schema"
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test DB: %v", err)
	}

	// Point Gorm to use the schema explicitly
	err = dbConn.Exec(`SET search_path TO match_schema`).Error
	if err != nil {
		t.Fatalf("Failed to set schema search path: %v", err)
	}

	// Repoint your global `db`
	db = dbConn
	clearDB(t)
}

func clearDB(t *testing.T) {
	// Clear the database
	err := db.Exec("TRUNCATE TABLE match_schema.course_sem, " +
		"match_schema.users, match_schema.preferences, match_schema.matchings," +
		" match_schema.roles, match_schema.auth," +
		" match_schema.matching_iterations, match_schema.courses CASCADE").Error
	if err != nil {
		t.Fatalf("Failed to truncate tables: %v", err)
	}
}

func TestAddAndFetchCourseSemester(t *testing.T) {
	setupTestDB(t)

	err := AddCourseSemester(1, "Fall")
	assert.Nil(t, err)

	results, err := FetchCourseSemesters()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(results))
	assert.Equal(t, "Fall", results[0].Semester)

	clearDB(t)
}

func TestFetchCourseSemester(t *testing.T) {
	setupTestDB(t)

	cs := CourseSemester{CourseID: 1, Semester: "Fall"}
	txn := db.Create(&cs)
	if txn.Error != nil {
		t.Fatalf("Failed to create course semester: %v", txn.Error)
	}
	fetched, err := FetchCourseSemester(cs.ID)
	assert.Nil(t, err)
	assert.Equal(t, cs.ID, fetched.ID)
}

func TestUpdateCourseSemester(t *testing.T) {
	setupTestDB(t)

	cs := CourseSemester{CourseID: 1, Semester: "Fall", MandatoryLevel: "Low", Timeslot: "TTh"}
	db.Create(&cs)

	err := UpdateCourseSemester(cs.ID, 2, "Spring", "High", "MWF")
	assert.Nil(t, err)

	var updated CourseSemester
	db.First(&updated, cs.ID)
	assert.Equal(t, uint(2), updated.CourseID)
	assert.Equal(t, "Spring", updated.Semester)
	assert.Equal(t, "High", updated.MandatoryLevel)
	assert.Equal(t, "MWF", updated.Timeslot)
}

func TestRemoveCourseSemester(t *testing.T) {
	setupTestDB(t)

	cs := CourseSemester{CourseID: 1, Semester: "Fall"}
	db.Create(&cs)

	err := RemoveCourseSemester(cs.ID)
	assert.Nil(t, err)

	var check CourseSemester
	result := db.First(&check, cs.ID)
	assert.Error(t, result.Error)
}

func TestCreateUserAndFetchAllUsers(t *testing.T) {
	setupTestDB(t)

	UpsertUser(User{Name: "Alice", Email: "alice@test.com", NumReqCourses: 3})
	UpsertUser(User{Name: "Bob", Email: "bob@test.com", NumReqCourses: 2})

	users, err := FetchAllUsers()
	assert.Nil(t, err)
	assert.Len(t, users, 2)
}

func TestRemoveUser(t *testing.T) {
	setupTestDB(t)

	user := User{Name: "Temp", Email: "temp@test.com", NumReqCourses: 1}
	db.Create(&user)

	err := DeleteUser(user.ID)
	assert.Nil(t, err)

	var u User
	result := db.First(&u, user.ID)
	assert.Error(t, result.Error)
}

func TestFetchPreferences(t *testing.T) {
	setupTestDB(t)

	p := Preferences{UserID: 1, Semester: "Fall"}
	db.Create(&p)

	prefs, err := FetchPreferences(1)
	println(len(prefs))
	assert.Nil(t, err)
	assert.Len(t, prefs, 1)
	assert.Equal(t, "Fall", prefs[0].Semester)
}

func TestFetchAllPreferences(t *testing.T) {
	setupTestDB(t)

	db.Create(&Preferences{UserID: 1, Semester: "Spring"})
	db.Create(&Preferences{UserID: 2, Semester: "Spring"})

	prefs, err := FetchAllPreferences()
	assert.Nil(t, err)
	assert.Len(t, prefs, 2)
}

func TestFetchMatchingsByIterationID(t *testing.T) {
	setupTestDB(t)

	m := Matching{MatchingIterationID: 7}
	db.Create(&m)

	result, err := FetchMatchingsByIterationID(7)
	assert.Nil(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, uint(7), result[0].MatchingIterationID)
}

func TestFetchLatestMatchings(t *testing.T) {
	setupTestDB(t)

	iter := MatchingIteration{}
	db.Create(&iter)
	m := Matching{MatchingIterationID: iter.ID}
	db.Create(&m)

	latest, err := FetchLatestMatchings()
	assert.Nil(t, err)
	assert.Len(t, latest, 1)
	assert.Equal(t, iter.ID, latest[0].MatchingIterationID)
}
