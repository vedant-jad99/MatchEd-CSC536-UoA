
import React from 'react';
import { MainLayout } from '@/components/layout/MainLayout';
import { Button } from '@/components/ui/button';
import { ChevronRight } from 'lucide-react';
import { useNavigate } from 'react-router-dom';

const Index = () => {
  const navigate = useNavigate();
  
  return (
    <MainLayout>
      <div className="flex flex-col items-center justify-center min-h-screen w-full text-center px-4 -mt-16">
        <div 
          className="relative w-full max-w-4xl mx-auto p-8 rounded-2xl overflow-hidden backdrop-blur-sm border shadow-lg"
          style={{
            background: "linear-gradient(to right bottom, rgba(255, 255, 255, 0.1), rgba(255, 255, 255, 0.05))"
          }}
        >
          {/* Background decoration elements - made larger */}
          <div className="absolute top-0 left-0 w-[40rem] h-[40rem] bg-purple-500/10 rounded-full -translate-x-1/2 -translate-y-1/2 blur-3xl"></div>
          <div className="absolute bottom-0 right-0 w-[50rem] h-[50rem] bg-blue-500/15 rounded-full translate-x-1/3 translate-y-1/3 blur-3xl"></div>
          <div className="absolute top-1/3 right-0 w-[35rem] h-[35rem] bg-teal-500/10 rounded-full translate-x-1/4 -translate-y-1/4 blur-3xl"></div>
          
          {/* Geometric shapes for Web3 aesthetic */}
          <div className="absolute top-12 right-12 w-16 h-16 border border-primary/20 rounded-lg rotate-12 animate-float"></div>
          <div className="absolute bottom-20 left-16 w-12 h-12 border border-primary/30 rounded-full animate-pulse-subtle"></div>
          <div className="absolute top-1/3 left-1/4 w-8 h-8 border border-primary/40 rounded-full animate-float"></div>
          
          {/* Content */}
          <div className="relative z-10">
            <h1 className="text-4xl md:text-6xl font-bold mb-6 bg-clip-text text-transparent bg-gradient-to-r from-primary to-blue-500 font-poppins">
              matchEd
            </h1>
            
            <p className="text-xl md:text-2xl text-foreground/90 max-w-2xl mx-auto leading-relaxed mb-8 font-poppins">
              A decision-support tool designed to efficiently assign individuals to roles
            </p>
            
            <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mb-10">
              <div className="p-4 rounded-lg bg-card/60 backdrop-blur-sm border shadow-sm">
                <h3 className="font-medium text-lg mb-2 font-poppins">Smart Matching</h3>
                <p className="text-muted-foreground">Intelligently match people to roles based on skills and requirements</p>
              </div>
              
              <div className="p-4 rounded-lg bg-card/60 backdrop-blur-sm border shadow-sm">
                <h3 className="font-medium text-lg mb-2 font-poppins">Drag & Drop</h3>
                <p className="text-muted-foreground">Intuitive interface for manual adjustments and assignments</p>
              </div>
              
              <div className="p-4 rounded-lg bg-card/60 backdrop-blur-sm border shadow-sm">
                <h3 className="font-medium text-lg mb-2 font-poppins">Optimize Teams</h3>
                <p className="text-muted-foreground">Create balanced and effective teams with the right skill mix</p>
              </div>
            </div>
            
            <Button
              size="lg"
              onClick={() => navigate('/courses')}
              className="px-8 py-6 text-lg gap-3 font-poppins"
            >
              Get Started
              <ChevronRight className="h-5 w-5" />
            </Button>
          </div>
        </div>
      </div>
    </MainLayout>
  );
};

export default Index;
