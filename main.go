package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"sync/atomic"
	"time"
)

// Global variables for metrics
var (
	startTime    = time.Now()
	requestCount int64
	helloCount   int64
)

// HelloWorldHandler handles requests to the root endpoint
func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	atomic.AddInt64(&requestCount, 1)
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	envValue := os.Getenv("APP_ENV")
	fmt.Fprintf(w, "Environment Value: %s, Timestamp: %s", envValue, timestamp)

	// evv debug_mode value gets
	fmt.Fprintf(w, "Update baru baru bangets22")
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
	atomic.AddInt64(&helloCount, 1)
	fmt.Fprintf(w, "Hello World 42222")
}

// MetricsHandler handles requests to the /metrics endpoint for Grafana/Prometheus
func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; version=0.0.4; charset=utf-8")

	// Calculate uptime
	uptime := time.Since(startTime).Seconds()

	// Get memory stats
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	// Get number of goroutines
	numGoroutines := runtime.NumGoroutine()

	// Write metrics in Prometheus format
	fmt.Fprintf(w, "# HELP app_uptime_seconds Application uptime in seconds\n")
	fmt.Fprintf(w, "# TYPE app_uptime_seconds counter\n")
	fmt.Fprintf(w, "app_uptime_seconds %s\n", strconv.FormatFloat(uptime, 'f', 2, 64))

	fmt.Fprintf(w, "\n# HELP app_requests_total Total number of requests to root endpoint\n")
	fmt.Fprintf(w, "# TYPE app_requests_total counter\n")
	fmt.Fprintf(w, "app_requests_total %d\n", atomic.LoadInt64(&requestCount))

	fmt.Fprintf(w, "\n# HELP app_hello_requests_total Total number of requests to hello endpoint\n")
	fmt.Fprintf(w, "# TYPE app_hello_requests_total counter\n")
	fmt.Fprintf(w, "app_hello_requests_total %d\n", atomic.LoadInt64(&helloCount))

	fmt.Fprintf(w, "\n# HELP app_memory_alloc_bytes Current allocated memory in bytes\n")
	fmt.Fprintf(w, "# TYPE app_memory_alloc_bytes gauge\n")
	fmt.Fprintf(w, "app_memory_alloc_bytes %d\n", m.Alloc)

	fmt.Fprintf(w, "\n# HELP app_memory_sys_bytes Total memory obtained from system in bytes\n")
	fmt.Fprintf(w, "# TYPE app_memory_sys_bytes gauge\n")
	fmt.Fprintf(w, "app_memory_sys_bytes %d\n", m.Sys)

	fmt.Fprintf(w, "\n# HELP app_goroutines_count Current number of goroutines\n")
	fmt.Fprintf(w, "# TYPE app_goroutines_count gauge\n")
	fmt.Fprintf(w, "app_goroutines_count %d\n", numGoroutines)

	fmt.Fprintf(w, "\n# HELP app_gc_runs_total Total number of GC runs\n")
	fmt.Fprintf(w, "# TYPE app_gc_runs_total counter\n")
	fmt.Fprintf(w, "app_gc_runs_total %d\n", m.NumGC)
}

func main() {
	// Register handlers for routes
	http.HandleFunc("/", HelloWorldHandler)
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/metrics", MetricsHandler)

	// Start the server on port 8080
	fmt.Println("Server startsssing on port 3001...")
	if err := http.ListenAndServe(":3001", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
