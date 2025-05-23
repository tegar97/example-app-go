package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// HelloWorldHandler handles requests to the root endpoint
func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(w, "Update 223 32324 Time : %s", timestamp)
}

// HelloHandler handles requests to the /hello endpoint
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World 42222")
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
