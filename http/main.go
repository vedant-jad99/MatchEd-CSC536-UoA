package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := 8000
	fmt.Printf("Starting server on port %d...\n", port)

	// Serve static files from the current directory
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
