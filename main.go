package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// HelloWorldHandler handles requests to the root endpoint
func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	envValue := os.Getenv("APP_ENV")
	fmt.Fprintf(w, "Environment Value: %s, Timestamp: %s", envValue, timestamp)

	// evv debug_mode value gets
	fmt.Fprintf(w, "baru banget ini jam 11:04")
	debugMode := os.Getenv("DEBUG_MODE")
	fmt.Fprintf(w, "Debug Mode: %s", debugMode)

	// get list env
	fmt.Fprintln(w, "\n\nEnvironment List:")
	for _, env := range os.Environ() {
		fmt.Fprintln(w, env)
	}

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
