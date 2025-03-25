
import React from 'react';
import { Link, useLocation } from 'react-router-dom';
import { Home, Users, Calendar, BookOpen, List, Plus, Settings, MoreVertical, Database } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { cn } from '@/lib/utils';

interface SidebarProps {
  open: boolean;
  setOpen: (open: boolean) => void;
}

// COMMENT: For Go migration, this would be defined server-side
// and passed as data to the template
const menuItems = [
  { name: 'Home', path: '/', icon: Home },
  { name: 'Courses', path: '/courses', icon: BookOpen },
  { name: 'People', path: '/people', icon: Users },
  { name: 'All', path: '/all', icon: Database },
  { name: 'Calendar', path: '/calendar', icon: Calendar },
  { name: 'New People', path: '/new-people', icon: Plus },
  { name: 'Settings', path: '/settings', icon: Settings },
];

export const Sidebar: React.FC<SidebarProps> = ({ open, setOpen }) => {
  // COMMENT: For Go migration, the current route would be determined server-side
  const location = useLocation();
  
  return (
    <>
      {/* Overlay for mobile */}
      {open && (
        <div 
          className="fixed inset-0 bg-black/20 backdrop-blur-xs z-20 md:hidden"
          onClick={() => setOpen(false)}
          aria-hidden="true"
        />
      )}
      
      <aside 
        className={cn(
          "fixed md:relative z-30 h-full flex flex-col border-r bg-background transition-all duration-300 ease-in-out",
          open ? "w-64" : "w-0 md:w-16 overflow-hidden"
        )}
      >
        <div className="flex items-center justify-between h-16 px-4 border-b">
          {open && <h1 className="text-lg font-medium">Menu</h1>}
          <Button 
            variant="ghost" 
            size="sm" 
            className={cn("ml-auto", !open && "w-full justify-center")} 
            onClick={() => setOpen(!open)}
            aria-label={open ? "Close sidebar" : "Open sidebar"}
          >
            <MoreVertical className="h-5 w-5" />
          </Button>
        </div>
        
        <nav className="flex-1 overflow-y-auto py-4">
          <ul className="space-y-1 px-2">
            {menuItems.map((item) => (
              <li key={item.name}>
                <Link 
                  to={item.path} 
                  className={cn(
                    "flex items-center px-3 py-2 rounded-lg transition-colors",
                    "hover:bg-muted hover:text-foreground focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring",
                    location.pathname === item.path ? "bg-primary/10 text-primary font-medium" : "text-muted-foreground"
                  )}
                >
                  <item.icon className={cn("h-5 w-5", !open && "mx-auto")} />
                  {open && <span className="ml-3">{item.name}</span>}
                </Link>
              </li>
            ))}
          </ul>
        </nav>
      </aside>
    </>
  );
};
