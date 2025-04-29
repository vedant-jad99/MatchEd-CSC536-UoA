//API requests 

// ******************************************************
// Functions to handle API requests for courses
export const fetchAllCourses = async () => {
    const response = await fetch('/api/course', { method: 'GET' });
    return await response.json();
  };

// fetch course_sem_id and course name
export const fetchAllCoursesFormatted = async () => {
  const response = await fetch('/api/course/formatted', { method: 'GET' });
  return await response.json();
};

export const fetchCourseById = async (id) => {
  const response = await fetch(`/api/course/${id}`, { method: 'GET' });
  return await response.json();
};

export const upsertCourse = async (courseData) => {
  const response = await fetch('/api/course', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(courseData)
  });
  return await response.json();
};

export const deleteCourse = async (id) => {
  const response = await fetch(`/api/course/${id}`, { method: 'DELETE' });
  return await response.json();
};

// Functions to handle API requests for users
export const fetchAllUsers = async () => {
  const response = await fetch('/api/users', { method: 'GET' });
  return await response.json();
};

export const fetchAllUsersFormatted = async () => {
  const response = await fetch('/api/users/fetch_formatted', { method: 'GET' });
  return await response.json();
};

export const fetchUserById = async (id) => {
  const response = await fetch(`/api/users/${id}`, { method: 'GET' });
  return await response.json();
};

export const upsertUser = async (userData) => {
  const response = await fetch('/api/users', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(userData)
  });
  return await response.json();
};

export const deleteUser = async (id) => {
  const response = await fetch(`/api/users/${id}`, { method: 'DELETE' });
  return await response.json();
};

// ******************************************************
// Functions to handle API requests for course semesters

export const upsertCourseAndSemester = async (courseData) => {
  const response = await fetch('/api/course_semester/upsert', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(courseData)
  });
  return await response.json();
};

export const fetchAllCourseSemesters = async () => {
  const response = await fetch('/api/course_semester', { method: 'GET' });
  return await response.json();
};

export const fetchCourseSemester = async (id) => {
  const response = await fetch(`/api/course_semester/get/${id}`, {
    method: 'GET',
    headers: { 'Content-Type': 'application/json' },
  });
  return await response.json();
};

export const addCourseSemester = async (courseSemesterData) => {
  const response = await fetch('/api/course_semester/add', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(courseSemesterData)
  });
  return await response.json();
};

export const removeCourseSemester = async (id) => {
  const response = await fetch(`/api/course_semester/remove/${id}`, {
    method: 'DELETE',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ id })
  });
  return await response.json();
};

export const updateCourseSemester = async (courseSemesterData) => {
  const response = await fetch('/api/course_semester/update', {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(courseSemesterData)
  });
  return await response.json();
};

// ******************************************************
// Functions to handle API requests for matchings
export const fetchLatestMatchings = async () => {
  const response = await fetch('/api/matchings/latest', { method: 'GET' });
  return await response.json();
};

export const triggerMatchingEngine = async () => {
  const response = await fetch('/api/matchings/trigger', { method: 'POST' });
  return await response.json();
};

export const fetchLatestMatchingsFormatted = async () => {
  const response = await fetch('/api/matchings/latest_formatted', { method: 'GET' });
  return await response.json();
};

export const fetchMatchingsByIterationId = async (id) => {
  const response = await fetch('/api/matchings/by_iteration', {
    method: 'GET',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ id })
  });
  return await response.json();
};

// ******************************************************
// Functions to handle API requests for preferences

export const upsertPreference = async (preferenceData) => {
  const response = await fetch('/api/preferences/upsert', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(preferenceData)
  });
  return await response.json();
};

export const fetchAllPreferences = async () => {
  const response = await fetch('/api/preferences/fetch_all', { method: 'GET' });
  return await response.json();
};

export const fetchAllPreferencesFormatted = async () => {
  const response = await fetch('/api/preferences/fetch_formatted', { method: 'GET' });
  return await response.json();
};


export const fetchPreferencesByUserId = async (userId) => {
  const response = await fetch('/api/preferences/fetch_by_user', {
    method: 'GET',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ user_id: userId })
  });
  return await response.json();
};

export const fetchPreferenceById = async (id) => {
  const response = await fetch(`/api/preferences/${id}`, { method: 'GET' });
  return await response.json();
};

// Ping function to test the server
export const pingServer = async () => {
  const response = await fetch('/api/ping', { method: 'GET' });
  return await response.json();
};
