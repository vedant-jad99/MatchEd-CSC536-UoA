
import { useMatches } from "@/hooks/useMatches";
import MatchResultCard from "./MatchResultCard";
import MatchFilters from "./MatchFilters";
import { Button } from "@/components/ui/button";
import { RefreshCw, AlertTriangle, List } from "lucide-react";
import { useState } from "react";
import { Pagination } from "@/components/ui/pagination";
import { Separator } from "@/components/ui/separator";

const RESULTS_PER_PAGE = 15;

const MatchingResults = () => {
  const {
    matches,
    isLoading,
    isError,
    error,
    filters,
    setStatusFilter,
    setSearchFilter,
    setConfidenceFilter,
    refetch
  } = useMatches();
  
  const [currentPage, setCurrentPage] = useState(1);
  
  // Calculate pagination
  const totalPages = Math.ceil(matches.length / RESULTS_PER_PAGE);
  const startIdx = (currentPage - 1) * RESULTS_PER_PAGE;
  const endIdx = startIdx + RESULTS_PER_PAGE;
  const currentMatches = matches.slice(startIdx, endIdx);

  // Handle pagination
  const handlePageChange = (page: number) => {
    setCurrentPage(page);
    window.scrollTo({ top: 0, behavior: 'smooth' });
  };

  if (isError) {
    return (
      <div className="flex flex-col items-center justify-center p-8 bg-red-50 rounded-lg border border-red-200">
        <AlertTriangle className="h-12 w-12 text-red-500 mb-4" />
        <h2 className="text-xl font-semibold mb-2">Error Loading Matches</h2>
        <p className="text-gray-600 mb-4">{error?.message || "Failed to load matching results. Please try again."}</p>
        <Button onClick={() => refetch()} className="flex items-center gap-2">
          <RefreshCw className="h-4 w-4" />
          Retry
        </Button>
      </div>
    );
  }

  return (
    <div className="w-full max-w-4xl mx-auto p-4">
      <div className="mb-4">
        <h1 className="text-2xl font-bold mb-1">Matching Results</h1>
        <p className="text-muted-foreground text-sm">
          View and filter matching results between courses and users
        </p>
      </div>

      <div className="mb-4">
        <MatchFilters
          statusFilter={filters.status || 'all'}
          searchTerm={filters.searchTerm || ''}
          minConfidence={filters.minConfidence || 0}
          onStatusFilterChange={setStatusFilter}
          onSearchChange={setSearchFilter}
          onConfidenceChange={setConfidenceFilter}
        />
      </div>

      {isLoading ? (
        <div className="space-y-2">
          {[1, 2, 3, 4, 5].map((i) => (
            <div key={i} className="w-full h-14 rounded-md bg-gray-100 animate-pulse-light"></div>
          ))}
        </div>
      ) : matches.length > 0 ? (
        <>
          <div className="flex justify-between items-center mb-3">
            <div className="flex items-center gap-2">
              <List className="h-4 w-4 text-muted-foreground" />
              <p className="text-sm text-muted-foreground">
                Showing {startIdx + 1}-{Math.min(endIdx, matches.length)} of {matches.length} result{matches.length !== 1 ? 's' : ''}
              </p>
            </div>
            <Button onClick={() => refetch()} size="sm" variant="outline" className="h-7 text-xs flex items-center gap-1">
              <RefreshCw className="h-3 w-3" />
              Refresh
            </Button>
          </div>
          
          <div className="border rounded-md mb-4">
            {currentMatches.map((match, index) => (
              <div key={match.id}>
                <MatchResultCard match={match} />
                {index < currentMatches.length - 1 && <Separator className="mx-4" />}
              </div>
            ))}
          </div>
          
          {totalPages > 1 && (
            <div className="flex justify-center mt-4">
              <Pagination>
                <Pagination.First 
                  onClick={() => handlePageChange(1)} 
                  disabled={currentPage === 1}
                />
                <Pagination.Prev 
                  onClick={() => handlePageChange(currentPage - 1)} 
                  disabled={currentPage === 1}
                />
                {[...Array(totalPages)].map((_, i) => {
                  const page = i + 1;
                  // Show active page, first, last, and pages around current
                  const shouldShowPage = 
                    page === 1 || 
                    page === totalPages || 
                    (page >= currentPage - 1 && page <= currentPage + 1) ||
                    (currentPage <= 3 && page <= 5) ||
                    (currentPage >= totalPages - 2 && page >= totalPages - 4);
                    
                  return shouldShowPage ? (
                    <Pagination.Item
                      key={page}
                      onClick={() => handlePageChange(page)} 
                      active={page === currentPage}
                    >
                      {page}
                    </Pagination.Item>
                  ) : (page === 2 || page === totalPages - 1) ? (
                    <Pagination.Ellipsis key={`ellipsis-${page}`} />
                  ) : null;
                })}
                <Pagination.Next 
                  onClick={() => handlePageChange(currentPage + 1)} 
                  disabled={currentPage === totalPages}
                />
                <Pagination.Last 
                  onClick={() => handlePageChange(totalPages)} 
                  disabled={currentPage === totalPages}
                />
              </Pagination>
            </div>
          )}
        </>
      ) : (
        <div className="flex flex-col items-center justify-center p-6 bg-gray-50 rounded-lg border border-gray-200">
          <h3 className="text-base font-medium mb-1">No matches found</h3>
          <p className="text-gray-500 mb-3 text-sm text-center">
            Try adjusting your filters or search criteria
          </p>
          <Button onClick={() => {
            setStatusFilter('all');
            setSearchFilter('');
            setConfidenceFilter(0);
          }} variant="outline" size="sm" className="h-8 text-xs">
            Clear Filters
          </Button>
        </div>
      )}
    </div>
  );
};

export default MatchingResults;
