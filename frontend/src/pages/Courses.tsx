
import React from 'react';
import { MainLayout } from '@/components/layout/MainLayout';
import { DragProvider } from '@/components/dragdrop/DragContext';
import { Board } from '@/components/dragdrop/Board';
import { CardProps } from '@/components/people/Card';
import { useDragContext } from '@/components/dragdrop/DragContext';

// Sample data
const initialRoles = [
  {
    id: 'role-1',
    title: 'Faculty',
    people: [
      { 
        id: 'course', 
        name: 'John Smith', 
        assigned: 'Frontend Developer', 
        details: 'Specializes in React and TypeScript. 5 years of experience.',
        type: 'human',
        status: 'active'
      } as CardProps,
      { 
        id: 'person-2', 
        name: 'Alice Johnson', 
        assigned: 'Backend Developer', 
        details: 'Specializes in Node.js and MongoDB. 3 years of experience.',
        type: 'human',
        status: 'inactive'
      } as CardProps,
    ],
  },
  {
    id: 'role-2',
    title: 'Faculty',
    people: [
      { 
        id: 'person-3', 
        name: 'Emma Davis', 
        assigned: 'UI Designer', 
        details: 'Specializes in user interface design. 4 years of experience.',
        type: 'human',
        status: 'active'
      } as CardProps,
    ],
  },
  {
    id: 'role-3',
    title: 'Courses',
    people: [
      { 
        id: 'course-1', 
        name: 'Web Development', 
        assigned: 'Core Course', 
        details: 'Covers HTML, CSS, JavaScript and responsive design.',
        type: 'book',
        status: 'active'
      } as CardProps,
      { 
        id: 'course-2', 
        name: 'Database Design', 
        assigned: 'Elective Course', 
        details: 'Covers relational databases, SQL, and data modeling.',
        type: 'book',
        status: 'inactive'
      } as CardProps,
    ],
  },
  {
    id: 'role-4',
    title: 'Unassigned',
    people: [],
  },
];

const CoursesContainer = () => {
  const { movePerson } = useDragContext();
  
  return (
    <div className="border rounded-lg p-4 bg-muted/20" 
        onDragOver={(e) => e.preventDefault()}
        onDrop={(e) => {
          e.preventDefault();
          // COMMENT: For Go migration, this will be a server request to 
          // update the assignment status in the database
          try {
            const data = JSON.parse(e.dataTransfer.getData('application/json'));
            if (data && data.personId && data.sourceRoleId) {
              // Find the unassigned role and move the person there
              const unassignedRole = initialRoles.find(role => role.title === 'Unassigned');
              if (unassignedRole) {
                movePerson(data.personId, data.sourceRoleId, unassignedRole.id);
              }
            }
          } catch (err) {
            console.error('Error parsing dropped data:', err);
          }
        }}>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        {initialRoles.map((role) => (
          <Board
            key={role.id}
            id={role.id}
            title={role.title}
            people={role.people}
          />
        ))}
      </div>
    </div>
  );
};

const Courses = () => {
  return (
    <MainLayout>
      <div className="space-y-6">
        <div className="flex flex-col md:flex-row md:items-center md:justify-between gap-4">
          <div>
            <h1 className="text-2xl font-bold tracking-tight">Courses</h1>
            <p className="text-muted-foreground">Manage course assignments and distribution</p>
          </div>
        </div>
        
        <DragProvider initialRoles={initialRoles}>
          <CoursesContainer />
        </DragProvider>
      </div>
    </MainLayout>
  );
};

export default Courses;
