
import { Toaster } from "@/components/ui/toaster";
import { Toaster as Sonner } from "@/components/ui/sonner";
import { TooltipProvider } from "@/components/ui/tooltip";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { ThemeProvider } from "@/hooks/use-theme";
import Index from "./pages/Index";
import NotFound from "./pages/NotFound";
import Courses from "./pages/Courses"; 
import Settings from "./pages/Settings";
import People from "./pages/People";
import Calendar from "./pages/Calendar";
import NewPeople from "./pages/NewPeople";
import All from "./pages/All";

// COMMENT: This section will need to be updated when migrating to Go
// The QueryClient is React-specific, and we'll need to implement equivalent
// state management and data fetching in Go using appropriate libraries
const queryClient = new QueryClient();

const App = () => (
  <QueryClientProvider client={queryClient}>
    <ThemeProvider>
      <TooltipProvider>
        <Toaster />
        <Sonner />
        <BrowserRouter>
          {/* COMMENT: When migrating to Go, this client-side routing will be replaced
              with server-side routing. Routes will need to be defined in Go's HTTP router */}
          <Routes>
            <Route path="/" element={<Index />} />
            <Route path="/courses" element={<Courses />} />
            <Route path="/people" element={<People />} />
            <Route path="/all" element={<All />} />
            <Route path="/settings" element={<Settings />} />
            <Route path="/calendar" element={<Calendar />} />
            <Route path="/new-people" element={<NewPeople />} />
            {/* ADD ALL CUSTOM ROUTES ABOVE THE CATCH-ALL "*" ROUTE */}
            <Route path="*" element={<NotFound />} />
          </Routes>
        </BrowserRouter>
      </TooltipProvider>
    </ThemeProvider>
  </QueryClientProvider>
);

export default App;
