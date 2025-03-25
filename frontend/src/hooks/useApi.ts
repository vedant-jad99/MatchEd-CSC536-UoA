
import { useCallback, useState } from 'react';
import { fetchWithAuth, API_ENDPOINTS } from '../config/api';

// Generic hook for API calls that can be used throughout the app
export function useApi<T, E = Error>() {
  const [data, setData] = useState<T | null>(null);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [error, setError] = useState<E | null>(null);

  // Generic fetch method
  const fetchData = useCallback(async (
    url: string,
    options: RequestInit = {}
  ): Promise<void> => {
    setIsLoading(true);
    setError(null);
    
    try {
      const result = await fetchWithAuth(url, options);
      setData(result);
    } catch (err) {
      setError(err as E);
    } finally {
      setIsLoading(false);
    }
  }, []);

  // Example for fetching all people
  const fetchPeople = useCallback(async () => {
    return fetchData(API_ENDPOINTS.PEOPLE.GET_ALL);
  }, [fetchData]);

  // Example for creating a new person
  const createPerson = useCallback(async (personData: any) => {
    return fetchData(API_ENDPOINTS.PEOPLE.CREATE, {
      method: 'POST',
      body: JSON.stringify(personData),
    });
  }, [fetchData]);

  // Similar functions can be created for other endpoints

  return {
    data,
    isLoading,
    error,
    fetchData,
    fetchPeople,
    createPerson,
  };
}
