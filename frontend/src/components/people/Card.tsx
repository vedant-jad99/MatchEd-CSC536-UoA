
import React, { useState } from 'react';
import { ChevronDown, ChevronUp, Eye, EyeOff } from 'lucide-react';
import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar';
import { Button } from '@/components/ui/button';
import { cn } from '@/lib/utils';
import { Badge } from '../ui/badge';

// COMMENT: For Go migration, this interface will become a Go struct
// and will be used to define database schema
export interface CardProps {
  id: string;
  name: string;
  image?: string;
  assigned?: string;
  details?: string;
  color?: string;
  className?: string;
  type: 'human' | 'book';
  status: 'active' | 'inactive';
  sourceRoleId?: string;
  draggable?: boolean;
  onDragStart?: (e: React.DragEvent) => void;
}

// COMMENT: For Go migration, this component will receive data from API calls
// The template will be rendered server-side or via a frontend framework
export const Card: React.FC<CardProps> = ({ 
  id, 
  name, 
  image, 
  assigned, 
  details,
  color = 'bg-primary',
  className,
  type,
  status,
  draggable = false,
  onDragStart
}) => {
  // COMMENT: Client-side state will need to be handled differently in Go
  // Either via htmx or a combination of Go templates and JavaScript
  const [expanded, setExpanded] = useState(false);
  const [visible, setVisible] = useState(true);
  
  // Generate initials from name
  const initials = name
    .split(' ')
    .map(part => part[0])
    .join('')
    .toUpperCase()
    .substring(0, 2);
  
  const statusColor = status === 'active' ? 'text-yellow-500' : 'text-red-500';
  const iconColor = `${statusColor} h-5 w-5`;
  
  return (
    <div 
      id={id}
      className={cn(
        "border rounded-lg p-4 bg-card shadow-sm transition-all",
        expanded ? "shadow-md" : "",
        draggable ? "cursor-grab active:cursor-grabbing" : "",
        className
      )}
      draggable={draggable}
      onDragStart={onDragStart}
    >
      <div className="flex items-center gap-3">
        <Avatar className={cn("h-14 w-14 border-2 border-muted", statusColor)}>
          {image ? (
            <AvatarImage src={image} alt={name} />
          ) : (
            <AvatarFallback className={cn("text-primary-foreground", color)}>
              {initials}
            </AvatarFallback>
          )}
        </Avatar>
        
        <div className="flex-1 min-w-0">
          <h3 className="font-medium text-foreground truncate">{name}</h3>
          {assigned && (
            <p className="text-sm text-muted-foreground truncate">
              {assigned}
            </p>
          )}
          <div className="flex mt-1 space-x-1">
            <Badge variant="outline" className={status === 'active' ? 'bg-yellow-100 text-yellow-800 dark:bg-yellow-900 dark:text-yellow-300' : 'bg-red-100 text-red-800 dark:bg-red-900 dark:text-red-300'}>
              {type === 'human' ? 'Person' : 'Course'}
            </Badge>
          </div>
        </div>
        
        <div className="flex flex-col gap-2">
          <Button 
            variant="ghost" 
            size="icon" 
            onClick={() => setVisible(!visible)}
            aria-label={visible ? "Hide item" : "Show item"}
            className="h-8 w-8"
          >
            {visible ? (
              <Eye className="h-4 w-4" />
            ) : (
              <EyeOff className="h-4 w-4" />
            )}
          </Button>
          
          {details && (
            <Button 
              variant="ghost" 
              size="icon" 
              onClick={() => setExpanded(!expanded)}
              aria-label={expanded ? "Hide details" : "Show details"}
              className="h-8 w-8"
            >
              {expanded ? (
                <ChevronUp className="h-4 w-4" />
              ) : (
                <ChevronDown className="h-4 w-4" />
              )}
            </Button>
          )}
        </div>
      </div>
      
      {expanded && details && (
        <div className="mt-3 pt-3 border-t text-sm">
          {details}
        </div>
      )}
    </div>
  );
};
