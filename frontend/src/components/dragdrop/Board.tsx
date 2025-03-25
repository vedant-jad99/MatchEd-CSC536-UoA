
import React from 'react';
import { useDragContext } from './DragContext';
import { Card } from '../people/Card';

// COMMENT: For Go migration, this component will need to be refactored
// to make API calls instead of using context state
const Board = ({ id, title, people }) => {
  const { movePerson } = useDragContext();
  
  const handleDragOver = (e) => {
    e.preventDefault();
    e.currentTarget.classList.add('active-drop');
  };
  
  const handleDragLeave = (e) => {
    e.currentTarget.classList.remove('active-drop');
  };
  
  const handleDrop = (e) => {
    e.preventDefault();
    e.currentTarget.classList.remove('active-drop');
    
    // COMMENT: For Go migration, this will be a server request to 
    // update the assignment in the database
    try {
      const data = JSON.parse(e.dataTransfer.getData('application/json'));
      if (data && data.personId && data.sourceRoleId) {
        movePerson(data.personId, data.sourceRoleId, id);
      }
    } catch (err) {
      console.error('Error parsing dropped data:', err);
    }
  };
  
  return (
    <div 
      className="board bg-card rounded-lg border shadow-sm overflow-hidden flex flex-col min-h-[200px] h-auto"
      onDragOver={handleDragOver}
      onDragLeave={handleDragLeave}
      onDrop={handleDrop}
    >
      <div className="bg-muted/40 px-4 py-3 border-b">
        <h3 className="font-medium">{title}</h3>
      </div>
      
      <div className="flex-1 p-3 space-y-3 overflow-y-auto">
        {people && people.length > 0 ? (
          people.map((person) => (
            <Card 
              key={person.id}
              {...person}
              sourceRoleId={id}
              draggable={true}
              onDragStart={(e) => {
                // COMMENT: For Go migration, this would prepare data for a server request
                e.dataTransfer.setData('application/json', JSON.stringify({
                  personId: person.id,
                  sourceRoleId: id
                }));
              }}
            />
          ))
        ) : (
          <div className="bg-muted/40 rounded-md p-3 text-center text-muted-foreground text-sm border border-dashed h-20 flex items-center justify-center">
            Drop here
          </div>
        )}
      </div>
    </div>
  );
};

export { Board };
