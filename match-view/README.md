# Match View Display System

This application provides an interface for viewing and filtering matching results between courses and users.

## Project Setup for Go Server Integration

This guide will help you set up the React application to be served by your Go server from a "frontend" folder.

### Prerequisites

- Node.js (v16 or newer)
- npm (comes with Node.js)
- Your Go server configured to serve static files from a "frontend" folder

### Step 1: Clone and Install Dependencies

```sh
# Clone the repository
git clone <your-repository-url>

# Navigate to project directory
cd match-view-display

# Install dependencies
npm install
```

### Step 2: Configure the Build Output

The application needs to be configured to build to your Go server's "frontend" directory. You have two options:

#### Option 1: Build directly to the frontend directory

Edit the `vite.config.ts` file to specify the build output directory:

```typescript
import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";
import path from "path";

// https://vitejs.dev/config/
export default defineConfig(({ mode }) => ({
  server: {
    host: "::",
    port: 8080,
  },
  plugins: [
    react(),
  ],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
    },
  },
  build: {
    outDir: '../path-to-your-go-project/frontend', // Adjust this to your Go project's frontend folder
  },
}));
```

#### Option 2: Build and copy (recommended)

Keep the default build configuration and copy the files after building:

```sh
# Build the application (creates a 'dist' folder)
npm run build

# Copy the build output to your Go server's frontend folder
cp -r dist/* /path-to-your-go-project/frontend/
```

### Step 3: Build for Production

```sh
# Run the production build
npm run build
```

This command creates optimized production files in the `dist` directory (or your configured output directory).

### Step 4: Move Files to Go Server (if using Option 2)

```sh
# Copy all built files to your Go server's frontend folder
cp -r dist/* /path-to-your-go-project/frontend/
```

### Step 5: Update Go Server Configuration (if needed)

Ensure your Go server is configured to:

1. Serve static files from the "frontend" directory
2. Redirect all non-API routes to the index.html file (for React Router to work)

Example Go code for serving static files:

```go
// Import required packages
import (
    "net/http"
)

func main() {
    // Serve static files from the "frontend" directory
    fs := http.FileServer(http.Dir("./frontend"))
    http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Check if the file exists
        path := "./frontend" + r.URL.Path
        _, err := os.Stat(path)
        
        // If file doesn't exist, serve index.html for client-side routing
        if os.IsNotExist(err) || r.URL.Path == "/" {
            http.ServeFile(w, r, "./frontend/index.html")
            return
        }
        
        // Otherwise serve the requested file
        fs.ServeHTTP(w, r)
    }))
    
    // Start the server
    http.ListenAndServe(":8080", nil)
}
```

### Development Workflow

For local development (without Go server):

```sh
# Start the development server
npm run dev
```

This will run the application on [http://localhost:8080](http://localhost:8080).

### Additional Notes

- Make sure that your Go server properly handles all routing, especially if you're using client-side routing in React.
- If you encounter CORS issues, ensure your Go server sets appropriate CORS headers for development.
- The built React app is entirely client-side and requires no server-side rendering.

## Application Features

- View matching results between courses and users
- Filter results by status and confidence level
- Search through matches
- Pagination for large result sets

## Troubleshooting

If you encounter issues with routes in your Go-served app, ensure that your server is configured to fall back to `index.html` for all non-file routes, as shown in the example Go code above.
