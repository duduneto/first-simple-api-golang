package main

import (
	"fmt"
	"net/http"

	"library.com/api/routes"
)

func main() {
	// Create a new HTTP server
	server := http.NewServeMux()

	// Set Routes
	routes.HandleRoutes(server)

	// Start the server on port 8080
	fmt.Println("Server listening on http://localhost:3000")
	http.ListenAndServe(":3000", server)
}
