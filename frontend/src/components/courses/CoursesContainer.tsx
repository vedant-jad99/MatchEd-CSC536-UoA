import { useToast } from '@/hooks/use-toast';
import { useDragContext } from '@/components/dragdrop/DragContext';
import { useQuery } from '@tanstack/react-query';
import axios from 'axios';
import { API_ENDPOINTS } from '@/config/api';
import { Plus, RefreshCw } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { Board } from '@/components/dragdrop/Board';
import { Course, CourseResponse } from '@/types/course';

const CoursesContainer = () => {
  console.log("Rendering CoursesContainer...");
  const { movePerson } = useDragContext();
  const { toast } = useToast();
  
  // Fetch courses from API using Axios
  const { data, isLoading, isError, error, refetch } = useQuery<CourseResponse>({
    queryKey: ['courses'],
    queryFn: async () => {
      console.warn("getting courses")
      try {
        const response = await axios.get(API_ENDPOINTS.COURSES.GET_ALL);
        return response.data;
      } catch (err) {
        console.error('Error fetching courses:', err);
        toast({
          title: 'Error',
          description: 'Failed to fetch courses. Please try again.',
          variant: 'destructive'
        });
        throw err;
      }
    },
    staleTime: 0,  // Forces fresh fetch
  });

  // uses the data.courses to create an array of objects
  const courseBoardData = data?.courses
  ? data.courses.map((course) => ({
      id: `course-${course.ID}`,
      title: course.Name, // Set board titles as course names
      people: [], // Empty initially, will hold faculty members later
      
    }))
  : [];
  
  if (isLoading) {
    return (
      <div className="flex justify-center items-center h-64">
        <RefreshCw className="animate-spin h-8 w-8 text-primary" />
        <span className="ml-2">Loading courses...</span>
      </div>
    );
  }

  if (isError) {
    return (
      <div className="border rounded-lg p-4 bg-destructive/10 text-destructive">
        <h3 className="font-semibold">Error loading courses</h3>
        <p>{(error as Error)?.message || 'An unknown error occurred'}</p>
        <Button onClick={() => refetch()} variant="outline" className="mt-2">
          Try Again
        </Button>
      </div>
    );
  }
  
  return (
    <div className="border rounded-lg p-4 bg-muted/20" 
        onDragOver={(e) => e.preventDefault()}
        onDrop={(e) => {
          e.preventDefault();
          try {
            const data = JSON.parse(e.dataTransfer.getData('application/json'));
            if (data && data.personId && data.sourceRoleId) {
              // Find the unassigned role and move the person there
              const unassignedRole = courseBoardData.find(course => course.title === 'Filtered Courses');
              if (unassignedRole) {
                movePerson(data.personId, data.sourceRoleId, unassignedRole.id);
              }
            }
          } catch (err) {
            console.error('Error parsing dropped data:', err);
          }
        }}>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
      {courseBoardData.map((courseBoard) => (
        <Board
          key={courseBoard.id}
          id={courseBoard.id}
          title={courseBoard.title}
          people={courseBoard.people} // Empty for now, faculty will go here later
        />
      ))}
      </div>
    </div>
  );
};

export default CoursesContainer