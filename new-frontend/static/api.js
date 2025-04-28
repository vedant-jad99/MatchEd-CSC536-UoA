//API requests 

// ******************************************************
// Functions to handle API requests for courses
const fetchAllCourses = async () => {
    const response = await fetch('/api/course', { method: 'GET' });
    return await response.json();
  };

const fetchCourseById = async (id) => {
  const response = await fetch(`/api/course/${id}`, { method: 'GET' });
  return await response.json();
};

const upsertCourse = async (courseData) => {
  const response = await fetch('/api/course', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(courseData)
  });
  return await response.json();
};

const deleteCourse = async (id) => {
  const response = await fetch(`/api/course/${id}`, { method: 'DELETE' });
  return await response.json();
};

// Functions to handle API requests for users
const fetchAllUsers = async () => {
  const response = await fetch('/api/users', { method: 'GET' });
  return await response.json();
};

const fetchUserById = async (id) => {
  const response = await fetch(`/api/users/${id}`, { method: 'GET' });
  return await response.json();
};

const upsertUser = async (userData) => {
  const response = await fetch('/api/users', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(userData)
  });
  return await response.json();
};

const deleteUser = async (id) => {
  const response = await fetch(`/api/users/${id}`, { method: 'DELETE' });
  return await response.json();
};

// ******************************************************
// Functions to handle API requests for course semesters
const fetchAllCourseSemesters = async () => {
  const response = await fetch('/api/course_semester', { method: 'GET' });
  return await response.json();
};

const fetchCourseSemester = async (id) => {
  const response = await fetch('/api/course_semester/get', {
    method: 'GET',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ id })
  });
  return await response.json();
};

const addCourseSemester = async (courseSemesterData) => {
  const response = await fetch('/api/course_semester/add', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(courseSemesterData)
  });
  return await response.json();
};

const removeCourseSemester = async (id) => {
  const response = await fetch('/api/course_semester/remove', {
    method: 'DELETE',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ id })
  });
  return await response.json();
};

const updateCourseSemester = async (courseSemesterData) => {
  const response = await fetch('/api/course_semester/update', {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(courseSemesterData)
  });
  return await response.json();
};

// ******************************************************
// Functions to handle API requests for matchings
const fetchLatestMatchings = async () => {
  const response = await fetch('/api/matchings/latest', { method: 'GET' });
  return await response.json();
};

const fetchMatchingsByIterationId = async (id) => {
  const response = await fetch('/api/matchings/by_iteration', {
    method: 'GET',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ id })
  });
  return await response.json();
};

// ******************************************************
// Functions to handle API requests for preferences
const fetchAllPreferences = async () => {
  const response = await fetch('/api/preferences/fetch_all', { method: 'GET' });
  return await response.json();
};

const fetchPreferencesByUserId = async (userId) => {
  const response = await fetch('/api/preferences/fetch_by_user', {
    method: 'GET',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ user_id: userId })
  });
  return await response.json();
};

// Ping function to test the server
const pingServer = async () => {
  const response = await fetch('/api/ping', { method: 'GET' });
  return await response.json();
};
