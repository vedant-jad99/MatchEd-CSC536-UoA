
export interface Role {
  id: string;
  title: string;
  description: string;
  priority: 'Low' | 'Medium' | 'High';
  status: 'To Do' | 'In Progress' | 'Completed';
  dueDate: string;
  progress: number;
  assignee?: {
    id: string;
    name: string;
    avatar?: string;
  };
}
