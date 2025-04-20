
import { API_ENDPOINTS } from "@/config/api";
import { MatchFilters, MatchResult } from "@/types/matching";

export const fetchMatches = async (filters?: MatchFilters): Promise<MatchResult[]> => {
  try {
    // Build query params based on filters
    const params = new URLSearchParams();
    
    if (filters?.status && filters.status !== 'all') {
      params.append('status', filters.status);
    }
    
    if (filters?.searchTerm) {
      params.append('search', filters.searchTerm);
    }
    
    if (filters?.minConfidence !== undefined) {
      params.append('minConfidence', filters.minConfidence.toString());
    }
    
    const queryString = params.toString() ? `?${params.toString()}` : '';
    const response = await fetch(API_ENDPOINTS.MATCHINGS.GET_ALL);//${queryString}`);
    
    if (!response.ok) {
      throw new Error(`Error fetching matches: ${response.statusText}`);
    }
    
    const data = await response.json();
    return data.matchings || [];
  } catch (error) {
    console.error('Error in fetchMatches:', error);
    throw error;
  }
};

// Mock data for development/preview purposes
export const getMockMatches = (): MatchResult[] => {
  return [
    {
      id: '1',
      course: {
        ID: 'CS101',
        number: '101',
        name: 'Introduction to Computer Science',
        campus: 'Main Campus',
        semesters: ['Fall 2024', 'Spring 2025']
      },
      user: {
        ID: 'U123',
        name: 'John Doe',
        email: 'john@example.com'
      },
      status: 'perfect',
      confidence: 0.98,
      details: [
        { field: 'name', match: true, sourceValue: 'John Doe', targetValue: 'John Doe' },
        { field: 'email', match: true, sourceValue: 'john@example.com', targetValue: 'john@example.com' },
        { field: 'course', match: true, sourceValue: 'CS101', targetValue: 'CS101' }
      ],
      timestamp: '2025-04-05T12:30:00Z'
    },
    {
      id: '2',
      course: {
        ID: 'ENG200',
        number: '200',
        name: 'Advanced English Composition',
        campus: 'Arts Campus',
        semesters: ['Spring 2025']
      },
      user: {
        ID: 'U456',
        name: 'Jane Smith',
        email: 'jane@example.com'
      },
      status: 'partial',
      confidence: 0.75,
      details: [
        { field: 'name', match: true, sourceValue: 'Jane Smith', targetValue: 'Jane Smith' },
        { field: 'email', match: false, sourceValue: 'jane@example.com', targetValue: 'jsmith@example.com' },
        { field: 'course', match: true, sourceValue: 'ENG200', targetValue: 'ENG200' }
      ],
      timestamp: '2025-04-04T15:45:00Z'
    },
    {
      id: '3',
      course: {
        ID: 'MATH300',
        number: '300',
        name: 'Calculus III',
        campus: 'Science Campus',
        semesters: ['Fall 2024', 'Summer 2025']
      },
      user: {
        ID: 'U789',
        name: 'Robert Johnson',
        email: 'robert@example.com'
      },
      status: 'none',
      confidence: 0.32,
      details: [
        { field: 'name', match: false, sourceValue: 'Robert Johnson', targetValue: 'Rob Johnson' },
        { field: 'email', match: false, sourceValue: 'robert@example.com', targetValue: 'rob@example.com' },
        { field: 'course', match: false, sourceValue: 'MATH300', targetValue: 'MATH301' }
      ],
      timestamp: '2025-04-03T09:15:00Z'
    },
    {
      id: '4',
      course: {
        ID: 'PHYS101',
        number: '101',
        name: 'Introduction to Physics',
        campus: 'Science Campus',
        semesters: ['Fall 2024', 'Spring 2025']
      },
      user: {
        ID: 'U101',
        name: 'Emily Wilson',
        email: 'emily@example.com'
      },
      status: 'perfect',
      confidence: 1.0,
      details: [
        { field: 'name', match: true, sourceValue: 'Emily Wilson', targetValue: 'Emily Wilson' },
        { field: 'email', match: true, sourceValue: 'emily@example.com', targetValue: 'emily@example.com' },
        { field: 'course', match: true, sourceValue: 'PHYS101', targetValue: 'PHYS101' }
      ],
      timestamp: '2025-04-02T14:20:00Z'
    },
    {
      id: '5',
      course: {
        ID: 'BIO250',
        number: '250',
        name: 'Molecular Biology',
        campus: 'Health Campus',
        semesters: ['Spring 2025', 'Summer 2025']
      },
      user: {
        ID: 'U202',
        name: 'Michael Brown',
        email: 'michael@example.com'
      },
      status: 'partial',
      confidence: 0.82,
      details: [
        { field: 'name', match: true, sourceValue: 'Michael Brown', targetValue: 'Michael Brown' },
        { field: 'email', match: true, sourceValue: 'michael@example.com', targetValue: 'michael@example.com' },
        { field: 'course', match: false, sourceValue: 'BIO250', targetValue: 'BIO251' }
      ],
      timestamp: '2025-04-01T08:50:00Z'
    }
  ];
};
