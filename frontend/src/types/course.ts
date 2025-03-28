
export interface Course {
  Number: string;
  Name: string;
  Campus: string;
  Semesters: string;
}

export interface CourseResponse {
  courses: Course[];
  message?: string;
  status: string;
}
