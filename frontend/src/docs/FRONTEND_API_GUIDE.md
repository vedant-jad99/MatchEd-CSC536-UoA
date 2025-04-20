
# Frontend API Integration Guide

This guide explains how to make API requests from the frontend to the Go/Gin backend server.

## API Configuration

All API endpoints are configured in `src/config/api.ts`. This is where you'll find the base URL and all available endpoints.

```typescript
// Example of API_ENDPOINTS structure
export const API_ENDPOINTS = {
  COURSES: {
    GET_ALL: `/api/courses`,
    GET_BY_ID: (id: string) => `/api/courses/${id}`,
    // more endpoints...
  },
  // other resource endpoints...
};
```

## Making API Requests with Axios

We use Axios for making HTTP requests to our Go/Gin backend.

### Direct Axios Usage

```typescript
// Example: Fetch all courses
const response = await axios.get(API_ENDPOINTS.COURSES.GET_ALL);
const courses = response.data;

// Example: Create a new course
const response = await axios.post(API_ENDPOINTS.COURSES.CREATE, courseData);
const newCourse = response.data;
```

### Authentication with Axios

For authenticated requests, you can use the `axiosWithAuth` helper from our API utils:

```typescript
import { axiosWithAuth } from '../config/api';

// Making an authenticated request
const data = await axiosWithAuth(API_ENDPOINTS.COURSES.GET_ALL);
```

Or, configure Axios with an interceptor:

```typescript
// Set up axios with auth token
axios.interceptors.request.use(config => {
  const token = localStorage.getItem('authToken');
  if (token) {
    config.headers['Authorization'] = `Bearer ${token}`;
  }
  return config;
});
```

## Using React Query with Axios

For most API requests, we use React Query (TanStack Query) with Axios for data fetching, caching, and state management.

### Setup Query Hook:

```typescript
// Example: Fetch all courses
const { data, isLoading, isError, error } = useQuery({
  queryKey: ['courses'],
  queryFn: async () => {
    const response = await axios.get(API_ENDPOINTS.COURSES.GET_ALL);
    return response.data;
  }
});
```

### Setup Mutation Hook:

```typescript
// Example: Create a new course
const createCourseMutation = useMutation({
  mutationFn: async (courseData) => {
    const response = await axios.post(API_ENDPOINTS.COURSES.CREATE, courseData);
    return response.data;
  },
  onSuccess: () => {
    // Invalidate and refetch queries
    queryClient.invalidateQueries({ queryKey: ['courses'] });
  }
});

// Usage:
createCourseMutation.mutate(newCourseData);
```

## Using the useApi Hook

We've created a custom `useApi` hook that simplifies API interactions using React Query:

```typescript
import { useApi } from '@/hooks/useApi';

function MyComponent() {
  const { useCourses } = useApi();
  
  // Get all courses
  const { data, isLoading, isError } = useCourses<CourseResponse>();
  
  // ... rest of your component
}
```

## GO/Gin Backend Integration

Our frontend is designed to work seamlessly with a Go/Gin backend:

1. **JSON Serialization**: Go/Gin returns JSON responses that match our TypeScript interfaces
2. **Authentication**: JWT tokens are sent in the Authorization header
3. **Error Handling**: Error responses from Go/Gin include status code and message

## Type Definitions

Define TypeScript interfaces for your API request and response objects in the `src/types` directory.

```typescript
// Example: Course interface matching Go GORM model
export interface Course {
  Number: string;  // matches GORM model field
  Name: string;    // matches GORM model field
  Campus: string;  // matches GORM model field
  Semesters: string; // matches GORM model field
}

export interface CourseResponse {
  courses: Course[];
  message?: string;
  status: string;
}
```

## Error Handling

Handle API request errors in your components:

```typescript
if (isError) {
  return <div>Error: {error.message}</div>;
}

if (isLoading) {
  return <div>Loading...</div>;
}
```

## Testing API Endpoints

Before integrating an API endpoint in your component, test it using tools like:
- The browser's Network tab in Developer Tools
- Postman
- curl

Example curl command:
```bash
curl -X GET http://localhost:3000/api/courses
```

## Authentication

For authenticated requests, ensure users are authenticated before making protected API calls:

```typescript
// Check if user is authenticated
const isAuthenticated = !!localStorage.getItem('authToken');

// Redirect if not authenticated
if (!isAuthenticated) {
  // Redirect to login page
}
```
