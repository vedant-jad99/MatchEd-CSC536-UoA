
import { Button } from "@/components/ui/button";
import { Link } from "react-router-dom";
import MatchingResults from "@/components/MatchingResults";

const Index = () => {
  return (
    <div className="min-h-screen flex flex-col">
      <header className="border-b">
        <div className="container mx-auto py-3 px-4 flex justify-between items-center">
          <h1 className="text-lg font-bold">Course-User Matching System</h1>
          <div>
            <Button asChild variant="outline" size="sm">
              <Link to="/">Home</Link>
            </Button>
          </div>
        </div>
      </header>
      
      <main className="flex-1 container mx-auto py-4 px-4">
        <MatchingResults />
      </main>
      
      <footer className="border-t mt-auto">
        <div className="container mx-auto py-3 px-4 text-center text-xs text-muted-foreground">
          &copy; 2025 Course-User Matching System
        </div>
      </footer>
    </div>
  );
};

export default Index;
