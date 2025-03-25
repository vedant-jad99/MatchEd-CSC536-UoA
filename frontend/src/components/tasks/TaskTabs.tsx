
import React from 'react';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import { Role } from '@/types/role';
import { TaskCard } from './TaskCard';
import { EmptyTasksView } from './EmptyTasksView';
import { filterRolesByStatus } from '@/utils/roleUtils';

interface RoleTabsProps {
  tasks: Role[];
}

export const TaskTabs: React.FC<RoleTabsProps> = ({ tasks }) => {
  const [activeTab, setActiveTab] = React.useState<'All' | Role['status']>('All');
  const filteredTasks = filterRolesByStatus(tasks, activeTab);
  
  return (
    <Tabs defaultValue="All" value={activeTab} onValueChange={(value) => setActiveTab(value as any)}>
      <div className="flex justify-between items-center">
        <TabsList>
          <TabsTrigger value="All">All</TabsTrigger>
          <TabsTrigger value="To Do">To Do</TabsTrigger>
          <TabsTrigger value="In Progress">In Progress</TabsTrigger>
          <TabsTrigger value="Completed">Completed</TabsTrigger>
        </TabsList>
        
        <div className="text-sm text-muted-foreground">
          Showing {filteredTasks.length} of {tasks.length} tasks
        </div>
      </div>
      
      <TabsContent value={activeTab} className="mt-6">
        {filteredTasks.length > 0 ? (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            {filteredTasks.map((task) => (
              <TaskCard key={task.id} task={task} />
            ))}
          </div>
        ) : (
          <EmptyTasksView activeTab={activeTab} />
        )}
      </TabsContent>
    </Tabs>
  );
};
