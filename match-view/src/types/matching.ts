
export type MatchStatus = 'perfect' | 'partial' | 'none';

export interface Course {
  ID: string;
  number: string;
  name: string;
  campus: string;
  semesters: string[];
}

export interface User {
  ID: string;
  name: string;
  email: string;
}

export interface MatchResult {
  id: string;
  course: Course;
  user: User;
  status: MatchStatus;
  confidence: number;
  details?: {
    field: string;
    match: boolean;
    sourceValue?: string;
    targetValue?: string;
  }[];
  timestamp: string;
}

export interface MatchFilters {
  status?: MatchStatus | 'all';
  searchTerm?: string;
  minConfidence?: number;
}
