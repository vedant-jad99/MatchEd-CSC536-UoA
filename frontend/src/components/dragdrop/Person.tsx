
import React, { useState } from 'react';
import { ChevronDown, ChevronUp, User } from 'lucide-react';
import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar';
import { Button } from '@/components/ui/button';
import { cn } from '@/lib/utils';

export interface PersonProps {
  id: string;
  name: string;
  image?: string;
  assigned?: string;
  details?: string;
  color?: string;
  className?: string;
  isDragging?: boolean;
}

export const Person: React.FC<PersonProps> = ({ 
  id, 
  name, 
  image, 
  assigned, 
  details,
  color = 'bg-primary',
  className,
  isDragging
}) => {
  const [expanded, setExpanded] = useState(false);
  
  // Generate initials from name
  const initials = name
    .split(' ')
    .map(part => part[0])
    .join('')
    .toUpperCase()
    .substring(0, 2);
  
  return (
    <div 
      id={id}
      className={cn(
        "border rounded-lg p-4 bg-card shadow-sm transition-all",
        isDragging ? "opacity-50" : "opacity-100",
        expanded ? "shadow-md" : "",
        className
      )}
    >
      <div className="flex items-center gap-3">
        <Avatar className="h-14 w-14 border-2 border-muted">
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
        </div>
        
        {details && (
          <Button 
            variant="ghost" 
            size="icon" 
            onClick={() => setExpanded(!expanded)}
            aria-label={expanded ? "Hide details" : "Show details"}
          >
            {expanded ? (
              <ChevronUp className="h-4 w-4" />
            ) : (
              <ChevronDown className="h-4 w-4" />
            )}
          </Button>
        )}
      </div>
      
      {expanded && details && (
        <div className="mt-3 pt-3 border-t text-sm">
          {details}
        </div>
      )}
    </div>
  );
};
