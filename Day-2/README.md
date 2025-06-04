# Day-2 HTTP Server

## Project Overview

This is a comprehensive Go HTTP server implementation following best practices and modern web development standards. The server demonstrates proper routing, error handling, middleware implementation, and security considerations.

## Functional Requirements Implementation

### ✅ Routes
- **GET /** - Returns "Welcome to the Go HTTP Server" (text/plain)
- **GET /hello?name=YourName** - Returns "Hello, YourName!" (text/plain), defaults to "Guest"
- **GET /health** - Returns 200 OK with JSON `{"status": "ok"}`

### ✅ Error Handling
- **404 (Unknown Routes)** - Returns "Route not found."
- **405 (Invalid Methods)** - Returns "Method Not Allowed."
- **500 (Server Errors)** - Proper internal error handling

### ✅ Headers & Security
- **Content-Type** headers set appropriately for each endpoint
- **Security Headers** implemented:
  - `X-Content-Type-Options: nosniff`
  - `X-XSS-Protection: 1; mode=block`
  - `X-Frame-Options: DENY`
  - Comments on additional production security headers

### ✅ Code Structure & Best Practices
- **http.NewServeMux** for better routing control
- **Modular handlers** with clear separation of concerns
- **defer** statements for cleanup
- **Comprehensive logging middleware** with method, URL, status, duration, and user-agent
- **Proper error checking** throughout the codebase
- **Idiomatic Go** patterns and conventions

## Project Structure

```
Day-2/
├── main.go      # Complete HTTP server implementation
├── go.mod       # Go module file
└── README.md    # This documentation
```

## Features

### 🔒 Security Features
- Basic security headers to prevent common web vulnerabilities
- Method validation for all endpoints
- Proper error responses without information leakage
- Input validation and sanitization

### 📊 Logging & Monitoring
- Request logging middleware with:
  - Timestamp
  - HTTP method and path
  - Response status code
  - Request duration
  - User-Agent string

### 🏗️ Architecture
- **Modular Design**: Each handler has a single responsibility
- **Middleware Pattern**: Logging implemented as middleware
- **Error Handling**: Centralized error response functions
- **Type Safety**: Proper struct definitions for JSON responses

## How to Run

### Prerequisites
- Go 1.21 or later installed

### Quick Start

```bash
# Navigate to Day-2 directory
cd Day-2

# Run the server
go run main.go
```

### Build and Run

```bash
# Build the executable
go build -o http-server main.go

# Run the built executable
./http-server
```

## API Endpoints

### 1. Root Endpoint
```http
GET /
```
**Response:** `Welcome to the Go HTTP Server` (text/plain)

**Example:**
```bash
curl http://localhost:8080/
```

### 2. Hello Endpoint
```http
GET /hello?name=YourName
```
**Parameters:**
- `name` (optional): Your name (defaults to "Guest")

**Response:** `Hello, YourName!` (text/plain)

**Examples:**
```bash
# With name parameter
curl "http://localhost:8080/hello?name=John"
# Response: Hello, John!

# Without name parameter
curl http://localhost:8080/hello
# Response: Hello, Guest!
```

### 3. Health Check
```http
GET /health
```
**Response:** JSON with status information

**Example:**
```bash
curl http://localhost:8080/health
# Response: {"status":"ok"}
```

## Error Handling Examples

### 404 - Route Not Found
```bash
curl http://localhost:8080/nonexistent
# Response: Route not found. (404)
```

### 405 - Method Not Allowed
```bash
curl -X POST http://localhost:8080/
# Response: Method Not Allowed. (405)
```

## Code Architecture

### Handler Functions
- `homeHandler` - Handles root path requests
- `helloHandler` - Handles personalized greetings
- `healthHandler` - Provides health check endpoint

### Error Handlers
- `handleNotFound` - 404 error responses
- `handleMethodNotAllowed` - 405 error responses
- `handleInternalServerError` - 500 error responses

### Middleware
- `loggingMiddleware` - Request logging and monitoring
- `responseWriter` - Custom response writer for status code capture

### Utility Functions
- `setCommonHeaders` - Security and response headers

## Security Considerations

### Implemented Headers
- **X-Content-Type-Options**: Prevents MIME type sniffing attacks
- **X-XSS-Protection**: Enables browser XSS filtering
- **X-Frame-Options**: Prevents clickjacking attacks

### Additional Production Headers (Commented)
- **Strict-Transport-Security**: Forces HTTPS connections
- **Content-Security-Policy**: Controls resource loading
- **Referrer-Policy**: Controls referrer information

## Testing the Server

### Using curl
```bash
# Test all endpoints
curl http://localhost:8080/
curl "http://localhost:8080/hello?name=Alice"
curl http://localhost:8080/health

# Test error conditions
curl http://localhost:8080/invalid-route
curl -X POST http://localhost:8080/
```

### Using a Web Browser
Visit these URLs in your browser:
- http://localhost:8080/
- http://localhost:8080/hello?name=YourName
- http://localhost:8080/health

## Logging Output Example

When running the server, you'll see logs like:
```
Starting HTTP server on port :8080...
Available endpoints:
  GET /           - Welcome message
  GET /hello      - Hello message (supports ?name=YourName)
  GET /health     - Health check
2024/01/15 10:30:45 [2024-01-15 10:30:45] GET / - Status: 200 - Duration: 123.456µs - User-Agent: curl/7.68.0
```

## Best Practices Demonstrated

1. **Separation of Concerns**: Each function has a single responsibility
2. **Error Handling**: Comprehensive error checking and user-friendly responses
3. **Security**: Multiple layers of security headers and input validation
4. **Logging**: Detailed request logging for monitoring and debugging
5. **Code Organization**: Clean, readable, and maintainable code structure
6. **HTTP Standards**: Proper use of HTTP status codes and headers
7. **Documentation**: Comprehensive inline comments and external documentation

## Production Considerations

For production deployment, consider:
- Enable additional security headers
- Implement rate limiting
- Add authentication/authorization
- Use HTTPS with TLS certificates
- Implement graceful shutdown
- Add health check dependencies
- Configure proper logging levels
- Add metrics and monitoring 