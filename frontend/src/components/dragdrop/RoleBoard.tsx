
import React, { useState, useCallback } from 'react';
import { Card, CardProps } from '../people/Card';
import { cn } from '@/lib/utils';

interface RoleBoardProps {
  title: string;
  people: CardProps[];
  className?: string;
  onPersonMove?: (personId: string, sourceRoleId: string, targetRoleId: string) => void;
  id: string;
}

export const RoleBoard: React.FC<RoleBoardProps> = ({ 
  title, 
  people, 
  className,
  onPersonMove,
  id
}) => {
  const [draggedPersonId, setDraggedPersonId] = useState<string | null>(null);
  const [isDropTarget, setIsDropTarget] = useState(false);
  
  const handleDragStart = useCallback((e: React.DragEvent, personId: string) => {
    // COMMENT: For Go migration, this would prepare data for a server request
    e.dataTransfer.setData('application/json', JSON.stringify({
      personId,
      sourceRoleId: id
    }));
    setDraggedPersonId(personId);
  }, [id]);
  
  const handleDragEnd = useCallback(() => {
    setDraggedPersonId(null);
  }, []);
  
  const handleDragOver = useCallback((e: React.DragEvent) => {
    e.preventDefault();
    setIsDropTarget(true);
  }, []);
  
  const handleDragLeave = useCallback(() => {
    setIsDropTarget(false);
  }, []);
  
  const handleDrop = useCallback((e: React.DragEvent) => {
    e.preventDefault();
    setIsDropTarget(false);
    
    try {
      const data = JSON.parse(e.dataTransfer.getData('application/json'));
      
      if (data && data.personId && data.sourceRoleId) {
        if (onPersonMove) {
          // COMMENT: For Go migration, this will be a server request to update 
          // the assignment in the database
          onPersonMove(data.personId, data.sourceRoleId, id);
        }
      }
    } catch (err) {
      console.error('Error parsing dropped data:', err);
    }
  }, [id, onPersonMove]);
  
  return (
    <div 
      className={cn(
        "border rounded-lg bg-card overflow-hidden transition-colors min-h-[200px] h-auto",
        isDropTarget && "ring-2 ring-primary/20 bg-primary/5",
        className
      )}
      onDragOver={handleDragOver}
      onDragLeave={handleDragLeave}
      onDrop={handleDrop}
      id={id}
    >
      <div className="border-b px-4 py-3 bg-muted/30">
        <h2 className="font-medium">{title}</h2>
        <p className="text-sm text-muted-foreground">{people.length} items</p>
      </div>
      
      <div className="p-3 space-y-3 min-h-[200px]">
        {people.length === 0 ? (
          <div className="h-full flex items-center justify-center text-sm text-muted-foreground p-6">
            Drop items here
          </div>
        ) : (
          people.map(person => (
            <div 
              key={person.id}
              draggable
              onDragStart={(e) => handleDragStart(e, person.id)}
              onDragEnd={handleDragEnd}
              className="drag-item cursor-grab active:cursor-grabbing"
            >
              <Card 
                {...person} 
              />
            </div>
          ))
        )}
      </div>
    </div>
  );
};
