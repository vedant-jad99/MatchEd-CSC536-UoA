#!/bin/bash
# Backup before testing
./db_testing.sh backup

# Set base URL
BASE_URL=http://localhost:3000/api

# Add timeouts to curl (in seconds)
curl --max-time 10 --connect-timeout 5 http://localhost:3000/api

# Add a CourseSemester
curl -X POST -H "Content-Type: application/json" \
  -d '{"course_id": 1, "semester": "Fall"}' \
  $BASE_URL/course_semester/add
echo -e "\nAdded CourseSemester"

# Fetch CourseSemesters for a semester
curl "$BASE_URL/course_semester/fetch?semester=Fall"
echo -e "\nFetched CourseSemesters"

# Fetch a single CourseSemester by ID
curl "$BASE_URL/course_semester/get?id=1"
echo -e "\nFetched CourseSemester by ID"

# Update a CourseSemester
curl -X PUT -H "Content-Type: application/json" \
  -d '{"id": 1, "course_id": 1, "semester": "Fall", "mandatory_level": "high", "timeslot": "MWF 9AM"}' \
  $BASE_URL/course_semester/update
echo -e "\nUpdated CourseSemester"

# Remove a CourseSemester
# curl -X DELETE "$BASE_URL/course_semester/remove?id=1"
# echo -e "\nDeleted CourseSemester"


# -- FACULTY / USERS TESTS --

# Create a User
curl -X POST -H "Content-Type: application/json" \
  -d '{"name": "John Smith", "email": "jsmith@example.com", "num_req_courses": 2}' \
  $BASE_URL/faculty/create
echo -e "\nCreated User"

# Get all Faculty
curl "$BASE_URL/faculty/fetch"
echo -e "\nFetched Faculty"

# Remove a Faculty (user)
# curl -X DELETE "$BASE_URL/faculty/remove?id=1"
# echo -e "\nDeleted User"


# -- PREFERENCES TESTS --

# Get all preferences for a semester
curl "$BASE_URL/preferences/fetch_all?semester=Fall"
echo -e "\nFetched Preferences by Semester"

# Get preferences for a specific user
curl "$BASE_URL/preferences/fetch_by_user?user_id=1"
echo -e "\nFetched Preferences by User"


# -- MATCHINGS TESTS --

# Get matchings by MatchingIterationID
curl "$BASE_URL/matchings/by_iteration?id=1"
echo -e "\nFetched Matchings by IterationID"

# Get latest matchings
curl "$BASE_URL/matchings/latest"
echo -e "\nFetched Latest Matchings"

# Restore after testing
./db_testing.sh restore
