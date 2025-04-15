
import React from 'react';
import { Card, CardHeader, CardTitle, CardDescription, CardContent } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Book } from "lucide-react";
import { Course } from "@/types/course";

interface CourseCardProps {
  course: Course;
}

const CourseCard: React.FC<CourseCardProps> = ({ course }) => {
  return (
    <Card className="h-full overflow-hidden hover:shadow-md transition-shadow">
      <CardHeader className="pb-2">
        <div className="flex justify-between items-start">
          <div className="flex items-center gap-2">
            <Book className="h-5 w-5 text-primary" />
            <CardTitle className="text-lg">{course.Name}</CardTitle>
          </div>
          <Badge variant="outline">
            {course.Campus}
          </Badge>
        </div>
        {course.Number && (
          <CardDescription>Course #: {course.Number}</CardDescription>
        )}
      </CardHeader>
      <CardContent>
        <div className="mt-2 text-xs text-muted-foreground">
          Semesters: {course.Semesters}
        </div>
      </CardContent>
    </Card>
  );
};

export default CourseCard;
