
import { Role } from '@/types/role';

// COMMENT: For Go migration, this data would come from a database query
export const roles: Role[] = [
  {
    id: 'role-1',
    title: 'Assign developers to projects',
    description: 'Review available developers and assign them to upcoming projects based on skills and availability.',
    dueDate: '2023-06-15',
    priority: 'High',
    status: 'In Progress',
    assignee: {
      id: 'person-1',
      name: 'John Doe',
    },
    progress: 60,
  },
  {
    id: 'role-2',
    title: 'Update role descriptions',
    description: 'Review and update all role descriptions to match current project requirements.',
    dueDate: '2023-06-20',
    priority: 'Medium',
    status: 'To Do',
    assignee: {
      id: 'person-2',
      name: 'Alice Johnson',
    },
    progress: 0,
  },
  {
    id: 'role-3',
    title: 'Conduct skills assessment',
    description: 'Assess team members\' skills to better match them with appropriate roles.',
    dueDate: '2023-06-10',
    priority: 'Medium',
    status: 'Completed',
    assignee: {
      id: 'person-3',
      name: 'Bob Smith',
    },
    progress: 100,
  },
  {
    id: 'role-4',
    title: 'Review role assignments',
    description: 'Quarterly review of all role assignments to ensure optimal team performance.',
    dueDate: '2023-06-30',
    priority: 'Low',
    status: 'To Do',
    progress: 0,
  },
  {
    id: 'role-5',
    title: 'Update project staffing plan',
    description: 'Revise the staffing plan for upcoming projects based on skill requirements.',
    dueDate: '2023-06-25',
    priority: 'High',
    status: 'In Progress',
    assignee: {
      id: 'person-4',
      name: 'Diana Lee',
    },
    progress: 35,
  },
];
