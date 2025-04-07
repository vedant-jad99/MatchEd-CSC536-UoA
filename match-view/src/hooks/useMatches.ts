
import { useQuery } from "@tanstack/react-query";
import { useState } from "react";
import { fetchMatches, getMockMatches } from "@/services/matchingService";
import { MatchFilters, MatchStatus } from "@/types/matching";

// Set this to true to use mock data instead of actual API calls
const USE_MOCK_DATA = true;

export function useMatches() {
  const [filters, setFilters] = useState<MatchFilters>({
    status: 'all',
    searchTerm: '',
    minConfidence: 0
  });

  const matchesQuery = useQuery({
    queryKey: ['matches', filters],
    queryFn: () => USE_MOCK_DATA ? getMockMatches() : fetchMatches(filters),
    staleTime: 1000 * 60 * 5, // 5 minutes
  });

  const updateFilters = (newFilters: Partial<MatchFilters>) => {
    setFilters(prev => ({
      ...prev,
      ...newFilters
    }));
  };
  
  const setStatusFilter = (status: MatchStatus | 'all') => {
    updateFilters({ status });
  };

  const setSearchFilter = (searchTerm: string) => {
    updateFilters({ searchTerm });
  };
  
  const setConfidenceFilter = (minConfidence: number) => {
    updateFilters({ minConfidence });
  };

  // Apply client-side filtering for mock data
  let filteredMatches = matchesQuery.data || [];
  
  if (USE_MOCK_DATA) {
    if (filters.status && filters.status !== 'all') {
      filteredMatches = filteredMatches.filter(match => match.status === filters.status);
    }
    
    if (filters.searchTerm) {
      const searchLower = filters.searchTerm.toLowerCase();
      filteredMatches = filteredMatches.filter(match => 
        match.course.name.toLowerCase().includes(searchLower) || 
        match.user.name.toLowerCase().includes(searchLower) ||
        match.user.email.toLowerCase().includes(searchLower) ||
        match.course.ID.toLowerCase().includes(searchLower)
      );
    }
    
    if (filters.minConfidence !== undefined && filters.minConfidence > 0) {
      filteredMatches = filteredMatches.filter(match => 
        match.confidence >= filters.minConfidence!
      );
    }
  }

  return {
    matches: filteredMatches,
    isLoading: matchesQuery.isLoading,
    isError: matchesQuery.isError,
    error: matchesQuery.error,
    filters,
    setStatusFilter,
    setSearchFilter,
    setConfidenceFilter,
    refetch: matchesQuery.refetch
  };
}
