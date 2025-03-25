
import React from 'react';
import { Button } from '@/components/ui/button';
import { CheckCircle, Plus } from 'lucide-react';

interface EmptyTasksViewProps {
  activeTab: 'All' | string;
}

export const EmptyTasksView: React.FC<EmptyTasksViewProps> = ({ activeTab }) => {
  return (
    <div className="flex flex-col items-center justify-center py-12 text-center">
      <div className="rounded-full bg-primary/10 p-4 mb-4">
        <CheckCircle className="h-8 w-8 text-primary" />
      </div>
      <h3 className="text-lg font-medium mb-1">No tasks found</h3>
      <p className="text-muted-foreground mb-4">
        {activeTab === 'All' 
          ? "You don't have any tasks yet." 
          : `You don't have any ${activeTab.toLowerCase()} tasks.`}
      </p>
      <Button>
        <Plus className="mr-2 h-4 w-4" />
        Create a new task
      </Button>
    </div>
  );
};
