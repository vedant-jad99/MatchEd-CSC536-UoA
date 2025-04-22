package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	dsn := "host=localhost user=match_user password=swifty dbname=match_db port=5432 sslmode=disable search_path=match_schema"
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test DB: %v", err)
	}

	// Optional safety check
	err = dbConn.Exec("CREATE SCHEMA IF NOT EXISTS match_schema").Error
	if err != nil {
		t.Fatalf("Failed to ensure schema: %v", err)
	}

	// Point Gorm to use the schema explicitly
	err = dbConn.Exec(`SET search_path TO match_schema`).Error
	if err != nil {
		t.Fatalf("Failed to set schema search path: %v", err)
	}

	// Repoint your global `db`
	db = dbConn
	return db
}

func TestAddAndFetchCourseSemester(t *testing.T) {
	db := setupTestDB(t)

	tx := db.Begin()
	defer tx.Rollback()

	// temporarily redirect global `db` to transaction
	db = tx

	err := AddCourseSemester(1, "Fall")
	assert.Nil(t, err)

	results, err := FetchCourseSemesters("Fall")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(results))
	assert.Equal(t, "Fall", results[0].Semester)
}

func TestFetchCourseSemester(t *testing.T) {
	db := setupTestDB(t)

	tx := db.Begin()
	defer tx.Rollback()

	// temporarily redirect global `db` to transaction
	db = tx

	cs := CourseSemester{CourseID: 1, Semester: "Fall"}
	db.Create(&cs)

	fetched, err := FetchCourseSemester(cs.ID)
	assert.Nil(t, err)
	assert.Equal(t, cs.ID, fetched.ID)
}

func TestUpdateCourseSemester(t *testing.T) {
	db := setupTestDB(t)

	tx := db.Begin()
	defer tx.Rollback()

	// temporarily redirect global `db` to transaction
	db = tx

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
	db := setupTestDB(t)

	tx := db.Begin()
	defer tx.Rollback()

	// temporarily redirect global `db` to transaction
	db = tx

	cs := CourseSemester{CourseID: 1, Semester: "Fall"}
	db.Create(&cs)

	err := RemoveCourseSemester(cs.ID)
	assert.Nil(t, err)

	var check CourseSemester
	result := db.First(&check, cs.ID)
	assert.Error(t, result.Error)
}

func TestCreateUserAndFetchAllFaculty(t *testing.T) {
	db := setupTestDB(t)

	tx := db.Begin()
	defer tx.Rollback()

	// temporarily redirect global `db` to transaction
	db = tx

	UpsertUser(User{Name: "Alice", Email: "alice@test.com", NumReqCourses: 3})
	UpsertUser(User{Name: "Bob", Email: "bob@test.com", NumReqCourses: 2})

	users, err := FetchAllFaculty()
	assert.Nil(t, err)
	assert.Len(t, users, 2)
}

func TestRemoveFaculty(t *testing.T) {
	db := setupTestDB(t)

	tx := db.Begin()
	defer tx.Rollback()

	// temporarily redirect global `db` to transaction
	db = tx

	user := User{Name: "Temp", Email: "temp@test.com", NumReqCourses: 1}
	db.Create(&user)

	err := RemoveFaculty(user.ID)
	assert.Nil(t, err)

	var u User
	result := db.First(&u, user.ID)
	assert.Error(t, result.Error)
}

func TestFetchPreferences(t *testing.T) {
	db := setupTestDB(t)

	tx := db.Begin()
	defer tx.Rollback()

	// temporarily redirect global `db` to transaction
	db = tx

	p := Preferences{UserID: 1, Semester: "Fall"}
	db.Create(&p)

	prefs, err := FetchPreferences(1)
	assert.Nil(t, err)
	assert.Len(t, prefs, 1)
	assert.Equal(t, "Fall", prefs[0].Semester)
}

func TestFetchAllPreferences(t *testing.T) {
	db := setupTestDB(t)

	tx := db.Begin()
	defer tx.Rollback()

	// temporarily redirect global `db` to transaction
	db = tx

	db.Create(&Preferences{UserID: 1, Semester: "Spring"})
	db.Create(&Preferences{UserID: 2, Semester: "Spring"})

	prefs, err := FetchAllPreferences()
	assert.Nil(t, err)
	assert.Len(t, prefs, 2)
}

func TestFetchMatchingsByIterationID(t *testing.T) {
	db := setupTestDB(t)

	tx := db.Begin()
	defer tx.Rollback()

	// temporarily redirect global `db` to transaction
	db = tx

	m := Matching{MatchingIterationID: 7}
	db.Create(&m)

	result, err := FetchMatchingsByIterationID(7)
	assert.Nil(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, uint(7), result[0].MatchingIterationID)
}

func TestFetchLatestMatchings(t *testing.T) {
	db := setupTestDB(t)

	tx := db.Begin()
	defer tx.Rollback()

	// temporarily redirect global `db` to transaction
	db = tx

	iter := MatchingIteration{}
	db.Create(&iter)
	m := Matching{MatchingIterationID: iter.ID}
	db.Create(&m)

	latest, err := FetchLatestMatchings()
	assert.Nil(t, err)
	assert.Len(t, latest, 1)
	assert.Equal(t, iter.ID, latest[0].MatchingIterationID)
}
