
import React from 'react';
import { MainLayout } from '@/components/layout/MainLayout';

const NewPeople = () => {
  return (
    <MainLayout>
      <div className="space-y-6">
        <div className="flex flex-col md:flex-row md:items-center md:justify-between gap-4">
          <div>
            <h1 className="text-2xl font-bold tracking-tight">New People</h1>
            <p className="text-muted-foreground">Add new individuals to the system</p>
          </div>
        </div>
        
        <div className="flex items-center justify-center h-64 border rounded-lg bg-muted/20">
          <p className="text-muted-foreground">New people management coming soon</p>
        </div>
      </div>
    </MainLayout>
  );
};

export default NewPeople;
