
import React from 'react';
import { MainLayout } from '@/components/layout/MainLayout';
import { Board } from '@/components/dragdrop/Board';
import { DragProvider, useDragContext } from '@/components/dragdrop/DragContext';
import { Users, Plus } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { Separator } from '@/components/ui/separator';
import { CardProps } from '@/components/people/Card';

// Mock data for people
const initialRoles = [
  {
    id: 'role-1',
    title: '',
    people: [
      {
        id: 'person-1',
        name: 'John Doe',
        assigned: 'Frontend Developers',
        details: 'Senior developer with 5 years of React experience. Currently working on the dashboard redesign project.',
        color: 'bg-blue-500',
        type: 'human',
        status: 'active'
      } as CardProps,
      {
        id: 'person-2',
        name: 'Sarah Smith',
        image: '/placeholder.svg',
        assigned: 'Frontend Developers',
        details: 'UI/UX specialist with expertise in accessible design patterns.',
        color: 'bg-purple-500',
        type: 'human',
        status: 'active'
      } as CardProps,
    ]
  },
  {
    id: 'role-2',
    title: 'CSC 200',
    people: [
      {
        id: 'person-3',
        name: 'Michael Chen',
        assigned: 'Backend Developers',
        details: 'Database expert with a focus on performance optimization.',
        color: 'bg-green-500',
        type: 'human',
        status: 'active'
      } as CardProps,
    ]
  },
  {
    id: 'role-3',
    title: 'CSC 150',
    people: [
      {
        id: 'person-4',
        name: 'Lisa Johnson',
        assigned: 'Project Managers',
        details: 'PMP certified with 8 years of experience managing agile teams.',
        color: 'bg-amber-500',
        type: 'human',
        status: 'active'
      } as CardProps,
    ]
  },
  {
    id: 'role-4',
    title: 'Unassigned',
    people: [
      {
        id: 'person-5',
        name: 'David Williams',
        details: 'Full-stack developer looking for project assignment.',
        color: 'bg-red-500',
        type: 'human',
        status: 'inactive'
      } as CardProps,
      {
        id: 'person-6',
        name: 'Emily Davis',
        details: 'New hire with strong React Native background.',
        color: 'bg-cyan-500',
        type: 'human',
        status: 'inactive'
      } as CardProps,
    ]
  },
];

// Component for the container that will use the context
const PeopleContainer = () => {
  const { movePerson } = useDragContext();
  
  return (
    <div className="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-4 gap-4"
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
      {initialRoles.map(role => (
        <Board
          key={role.id}
          id={role.id}
          title={role.title}
          people={role.people}
        />
      ))}
    </div>
  );
};

const People = () => {
  return (
    <MainLayout>
      <div className="space-y-6">
        <div className="flex flex-col md:flex-row md:items-center md:justify-between gap-4">
          <div>
            <h1 className="text-2xl font-bold tracking-tight">People</h1>
            <p className="text-muted-foreground">Manage team members and role assignments</p>
          </div>
          
          <Button>
            <Plus className="mr-2 h-4 w-4" />
            Add Person
          </Button>
        </div>
        
        <Separator />
        
        <div className="flex items-center gap-2 text-muted-foreground">
          <Users className="h-4 w-4" />
          <span className="text-sm">Drag people between roles to reassign them</span>
        </div>
        
        <DragProvider initialRoles={initialRoles}>
          <PeopleContainer />
        </DragProvider>
      </div>
    </MainLayout>
  );
};

export default People;
