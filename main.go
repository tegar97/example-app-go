package main

import (
	"fmt"
	"log"
	"net/http"
)

// HelloWorldHandler handles requests to the root endpoint
func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test webhook")
}

// HelloHandler handles requests to the /hello endpoint
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World 192922a2sd3222222222222ss222s222")
}

func main() {
	// Register handlers for routes
	http.HandleFunc("/", HelloWorldHandler)
	http.HandleFunc("/hello", HelloHandler)

	// Start the server on port 8080
	fmt.Println("Server starting on port 3001...")
	if err := http.ListenAndServe(":3001", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
