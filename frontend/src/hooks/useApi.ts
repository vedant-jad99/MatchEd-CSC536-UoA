
import { useCallback } from 'react';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import axios from 'axios';
import { API_ENDPOINTS, createHeaders } from '../config/api';

// Helper function using axios instead of fetch
const axiosWithAuth = async (url: string, options: any = {}) => {
  try {
    // Get auth token from localStorage
    const token = localStorage.getItem('authToken');
    
    // Set up default config
    const config = {
      ...options,
      url,
      headers: {
        ...createHeaders(token),
        ...(options.headers || {})
      }
    };
    
    const response = await axios(config);
    return response.data;
  } catch (error) {
    console.error('API axios error:', error);
    throw error;
  }
};

// Generic type for API responses
export interface ApiResponse<T> {
  data: T;
  message?: string;
  statusCode: number;
}

// Hook for handling API calls with React Query
export function useApi() {
  const queryClient = useQueryClient();

  // People API hooks
  const usePeople = <T>() => {
    return useQuery<T>({
      queryKey: ['people'],
      queryFn: async () => {
        return await axiosWithAuth(API_ENDPOINTS.PEOPLE.GET_ALL);
      }
    });
  };

  const usePersonById = <T>(id: string) => {
    return useQuery<T>({
      queryKey: ['people', id],
      queryFn: async () => {
        return await axiosWithAuth(API_ENDPOINTS.PEOPLE.GET_BY_ID(id));
      },
      enabled: !!id
    });
  };

  const useCreatePerson = <T, D>() => {
    return useMutation<T, Error, D>({
      mutationFn: async (personData: D) => {
        return await axiosWithAuth(API_ENDPOINTS.PEOPLE.CREATE, {
          method: 'POST',
          data: personData,
        });
      },
      onSuccess: () => {
        // Invalidate and refetch people queries
        queryClient.invalidateQueries({ queryKey: ['people'] });
      }
    });
  };

  const useUpdatePerson = <T, D>() => {
    return useMutation<T, Error, { id: string; data: D }>({
      mutationFn: async ({ id, data }) => {
        return await axiosWithAuth(API_ENDPOINTS.PEOPLE.UPDATE(id), {
          method: 'PUT',
          data,
        });
      },
      onSuccess: (_, variables) => {
        // Invalidate and refetch related queries
        queryClient.invalidateQueries({ queryKey: ['people'] });
        queryClient.invalidateQueries({ queryKey: ['people', variables.id] });
      }
    });
  };

  const useDeletePerson = <T>() => {
    return useMutation<T, Error, string>({
      mutationFn: async (id: string) => {
        return await axiosWithAuth(API_ENDPOINTS.PEOPLE.DELETE(id), {
          method: 'DELETE',
        });
      },
      onSuccess: () => {
        // Invalidate and refetch people queries
        queryClient.invalidateQueries({ queryKey: ['people'] });
      }
    });
  };

  // Courses API hooks
  const useCourses = <T>() => {
    return useQuery<T>({
      queryKey: ['courses'],
      queryFn: async () => {
        return await axiosWithAuth(API_ENDPOINTS.COURSES.GET_ALL);
      }
    });
  };

  const useCourseById = <T>(id: string) => {
    return useQuery<T>({
      queryKey: ['courses', id],
      queryFn: async () => {
        return await axiosWithAuth(API_ENDPOINTS.COURSES.GET_BY_ID(id));
      },
      enabled: !!id
    });
  };

  const useCreateCourse = <T, D>() => {
    return useMutation<T, Error, D>({
      mutationFn: async (courseData: D) => {
        return await axiosWithAuth(API_ENDPOINTS.COURSES.CREATE, {
          method: 'POST',
          data: courseData,
        });
      },
      onSuccess: () => {
        // Invalidate and refetch courses queries
        queryClient.invalidateQueries({ queryKey: ['courses'] });
      }
    });
  };

  // Roles/assignments API hooks
  const useAssignPerson = <T>() => {
    return useMutation<T, Error, { personId: string; roleId: string }>({
      mutationFn: async ({ personId, roleId }) => {
        return await axiosWithAuth(API_ENDPOINTS.ROLES.ASSIGN_PERSON, {
          method: 'POST',
          data: { personId, roleId },
        });
      },
      onSuccess: () => {
        // Invalidate and refetch relevant queries
        queryClient.invalidateQueries({ queryKey: ['people'] });
        queryClient.invalidateQueries({ queryKey: ['roles'] });
      }
    });
  };

  const useUnassignPerson = <T>() => {
    return useMutation<T, Error, { personId: string; roleId: string }>({
      mutationFn: async ({ personId, roleId }) => {
        return await axiosWithAuth(API_ENDPOINTS.ROLES.UNASSIGN_PERSON, {
          method: 'POST',
          data: { personId, roleId },
        });
      },
      onSuccess: () => {
        // Invalidate and refetch relevant queries
        queryClient.invalidateQueries({ queryKey: ['people'] });
        queryClient.invalidateQueries({ queryKey: ['roles'] });
      }
    });
  };

  const useRoles = <T>() => {
    return useQuery<T>({
      queryKey: ['roles'],
      queryFn: async () => {
        return await axiosWithAuth(API_ENDPOINTS.ROLES.GET_ALL);
      }
    });
  };

  // Generic fetcher for custom endpoints
  const useFetchData = <T>(url: string, options: any = {}) => {
    const fetchData = useCallback(async (): Promise<T> => {
      return await axiosWithAuth(url, options);
    }, [url, JSON.stringify(options)]);

    return fetchData;
  };

  return {
    // People hooks
    usePeople,
    usePersonById,
    useCreatePerson,
    useUpdatePerson,
    useDeletePerson,
    
    // Courses hooks
    useCourses,
    useCourseById,
    useCreateCourse,
    
    // Roles hooks
    useRoles,
    useAssignPerson,
    useUnassignPerson,
    
    // Generic fetcher
    useFetchData
  };
}
