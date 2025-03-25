
import { Role } from '@/types/role';

export const priorityColorMap = {
  'Low': 'bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-300',
  'Medium': 'bg-yellow-100 text-yellow-800 dark:bg-yellow-900 dark:text-yellow-300',
  'High': 'bg-red-100 text-red-800 dark:bg-red-900 dark:text-red-300'
};

export const statusColorMap = {
  'To Do': 'bg-slate-100 text-slate-800 dark:bg-slate-900 dark:text-slate-300',
  'In Progress': 'bg-purple-100 text-purple-800 dark:bg-purple-900 dark:text-purple-300',
  'Completed': 'bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-300'
};

export const filterRolesByStatus = (roles: Role[], statusFilter: 'All' | Role['status']) => {
  if (statusFilter === 'All') {
    return roles;
  }
  return roles.filter(role => role.status === statusFilter);
};
