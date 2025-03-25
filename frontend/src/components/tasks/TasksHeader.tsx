
import React from 'react';
import { Button } from '@/components/ui/button';
import { Plus } from 'lucide-react';

export const TasksHeader: React.FC = () => {
  return (
    <div className="flex flex-col md:flex-row md:items-center md:justify-between gap-4">
      <div>
        <h1 className="text-2xl font-bold tracking-tight">Tasks</h1>
        <p className="text-muted-foreground">Manage your tasks and projects</p>
      </div>
      
      <Button>
        <Plus className="mr-2 h-4 w-4" />
        New Task
      </Button>
    </div>
  );
};
