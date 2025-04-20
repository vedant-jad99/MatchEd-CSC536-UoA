
import { useState } from "react";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { ToggleGroup, ToggleGroupItem } from "@/components/ui/toggle-group";
import { Search, X, SlidersHorizontal, CheckCircle2, AlertCircle, XCircle } from "lucide-react";
import { Slider } from "@/components/ui/slider";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { Label } from "@/components/ui/label";
import { MatchStatus } from "@/types/matching";

interface MatchFiltersProps {
  statusFilter: string;
  searchTerm: string;
  minConfidence: number;
  onStatusFilterChange: (status: MatchStatus | 'all') => void;
  onSearchChange: (search: string) => void;
  onConfidenceChange: (confidence: number) => void;
}

const MatchFilters = ({
  statusFilter,
  searchTerm,
  minConfidence,
  onStatusFilterChange,
  onSearchChange,
  onConfidenceChange
}: MatchFiltersProps) => {
  const [isFilterExpanded, setIsFilterExpanded] = useState(false);
  const [localSearch, setLocalSearch] = useState(searchTerm);
  
  const handleSearchSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSearchChange(localSearch);
  };

  const clearSearch = () => {
    setLocalSearch('');
    onSearchChange('');
  };

  return (
    <div className="w-full space-y-4">
      <div className="flex flex-col sm:flex-row gap-2">
        <form onSubmit={handleSearchSubmit} className="relative flex-1">
          <Search className="absolute left-2.5 top-2.5 h-4 w-4 text-muted-foreground" />
          <Input
            type="search"
            placeholder="Search by name..."
            className="pl-8 pr-10"
            value={localSearch}
            onChange={(e) => setLocalSearch(e.target.value)}
          />
          {localSearch && (
            <Button
              variant="ghost"
              size="sm"
              className="absolute right-0 top-0 h-full px-2"
              onClick={clearSearch}
              type="button"
            >
              <X className="h-4 w-4" />
              <span className="sr-only">Clear search</span>
            </Button>
          )}
        </form>

        <Popover open={isFilterExpanded} onOpenChange={setIsFilterExpanded}>
          <PopoverTrigger asChild>
            <Button variant="outline" className="flex gap-1">
              <SlidersHorizontal className="h-4 w-4" />
              <span className="hidden sm:inline">Filters</span>
            </Button>
          </PopoverTrigger>
          <PopoverContent className="w-80">
            <div className="space-y-4">
              <div className="space-y-2">
                <Label htmlFor="min-confidence">Minimum Confidence ({minConfidence * 100}%)</Label>
                <Slider
                  id="min-confidence"
                  defaultValue={[minConfidence]}
                  max={1}
                  step={0.05}
                  onValueChange={(values) => onConfidenceChange(values[0])}
                />
              </div>

              <Button onClick={() => setIsFilterExpanded(false)} className="w-full">
                Apply Filters
              </Button>
            </div>
          </PopoverContent>
        </Popover>
      </div>
      
      <div className="flex justify-center sm:justify-start">
        <ToggleGroup type="single" value={statusFilter} onValueChange={(value) => {
          if (value) onStatusFilterChange(value as MatchStatus | 'all');
        }}>
          <ToggleGroupItem value="all">
            All Results
          </ToggleGroupItem>
          <ToggleGroupItem value="perfect" className="flex items-center gap-1">
            <CheckCircle2 className="h-3 w-3 text-matching-perfect" />
            <span className="hidden sm:inline">Perfect</span>
          </ToggleGroupItem>
          <ToggleGroupItem value="partial" className="flex items-center gap-1">
            <AlertCircle className="h-3 w-3 text-matching-partial" />
            <span className="hidden sm:inline">Partial</span>
          </ToggleGroupItem>
          <ToggleGroupItem value="none" className="flex items-center gap-1">
            <XCircle className="h-3 w-3 text-matching-none" />
            <span className="hidden sm:inline">No Match</span>
          </ToggleGroupItem>
        </ToggleGroup>
      </div>
    </div>
  );
};

export default MatchFilters;
