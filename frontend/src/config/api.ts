
// Configuration for API endpoints
// This will be used throughout the application to make requests to the Go/Gin backend

const API_BASE_URL = process.env.NODE_ENV === 'production' 
  ? '/api' // In production, served directly from Gin
  : '/api'; // In development, proxied through Vite

export const API_ENDPOINTS = {
  // Auth endpoints
  AUTH: {
    LOGIN: `${API_BASE_URL}/auth/login`,
    REGISTER: `${API_BASE_URL}/auth/register`,
    LOGOUT: `${API_BASE_URL}/auth/logout`,
  },
  // Users endpoints
  USERS: {
    GET_ALL: `${API_BASE_URL}/users`,
    GET_BY_ID: (id: string) => `${API_BASE_URL}/users/${id}`,
    CREATE: `${API_BASE_URL}/users`,
    UPDATE: (id: string) => `${API_BASE_URL}/users/${id}`,
    DELETE: (id: string) => `${API_BASE_URL}/users/${id}`,
  },
  // Courses endpoints
  COURSES: {
    GET_ALL: `${API_BASE_URL}/courses`,
    GET_BY_ID: (id: string) => `${API_BASE_URL}/courses/${id}`,
    CREATE: `${API_BASE_URL}/courses`,
    UPDATE: (id: string) => `${API_BASE_URL}/courses/${id}`,
    DELETE: (id: string) => `${API_BASE_URL}/courses/${id}`,
  },
  // People endpoints
  PEOPLE: {
    GET_ALL: `${API_BASE_URL}/people`,
    GET_BY_ID: (id: string) => `${API_BASE_URL}/people/${id}`,
    CREATE: `${API_BASE_URL}/people`,
    UPDATE: (id: string) => `${API_BASE_URL}/people/${id}`,
    DELETE: (id: string) => `${API_BASE_URL}/people/${id}`,
  },
  // Roles/assignments endpoints
  ROLES: {
    GET_ALL: `${API_BASE_URL}/roles`,
    ASSIGN_PERSON: `${API_BASE_URL}/roles/assign`,
    UNASSIGN_PERSON: `${API_BASE_URL}/roles/unassign`,
  }
};

// Helper function to create request headers with authentication
export const createHeaders = (token?: string) => {
  const headers: Record<string, string> = {
    'Content-Type': 'application/json',
  };
  
  if (token) {
    headers['Authorization'] = `Bearer ${token}`;
  }
  
  return headers;
};

// Fetch with authentication function for API requests
export const fetchWithAuth = async (url: string, options: RequestInit = {}): Promise<any> => {
  try {
    // Get auth token from localStorage (adjust based on your auth implementation)
    const token = localStorage.getItem('authToken');
    
    // Merge default headers with provided options
    const fetchOptions: RequestInit = {
      ...options,
      headers: {
        ...createHeaders(token),
        ...(options.headers || {})
      }
    };
    
    const response = await fetch(url, fetchOptions);
    
    // Check if the response is successful
    if (!response.ok) {
      throw new Error(`API error: ${response.status} ${response.statusText}`);
    }
    
    // Parse JSON response
    const data = await response.json();
    return data;
  } catch (error) {
    console.error('API fetch error:', error);
    throw error;
  }
};
