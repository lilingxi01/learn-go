# HTTP Basics in Go

Learn to build HTTP servers and clients with Go's standard library.

## Prerequisites

- Completed [12-file-io](../12-file-io/)
- Understanding of functions and error handling

## What You'll Learn

- HTTP server basics
- Handling requests and responses
- Routing patterns
- HTTP methods (GET, POST, PUT, DELETE)
- Request parsing (query params, JSON body)
- HTTP client operations
- Middleware basics
- Common HTTP patterns

## Basic HTTP Server

```go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

## Request and Response

### Request Object

```go
func handler(w http.ResponseWriter, r *http.Request) {
    r.Method        // HTTP method (GET, POST, etc.)
    r.URL           // Request URL
    r.URL.Path      // URL path
    r.URL.Query()   // Query parameters
    r.Header        // HTTP headers
    r.Body          // Request body (io.ReadCloser)
}
```

### Response Writer

```go
func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("response"))
}
```

## Routing

### Multiple Handlers

```go
http.HandleFunc("/", homeHandler)
http.HandleFunc("/about", aboutHandler)
http.HandleFunc("/api/users", usersHandler)
```

### Method-Specific Handling

```go
func handler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        // Handle GET
    case http.MethodPost:
        // Handle POST
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}
```

## Parsing Requests

### Query Parameters

```go
// GET /search?q=golang&limit=10
q := r.URL.Query().Get("q")
limit := r.URL.Query().Get("limit")
```

### Form Data

```go
r.ParseForm()
name := r.FormValue("name")
email := r.FormValue("email")
```

### JSON Body

```go
var data MyStruct
decoder := json.NewDecoder(r.Body)
err := decoder.Decode(&data)
```

## HTTP Client

```go
// GET request
resp, err := http.Get("https://api.example.com")
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

body, err := io.ReadAll(resp.Body)

// POST request
data := bytes.NewBuffer(jsonData)
resp, err := http.Post("https://api.example.com", "application/json", data)
```

## Running the Examples

```bash
# Run server
go run server.go

# In another terminal, test the server
curl http://localhost:8080
curl http://localhost:8080/api/users

# Run client
go run client.go
```

## Common Patterns

### JSON Response

```go
func jsonResponse(w http.ResponseWriter, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}
```

### Error Response

```go
func errorResponse(w http.ResponseWriter, message string, code int) {
    w.WriteHeader(code)
    json.NewEncoder(w).Encode(map[string]string{
        "error": message,
    })
}
```

### Middleware Pattern

```go
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s", r.Method, r.URL.Path)
        next(w, r)
    }
}
```

## Best Practices

1. **Always defer Body.Close()**: On HTTP responses
2. **Set Content-Type header**: For responses
3. **Handle all HTTP methods**: Return 405 for unsupported
4. **Validate input**: Never trust user input
5. **Use context for cancellation**: http.Request.Context()
6. **Return proper status codes**: Use constants from net/http

## HTTP Status Codes

Common codes:
- `200 OK`: Success
- `201 Created`: Resource created
- `400 Bad Request`: Invalid input
- `401 Unauthorized`: Authentication required
- `404 Not Found`: Resource not found
- `500 Internal Server Error`: Server error

## Common Mistakes

1. **Not closing response body**: Memory leak
2. **Ignoring errors**: Always check error returns
3. **Wrong Content-Type**: Mismatched content and header
4. **Not validating input**: Security vulnerability
5. **Blocking operations**: Use goroutines for long tasks

## Quick Reference

```go
// Server
http.HandleFunc("/path", handler)
http.ListenAndServe(":8080", nil)

// Handler
func handler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("response"))
}

// Client GET
resp, err := http.Get("http://example.com")
defer resp.Body.Close()
body, _ := io.ReadAll(resp.Body)

// Client POST
resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))

// JSON response
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(data)

// Query params
q := r.URL.Query().Get("q")

// Path
path := r.URL.Path
```

## Next Steps

1. Run the HTTP server example
2. Test with curl or browser
3. Build a simple API
4. Move to **14-rest-api-fundamentals** for production APIs

## Further Reading

- [net/http Package](https://pkg.go.dev/net/http)
- [Writing Web Applications](https://go.dev/doc/articles/wiki/)
- [HTTP Server Examples](https://gobyexample.com/http-servers)
- [HTTP Client Examples](https://gobyexample.com/http-clients)

