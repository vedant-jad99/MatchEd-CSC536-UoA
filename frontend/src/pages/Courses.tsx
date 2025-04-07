
import React from 'react';
import { useQuery } from '@tanstack/react-query';
import axios from 'axios';
import { MainLayout } from '@/components/layout/MainLayout';
import { DragProvider } from '@/components/dragdrop/DragContext';
import { Board } from '@/components/dragdrop/Board';
import { CardProps } from '@/components/dragdrop/Card';
import { useDragContext } from '@/components/dragdrop/DragContext';
import { API_ENDPOINTS } from '@/config/api';
import { Course, CourseResponse } from '@/types/course';
import CourseCard from '@/components/courses/CourseCard';
import CoursesContainer from '@/components/courses/CoursesContainer';
import { Button } from '@/components/ui/button';
import { Plus, RefreshCw } from 'lucide-react';
import { useToast } from '@/hooks/use-toast';


const Courses = () => {
  return (
    <MainLayout>
      <div className="space-y-6">
        <div className="flex flex-col md:flex-row md:items-center md:justify-between gap-4">
          <div>
            <h1 className="text-2xl font-bold tracking-tight">Courses</h1>
            <p className="text-muted-foreground">Manage course assignments and distribution</p>
          </div>
          <Button className="self-start">
            <Plus className="mr-2 h-4 w-4" />
            Add New Course
          </Button>
        </div>
        
        <DragProvider initialRoles={[]}>
          <CoursesContainer/>
        </DragProvider>
      </div>
    </MainLayout>
  );
};

export default Courses;
