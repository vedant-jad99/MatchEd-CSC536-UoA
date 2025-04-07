
import React from 'react';
import { MainLayout } from '@/components/layout/MainLayout';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { useState, useEffect } from 'react';
import { Preferences } from "../types/all";
import axios from 'axios';

import { 
  Pagination, 
  PaginationContent, 
  PaginationItem, 
  PaginationNext, 
  PaginationPrevious 
} from '@/components/ui/pagination';
import { Search, ArrowUpAZ, ArrowDownAZ, Upload, Download } from 'lucide-react';
import { types } from 'util';

const All = () => {
  const [preferences, setPreferences] = useState<Preferences[]>([]); // State to hold preferences
  const [loading, setLoading] = useState(false); 
  const [error, setError] = useState<string | null>(null); 
  const [currentPage, setCurrentPage] = useState(1);

  // COMMENT: For Go/Gin migration, this would be replaced with server-side data fetching
  // The following functions would trigger API calls to the Gin backend

  const handleSearch = (query: string) => {
    // COMMENT: This should trigger a SQL query on the Go backend
    // Example: SELECT * FROM items WHERE name LIKE '%query%' OR description LIKE '%query%'
    console.log('Search query:', query);
  };

  const handleSort = (column: string, direction: 'asc' | 'desc') => {
    // COMMENT: This should trigger a SQL ORDER BY clause on the Go backend
    // Example: SELECT * FROM items ORDER BY ${column} ${direction}
    console.log(`Sort by ${column} ${direction}`);
  };

  const handlePagination = (page: number) => {
    
    //server call (api/people/get)


    // COMMENT: This should trigger SQL pagination on the Go backend
    // Example: SELECT * FROM items LIMIT 10 OFFSET (page - 1) * 10
    setCurrentPage(page); // Update current page state
    console.log('Go to page:', page);
  };

  const handleImport = (event: React.ChangeEvent<HTMLInputElement>) => {
    // COMMENT: This should trigger a file upload to the Go backend
    // The Go server would parse the CSV/Excel and insert data into the database
    console.log('Import file:', event.target.files?.[0]?.name);
  };

  const handleExport = () => {
    // COMMENT: This should trigger a SQL query on the Go backend to export data
    // The Go server would format the data as CSV/Excel and return it as a download
    console.log('Export data');
  };

  const ShowPreferencesButton = async () => {
      setLoading(true);
      setError(null);
  
      try {
        // API request to backend
        // useEffect(() => {
        //     axios.get<Preference[]>("http://localhost:3000/api/all")
        //     .then((res) => {
        //       setPreferences(res.data);
        //     })
        //     .catch((err) => console.error(err));
        // }, []);


        // return (
        //   <div>
        //     <h2>API Data:</h2>
        //     <pre>{JSON.stringify(data, null, 2)}</pre>
        //   </div>
        // );
        const response = await axios.get('http://localhost:3000/api/all');
        console.log("response:", response.data)
        setPreferences(response.data); // Storing preferences
      } catch (err) {
        setError('Failed to load preferences');
      } finally {
        setLoading(false);
      }
  };
  
  
  return (
    <MainLayout>
      <div className="container mx-auto py-6">
        <h1 className="text-3xl font-bold mb-6">All Data</h1>
        
        <div className="mb-8 p-6 border rounded-lg bg-card shadow-sm">
          <p className="text-center text-muted-foreground mb-4">
            This page will display data from the database.
            <br />
            The actual data display will be implemented with the Go backend.
          </p>
          
          {/* Database Operation Controls */}
          <div className="space-y-6">
            {/* Search */}
            <div className="flex items-center gap-2">
              <Input 
                placeholder="Search..." 
                className="max-w-sm"
                onChange={(e) => handleSearch(e.target.value)}
              />
              <Button variant="outline" size="icon">
                <Search className="h-4 w-4" />
              </Button>
            </div>
            
            {/* Sort */}
            <div className="flex items-center gap-2">
              <span className="text-sm font-medium">Sort By:</span>
              <Button 
                variant="outline" 
                size="sm"
                onClick={() => handleSort('name', 'asc')}
                className="gap-1"
              >
                Name <ArrowUpAZ className="h-3 w-3" />
              </Button>
              <Button 
                variant="outline" 
                size="sm"
                onClick={() => handleSort('name', 'desc')}
                className="gap-1"
              >
                Name <ArrowDownAZ className="h-3 w-3" />
              </Button>
              <Button 
                variant="outline" 
                size="sm"
                onClick={() => handleSort('date', 'desc')}
                className="gap-1"
              >
                Date <ArrowDownAZ className="h-3 w-3" />
              </Button>
            </div>
            
            {/* Import/Export */}
            <div className="flex items-center gap-2">
              <Button 
                variant="outline" 
                className="gap-2"
                onClick={() => document.getElementById('file-upload')?.click()}
              >
                <Upload className="h-4 w-4" /> Import
              </Button>
              <input 
                id="file-upload" 
                type="file" 
                accept=".csv,.xlsx,.xls" 
                className="hidden" 
                onChange={handleImport}
              />
              <Button 
                variant="outline" 
                className="gap-2"
                onClick={handleExport}
              >
                <Download className="h-4 w-4" /> Export
              </Button>
            </div>

            {/* {Show Preferences Button} */}
            <Button variant="outline" className="gap-2" onClick={ShowPreferencesButton}>
              <Upload className="h-4 w-4" /> Show Preferences
            </Button>

            {error && <p>{error}</p>}

            {/* {Rendering preferences table} */}
            {!loading && !error && preferences.length > 0 && (
              <table className="min-w-full mt-4">
                <thead>
                  <tr>
                    <th className="border px-4 py-2">ID</th>
                    <th className="border px-4 py-2">Semester</th>
                    <th className="border px-4 py-2">UserID</th>
                    <th className="border px-4 py-2">Name</th>
                    <th className="border px-4 py-2">CourseSemID</th>
                    <th className="border px-4 py-2">Preference Level</th>
                  </tr>
                </thead>
                <tbody>
                  {preferences.map((preference) => (
                    <tr key={preference.id}>
                      <td className="border px-4 py-2">{preference.id}</td>
                      <td className="border px-4 py-2">{preference.semester}</td>
                      <td className="border px-4 py-2">{preference.user_id}</td>
                      <td className="border px-4 py-2">{preference.course_sem_id}</td>
                      <td className="border px-4 py-2">{preference.preference_level}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            )} 

            {/* Pagination */}
            <Pagination>
              <PaginationContent>
                <PaginationItem>
                  <PaginationPrevious onClick={() => handlePagination(currentPage -1)} />
                </PaginationItem>
                <PaginationItem>
                  <PaginationNext onClick={() => handlePagination(currentPage + 1)} />
                </PaginationItem>
              </PaginationContent>
            </Pagination>
          </div>
        </div>
      </div>
    </MainLayout>
  );
};

export default All;
