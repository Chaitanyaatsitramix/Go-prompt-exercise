package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// HealthResponse represents the JSON response for health endpoint
type HealthResponse struct {
	Status string `json:"status"`
}

// ErrorResponse represents the JSON response for errors
type ErrorResponse struct {
	Error string `json:"error"`
}

func main() {
	// Check if we have command line arguments for file reading demonstration
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "test-file-reader":
			testReadFileContent()
			return
		case "create-test-files":
			createTestFiles()
			return
		case "cleanup-test-files":
			cleanupTestFiles()
			return
		case "demo-file-reader":
			// Simple demo using the existing functions
			fmt.Println("=== File Reading Demo ===")
			fmt.Println("This demonstrates the ReadFileContent function.")
			fmt.Println("For comprehensive testing, use: go run . test-file-reader")
			fmt.Println()

			// Create sample file if it doesn't exist
			if _, err := os.Stat("sample.txt"); os.IsNotExist(err) {
				fmt.Println("Creating sample.txt for demonstration...")
				createTestFiles()
				fmt.Println()
			}

			// Demo reading the sample file
			fmt.Println("Reading sample.txt:")
			content, err := ReadFileContent("sample.txt")
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Printf("✅ Successfully read %d characters\n", len(content))
				fmt.Printf("Content preview: %q\n", truncateString(content, 150))
			}
			return
		}
	}

	// Default behavior: start the HTTP server
	startHTTPServer()
}

// startHTTPServer starts the HTTP server and contains the main server logic
func startHTTPServer() {
	// Create a new ServeMux for better routing control
	mux := http.NewServeMux()

	// Register routes with the mux
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/health", healthHandler)
	// Add new endpoint for file reading demonstration
	mux.HandleFunc("/read-file", fileReadHandler)

	// Wrap the mux with logging middleware
	loggedMux := loggingMiddleware(mux)

	// Start the server
	port := ":8080"
	fmt.Printf("Starting HTTP server on port %s...\n", port)
	fmt.Println("Available endpoints:")
	fmt.Println("  GET /           - Welcome message")
	fmt.Println("  GET /hello      - Hello message (supports ?name=YourName)")
	fmt.Println("  GET /health     - Health check")
	fmt.Println("  GET /read-file  - Demonstrate file reading (supports ?file=filename)")
	fmt.Println()
	fmt.Println("File reading options:")
	fmt.Println("  go run . demo-file-reader     # Demo file reading functions")
	fmt.Println("  go run . test-file-reader     # Run comprehensive tests")
	fmt.Println("  go run . create-test-files    # Create test files")
	fmt.Println("  go run . cleanup-test-files   # Remove test files")

	// Use defer for cleanup (though log.Fatal will exit before defer executes)
	defer func() {
		fmt.Println("Server shutting down...")
	}()

	// Start server with proper error handling
	if err := http.ListenAndServe(port, loggedMux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// homeHandler handles GET requests to the root path "/"
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Method validation - only allow GET requests
	if r.Method != http.MethodGet {
		handleMethodNotAllowed(w, r)
		return
	}

	// Route validation - ensure exact path match for root
	if r.URL.Path != "/" {
		handleNotFound(w, r)
		return
	}

	// Set appropriate headers
	setCommonHeaders(w)
	w.Header().Set("Content-Type", "text/plain")

	// Send response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Welcome to the Go HTTP Server")
}

// helloHandler handles GET requests to "/hello" with optional name parameter
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Method validation - only allow GET requests
	if r.Method != http.MethodGet {
		handleMethodNotAllowed(w, r)
		return
	}

	// Extract name parameter from query string, default to "Guest"
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}

	// Set appropriate headers
	setCommonHeaders(w)
	w.Header().Set("Content-Type", "text/plain")

	// Send personalized greeting
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, %s!", name)
}

// healthHandler handles GET requests to "/health" endpoint
func healthHandler(w http.ResponseWriter, r *http.Request) {
	// Method validation - only allow GET requests
	if r.Method != http.MethodGet {
		handleMethodNotAllowed(w, r)
		return
	}

	// Set appropriate headers
	setCommonHeaders(w)
	w.Header().Set("Content-Type", "application/json")

	// Create health response
	response := HealthResponse{
		Status: "ok",
	}

	// Send JSON response
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding health response: %v", err)
		handleInternalServerError(w, r)
		return
	}
}

// fileReadHandler handles requests to demonstrate file reading via HTTP
func fileReadHandler(w http.ResponseWriter, r *http.Request) {
	// Method validation - only allow GET requests
	if r.Method != http.MethodGet {
		handleMethodNotAllowed(w, r)
		return
	}

	// Get filename from query parameter
	filename := r.URL.Query().Get("file")
	if filename == "" {
		filename = "sample.txt" // Default file
	}

	// Set appropriate headers
	setCommonHeaders(w)
	w.Header().Set("Content-Type", "text/plain")

	// Try to read the file
	content, err := ReadFileContent(filename)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error reading file '%s': %v\n", filename, err)
		return
	}

	// Send file content
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Content of '%s' (%d characters):\n\n%s", filename, len(content), content)
}

// handleNotFound handles 404 errors for unknown routes
func handleNotFound(w http.ResponseWriter, r *http.Request) {
	setCommonHeaders(w)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Route not found.")
}

// handleMethodNotAllowed handles 405 errors for invalid HTTP methods
func handleMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	setCommonHeaders(w)
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Allow", "GET") // Specify allowed methods
	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprint(w, "Method Not Allowed.")
}

// handleInternalServerError handles 500 errors
func handleInternalServerError(w http.ResponseWriter, r *http.Request) {
	setCommonHeaders(w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	response := ErrorResponse{
		Error: "Internal Server Error",
	}
	json.NewEncoder(w).Encode(response)
}

// setCommonHeaders sets common security and response headers
func setCommonHeaders(w http.ResponseWriter) {
	// Basic security headers - important for production applications

	// Prevents the browser from interpreting files as a different MIME type
	w.Header().Set("X-Content-Type-Options", "nosniff")

	// Enables the Cross-site scripting (XSS) filter built into modern browsers
	w.Header().Set("X-XSS-Protection", "1; mode=block")

	// Prevents the page from being embedded in frames (clickjacking protection)
	w.Header().Set("X-Frame-Options", "DENY")

	// Additional security headers to consider for production:
	// w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
	// w.Header().Set("Content-Security-Policy", "default-src 'self'")
	// w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
}

// loggingMiddleware provides basic request logging
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Record start time for request duration calculation
		start := time.Now()

		// Create a custom ResponseWriter to capture status code
		wrapped := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		// Process the request
		next.ServeHTTP(wrapped, r)

		// Log request details
		duration := time.Since(start)
		log.Printf(
			"[%s] %s %s - Status: %d - Duration: %v - User-Agent: %s",
			time.Now().Format("2006-01-02 15:04:05"),
			r.Method,
			r.URL.Path,
			wrapped.statusCode,
			duration,
			r.UserAgent(),
		)
	})
}

// responseWriter wraps http.ResponseWriter to capture status codes
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code before writing
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// truncateString truncates a string to maxLength and adds "..." if needed
func truncateString(s string, maxLength int) string {
	if len(s) <= maxLength {
		return s
	}
	return s[:maxLength] + "..."
}
