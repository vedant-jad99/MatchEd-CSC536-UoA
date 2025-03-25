
import React from 'react';
import { Role } from '@/types/role';
import { priorityColorMap, statusColorMap } from '@/utils/roleUtils';
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';
import { Progress } from '@/components/ui/progress';
import { Button } from '@/components/ui/button';
import { Clock } from 'lucide-react';
import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar';

interface RoleCardProps {
  task: Role;
}

export const TaskCard: React.FC<RoleCardProps> = ({ task }) => {
  return (
    <Card className="animate-fade-in transition-all duration-200 hover:shadow-md">
      <CardHeader className="pb-2">
        <div className="flex justify-between items-start">
          <CardTitle className="text-lg">{task.title}</CardTitle>
          <Badge className={priorityColorMap[task.priority]} variant="outline">
            {task.priority}
          </Badge>
        </div>
        <CardDescription>{task.description}</CardDescription>
      </CardHeader>
      <CardContent className="pb-2">
        <div className="flex flex-wrap gap-2 mb-3">
          <Badge variant="secondary" className={statusColorMap[task.status]}>
            {task.status}
          </Badge>
          <Badge variant="outline" className="flex items-center gap-1">
            <Clock className="h-3 w-3" />
            Due: {new Date(task.dueDate).toLocaleDateString()}
          </Badge>
        </div>
        
        <div className="space-y-2">
          <div className="flex justify-between items-center text-sm">
            <span>Progress</span>
            <span>{task.progress}%</span>
          </div>
          <Progress value={task.progress} className="h-2" />
        </div>
      </CardContent>
      <CardFooter className="pt-2">
        <div className="flex justify-between items-center w-full">
          {task.assignee ? (
            <div className="flex items-center gap-2">
              <Avatar className="h-6 w-6">
                <AvatarFallback className="text-xs">
                  {task.assignee.name.substring(0, 2).toUpperCase()}
                </AvatarFallback>
                {task.assignee.avatar && (
                  <AvatarImage src={task.assignee.avatar} alt={task.assignee.name} />
                )}
              </Avatar>
              <span className="text-sm">{task.assignee.name}</span>
            </div>
          ) : (
            <span className="text-sm text-muted-foreground">Unassigned</span>
          )}
          
          <Button variant="ghost" size="sm" className="ml-auto">
            View
          </Button>
        </div>
      </CardFooter>
    </Card>
  );
};
