
import React, { useState } from 'react';
import { MainLayout } from '@/components/layout/MainLayout';
import { Button } from '@/components/ui/button';
import { Calendar as CalendarComponent } from '@/components/ui/calendar';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { ChevronLeft, ChevronRight, Plus } from 'lucide-react';

interface Role {
  id: string;
  title: string;
  startTime: string;
  endTime: string;
  color: string;
  day: number; // 0-6 (Sunday-Saturday)
}

// Sample roles data
const roles: Role[] = [
  {
    id: 'role-1',
    title: 'Developer',
    startTime: '09:00',
    endTime: '12:00',
    color: 'border-blue-500',
    day: 1, // Monday
  },
  {
    id: 'role-2',
    title: 'Designer',
    startTime: '13:00',
    endTime: '17:00',
    color: 'border-purple-500',
    day: 2, // Tuesday
  },
  {
    id: 'role-3',
    title: 'Manager',
    startTime: '10:00',
    endTime: '15:00',
    color: 'border-green-500',
    day: 3, // Wednesday
  },
  {
    id: 'role-4',
    title: 'Tester',
    startTime: '14:00',
    endTime: '16:00',
    color: 'border-yellow-500',
    day: 4, // Thursday
  },
];

const CalendarSlot: React.FC<{ role?: Role }> = ({ role }) => {
  return role ? (
    <div 
      className={`p-2 border-l-4 ${role.color} bg-opacity-10 rounded`}
      style={{ backgroundColor: role.color.replace('border-', 'bg-').replace('-500', '-100') }}
    >
      <p className="font-medium text-sm">{role.title}</p>
      <p className="text-xs text-muted-foreground">{role.startTime} - {role.endTime}</p>
    </div>
  ) : (
    <div className="p-2 border border-dashed border-muted rounded h-full min-h-[60px] flex items-center justify-center">
      <span className="text-xs text-muted-foreground">Available</span>
    </div>
  );
};

const WeekView: React.FC<{ roles: Role[], currentDate: Date }> = ({ roles, currentDate }) => {
  const weekDays = ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday'];
  const timeSlots = ['9:00', '10:00', '11:00', '12:00', '13:00', '14:00', '15:00', '16:00', '17:00'];
  
  // Get the start of the week (Sunday)
  const startOfWeek = new Date(currentDate);
  startOfWeek.setDate(currentDate.getDate() - currentDate.getDay());
  
  return (
    <div className="overflow-x-auto">
      <div className="min-w-[800px]">
        <div className="grid grid-cols-8 gap-2 mb-2">
          <div className="font-medium text-center p-2"></div>
          {weekDays.map((day, index) => {
            const date = new Date(startOfWeek);
            date.setDate(startOfWeek.getDate() + index);
            
            return (
              <div key={index} className="font-medium text-center p-2 bg-muted/20 rounded">
                <div>{day}</div>
                <div className="text-sm text-muted-foreground">
                  {date.toLocaleDateString(undefined, { month: 'short', day: 'numeric' })}
                </div>
              </div>
            );
          })}
        </div>
        
        {timeSlots.map((time, timeIndex) => (
          <div key={timeIndex} className="grid grid-cols-8 gap-2 mb-2">
            <div className="font-medium text-right p-2 text-sm">{time}</div>
            {weekDays.map((_, dayIndex) => {
              const matchingRole = roles.find(role => 
                role.day === dayIndex && 
                role.startTime <= time && 
                role.endTime > time
              );
              
              return (
                <div key={dayIndex} className="min-h-[60px]">
                  <CalendarSlot role={matchingRole} />
                </div>
              );
            })}
          </div>
        ))}
      </div>
    </div>
  );
};

const Calendar = () => {
  const [date, setDate] = useState<Date>(new Date());
  
  const handlePrevWeek = () => {
    const newDate = new Date(date);
    newDate.setDate(newDate.getDate() - 7);
    setDate(newDate);
  };
  
  const handleNextWeek = () => {
    const newDate = new Date(date);
    newDate.setDate(newDate.getDate() + 7);
    setDate(newDate);
  };
  
  return (
    <MainLayout>
      <div className="space-y-6">
        <div className="flex flex-col md:flex-row md:items-center md:justify-between gap-4">
          <div>
            <h1 className="text-2xl font-bold tracking-tight">Calendar</h1>
            <p className="text-muted-foreground">View role allocations and conflicts</p>
          </div>
          
          <div className="flex items-center space-x-2">
            <Button>
              <Plus className="mr-2 h-4 w-4" />
              New Role
            </Button>
          </div>
        </div>
        
        <Card>
          <CardHeader className="px-6 py-3 border-b">
            <div className="flex items-center justify-between">
              <CardTitle className="text-lg">
                Week of {date.toLocaleDateString(undefined, { month: 'long', day: 'numeric', year: 'numeric' })}
              </CardTitle>
              
              <div className="flex items-center space-x-2">
                <Button variant="outline" size="icon" onClick={handlePrevWeek}>
                  <ChevronLeft className="h-4 w-4" />
                </Button>
                <Button variant="outline" size="icon" onClick={() => setDate(new Date())}>
                  Today
                </Button>
                <Button variant="outline" size="icon" onClick={handleNextWeek}>
                  <ChevronRight className="h-4 w-4" />
                </Button>
              </div>
            </div>
          </CardHeader>
          <CardContent className="p-4">
            <WeekView roles={roles} currentDate={date} />
          </CardContent>
        </Card>
      </div>
    </MainLayout>
  );
};

export default Calendar;
