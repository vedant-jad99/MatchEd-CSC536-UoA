
import React, { createContext, useContext, useState, ReactNode } from 'react';
import { CardProps } from '../people/Card';

// COMMENT: For Go migration, this interface will become a Go struct
// and the data will be stored in a database instead of in-memory state
interface RoleData {
  id: string;
  title: string;
  people: CardProps[];
}

// COMMENT: For Go migration, these operations will be handled by
// API endpoints that update the database
interface DragContextType {
  roles: RoleData[];
  movePerson: (personId: string, sourceRoleId: string, targetRoleId: string) => void;
}

const DragContext = createContext<DragContextType | undefined>(undefined);

interface DragProviderProps {
  initialRoles: RoleData[];
  children: ReactNode;
}

// COMMENT: For Go migration, this component will be replaced by
// API calls to fetch and update data from the Gin backend
export const DragProvider: React.FC<DragProviderProps> = ({ 
  initialRoles, 
  children 
}) => {
  // COMMENT: State will be moved to a database in Go migration
  const [roles, setRoles] = useState<RoleData[]>(initialRoles);
  
  // COMMENT: This function will become a Gin API endpoint in Go
  // that updates the database and returns the updated data
  const movePerson = (personId: string, sourceRoleId: string, targetRoleId: string) => {
    if (sourceRoleId === targetRoleId) return;
    
    // COMMENT: For Go migration, this will be a server request to update 
    // the assignment in the database via a Gin endpoint
    setRoles(prevRoles => {
      const newRoles = [...prevRoles];
      
      // Find the source and target roles
      const sourceRoleIndex = newRoles.findIndex(role => role.id === sourceRoleId);
      const targetRoleIndex = newRoles.findIndex(role => role.id === targetRoleId);
      
      if (sourceRoleIndex === -1 || targetRoleIndex === -1) return prevRoles;
      
      // Find the person in the source role
      const personIndex = newRoles[sourceRoleIndex].people.findIndex(
        person => person.id === personId
      );
      
      if (personIndex === -1) return prevRoles;
      
      // Remove the person from the source role
      const [person] = newRoles[sourceRoleIndex].people.splice(personIndex, 1);
      
      // Add the person to the target role
      newRoles[targetRoleIndex].people.push({
        ...person,
        assigned: newRoles[targetRoleIndex].title || 'Unassigned'
      });
      
      return newRoles;
    });
  };
  
  return (
    <DragContext.Provider value={{ roles, movePerson }}>
      {children}
    </DragContext.Provider>
  );
};

export const useDragContext = () => {
  const context = useContext(DragContext);
  if (context === undefined) {
    throw new Error('useDragContext must be used within a DragProvider');
  }
  return context;
};
