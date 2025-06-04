# Day-2 HTTP Server with File Reading Capabilities

## Project Overview

This is a comprehensive Go HTTP server implementation with advanced file reading capabilities. The project demonstrates proper routing, error handling, middleware implementation, security considerations, and robust file I/O operations following Go best practices.

## Functional Requirements Implementation

### ✅ HTTP Routes
- **GET /** - Returns "Welcome to the Go HTTP Server" (text/plain)
- **GET /hello?name=YourName** - Returns "Hello, YourName!" (text/plain), defaults to "Guest"
- **GET /health** - Returns 200 OK with JSON `{"status": "ok"}`
- **GET /read-file?file=filename** - Demonstrates file reading via HTTP endpoint

### ✅ File Reading Functionality
- **ReadFileContent()** - Main file reading function with comprehensive error handling
- **ReadFileContentSimple()** - Alternative implementation using os.ReadFile
- **ReadFileContentIoutil()** - Educational implementation showing different approaches
- **StreamFileLines()** - Memory-efficient line-by-line file processing

### ✅ Command Line Interface
- **Default behavior** - Starts HTTP server
- **demo-file-reader** - Quick file reading demonstration
- **test-file-reader** - Comprehensive testing suite
- **create-test-files** - Creates test files for demonstration
- **cleanup-test-files** - Removes created test files

## Project Structure

```
Day-2/
├── main.go              # HTTP server + command line interface
├── file_reader.go       # Core file reading functionality
├── test_file_reader.go  # Comprehensive testing functions
├── sample.txt           # Test file with content
├── empty.txt            # Empty test file
├── go.mod              # Go module file
└── README.md           # This documentation
```

## How to Run

### Prerequisites
- Go 1.21 or later installed

### Quick Start

```bash
# Navigate to Day-2 directory
cd Day-2

# Start HTTP server (default)
go run .

# File reading demonstrations
go run . demo-file-reader     # Quick demo
go run . test-file-reader     # Comprehensive tests
go run . create-test-files    # Create test files
go run . cleanup-test-files   # Remove test files
```

## API Endpoints

### 1. Root Endpoint
```http
GET /
```
**Response:** `Welcome to the Go HTTP Server` (text/plain)

### 2. Hello Endpoint
```http
GET /hello?name=YourName
```
**Parameters:**
- `name` (optional): Your name (defaults to "Guest")

**Response:** `Hello, YourName!` (text/plain)

### 3. Health Check
```http
GET /health
```
**Response:** JSON with status information

### 4. File Reading Endpoint
```http
GET /read-file?file=filename
```
**Parameters:**
- `file` (optional): Filename to read (defaults to "sample.txt")

**Response:** File content with character count (text/plain)

## Postman Testing URLs

Test these endpoints in Postman:

### **Basic Endpoints:**
1. **Welcome Message**
   ```
   GET http://localhost:8080/
   ```

2. **Hello Endpoint (Default)**
   ```
   GET http://localhost:8080/hello
   ```

3. **Hello with Name Parameter**
   ```
   GET http://localhost:8080/hello?name=YourName
   ```

4. **Health Check (JSON Response)**
   ```
   GET http://localhost:8080/health
   ```

### **File Reading Endpoints:**
5. **Read Default File (sample.txt)**
   ```
   GET http://localhost:8080/read-file
   ```

6. **Read Specific File**
   ```
   GET http://localhost:8080/read-file?file=sample.txt
   ```

7. **Read Empty File**
   ```
   GET http://localhost:8080/read-file?file=empty.txt
   ```

8. **Test Error Handling (Non-existent file)**
   ```
   GET http://localhost:8080/read-file?file=nonexistent.txt
   ```

9. **Test Error Handling (Directory)**
   ```
   GET http://localhost:8080/read-file?file=.
   ```

## Command Line Usage

### File Reading Demonstrations

```bash
# Quick demonstration of file reading
go run . demo-file-reader

# Comprehensive testing suite
go run . test-file-reader

# Create additional test files
go run . create-test-files

# Clean up created test files
go run . cleanup-test-files
```

### Example Output

```bash
$ go run . demo-file-reader
=== File Reading Demo ===
This demonstrates the ReadFileContent function.
For comprehensive testing, use: go run . test-file-reader

Reading sample.txt:
✅ Successfully read 298 characters
Content preview: "Hello, World!\nThis is a sample file for testing..."
```

## File Reading Functions

### ReadFileContent(filePath string) (string, error)
**Main implementation using bufio.Scanner**
- Line-by-line reading with built-in buffering
- Memory efficient for large files
- Handles different line endings automatically
- Comprehensive error handling

### ReadFileContentSimple(filePath string) (string, error)
**Alternative implementation using os.ReadFile**
- Reads entire file at once
- Simpler code, less control
- Good for small to medium files

### StreamFileLines(filePath string, processor func(line string, lineNumber int) error) error
**Memory-efficient streaming processor**
- Processes files line by line
- Minimal memory footprint
- Suitable for very large files
- Customizable line processing

## Error Handling Examples

### HTTP Errors
```bash
# 404 - Route Not Found
curl http://localhost:8080/nonexistent
# Response: Route not found. (404)

# 405 - Method Not Allowed
curl -X POST http://localhost:8080/
# Response: Method Not Allowed. (405)
```

### File Reading Errors
```bash
# Non-existent file
curl "http://localhost:8080/read-file?file=missing.txt"
# Response: Error reading file 'missing.txt': failed to open file...

# Directory instead of file
curl "http://localhost:8080/read-file?file=."
# Response: Error reading file '.': '.' is a directory, not a file

# Empty file path
curl "http://localhost:8080/read-file?file="
# Response: Error reading file '': file path cannot be empty
```

## Security Considerations

### Implemented Headers
- **X-Content-Type-Options**: Prevents MIME type sniffing attacks
- **X-XSS-Protection**: Enables browser XSS filtering
- **X-Frame-Options**: Prevents clickjacking attacks

### File Security
- **Path Validation**: Prevents directory traversal attacks
- **File Type Checking**: Ensures only files (not directories) are read
- **Error Information**: Limited error details to prevent information leakage

## Testing the Server

### Using curl
```bash
# Test all HTTP endpoints
curl http://localhost:8080/
curl "http://localhost:8080/hello?name=Alice"
curl http://localhost:8080/health
curl http://localhost:8080/read-file

# Test file reading with different files
curl "http://localhost:8080/read-file?file=sample.txt"
curl "http://localhost:8080/read-file?file=empty.txt"

# Test error conditions
curl http://localhost:8080/invalid-route
curl -X POST http://localhost:8080/
curl "http://localhost:8080/read-file?file=nonexistent.txt"
```

### Using Command Line
```bash
# Test file reading functions directly
go run . demo-file-reader
go run . test-file-reader

# Create and test with large files
go run . create-test-files
go run . test-file-reader
go run . cleanup-test-files
```

### Using a Web Browser
Visit these URLs in your browser:
- http://localhost:8080/
- http://localhost:8080/hello?name=YourName
- http://localhost:8080/health
- http://localhost:8080/read-file
- http://localhost:8080/read-file?file=sample.txt

## Performance Considerations

### Memory Usage
- **bufio.Scanner**: Memory efficient for large files
- **os.ReadFile**: Loads entire file into memory
- **Streaming**: Minimal memory footprint for any file size

### File Size Recommendations
- **Small files (< 1MB)**: Use `ReadFileContentSimple`
- **Medium files (1-10MB)**: Use `ReadFileContent`
- **Large files (> 10MB)**: Use `StreamFileLines`

---

**🚀 Happy coding and testing!** 