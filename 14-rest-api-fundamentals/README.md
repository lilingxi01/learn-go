# REST API Fundamentals

Build RESTful APIs using Go and the chi router.

## Prerequisites

- Completed [13-http-basics](../13-http-basics/)
- Understanding of HTTP and JSON

## What You'll Learn

- REST principles and constraints
- HTTP methods and their meaning
- Status codes and when to use them
- RESTful URL design
- Chi router for routing
- Request/Response patterns
- API versioning basics

## What is REST?

REST (Representational State Transfer) is an architectural style for designing networked applications.

### REST Principles

1. **Client-Server**: Separation of concerns
2. **Stateless**: Each request contains all needed information
3. **Cacheable**: Responses can be cached
4. **Uniform Interface**: Consistent resource manipulation
5. **Layered System**: Client can't tell if connected directly to server
6. **Code on Demand** (optional): Server can extend client functionality

### RESTful Resource Design

```
GET    /users          # List all users
GET    /users/123      # Get user with ID 123
POST   /users          # Create new user
PUT    /users/123      # Update user 123 (full update)
PATCH  /users/123      # Update user 123 (partial)
DELETE /users/123      # Delete user 123
```

## HTTP Methods

| Method | Purpose                 | Idempotent | Safe |
| ------ | ----------------------- | ---------- | ---- |
| GET    | Retrieve resource       | ✅         | ✅   |
| POST   | Create resource         | ❌         | ❌   |
| PUT    | Update/Replace resource | ✅         | ❌   |
| PATCH  | Partial update          | ❌         | ❌   |
| DELETE | Delete resource         | ✅         | ❌   |

## HTTP Status Codes

### Success (2xx)

- `200 OK`: Standard success
- `201 Created`: Resource created
- `204 No Content`: Success, no body

### Client Errors (4xx)

- `400 Bad Request`: Invalid input
- `401 Unauthorized`: Authentication required
- `403 Forbidden`: Authenticated but not authorized
- `404 Not Found`: Resource doesn't exist
- `409 Conflict`: Conflict with current state
- `422 Unprocessable Entity`: Validation failed

### Server Errors (5xx)

- `500 Internal Server Error`: Generic server error
- `503 Service Unavailable`: Temporary unavailability

## Chi Router

Chi is a lightweight, idiomatic router for Go.

```go
r := chi.NewRouter()

r.Get("/users", listUsers)
r.Post("/users", createUser)
r.Get("/users/{id}", getUser)
r.Put("/users/{id}", updateUser)
r.Delete("/users/{id}", deleteUser)

http.ListenAndServe(":8080", r)
```

### Path Parameters

```go
func getUser(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    // Use id...
}
```

## Running the Example

```bash
cd 14-rest-api-fundamentals

# Install dependencies
go mod download

# Run the server
go run main.go

# Test with curl
curl http://localhost:8080/api/users
curl -X POST http://localhost:8080/api/users -d '{"name":"Alice","email":"alice@example.com"}'
curl http://localhost:8080/api/users/1
```

## Best Practices

1. **Use plural nouns for resources**: `/users`, not `/user`
2. **Use nouns, not verbs**: `/users`, not `/getUsers`
3. **Nest resources logically**: `/users/123/posts`
4. **Version your API**: `/v1/users` or header-based
5. **Return appropriate status codes**: Don't return 200 for everything
6. **Use JSON for request/response**: Unless you need something else
7. **Consistent error format**: Same error structure across API

## RESTful URL Design

### Good Examples

```
GET    /api/v1/users
GET    /api/v1/users/123
POST   /api/v1/users
GET    /api/v1/users/123/posts
GET    /api/v1/products?category=electronics&limit=10
```

### Bad Examples

```
GET    /api/v1/getUser?id=123        # Verb in URL
POST   /api/v1/users/create          # Redundant with method
GET    /api/v1/user                  # Singular (inconsistent)
DELETE /api/v1/deleteUserById/123    # Verb + redundant
```

## Error Response Format

Consistent error format helps clients:

```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Invalid email format",
    "details": {
      "field": "email",
      "value": "invalid-email"
    }
  }
}
```

## Common Patterns

### Response Wrapper

```go
type Response struct {
    Success bool        `json:"success"`
    Data    interface{} `json:"data,omitempty"`
    Error   *ErrorInfo  `json:"error,omitempty"`
}
```

### Pagination

```go
type PaginatedResponse struct {
    Data       []interface{} `json:"data"`
    Page       int          `json:"page"`
    PerPage    int          `json:"per_page"`
    Total      int          `json:"total"`
    TotalPages int          `json:"total_pages"`
}
```

## Common Mistakes

1. **Using GET for actions**: Use POST/PUT/DELETE
2. **Not using proper status codes**: 200 for everything
3. **Inconsistent URLs**: Mixed singular/plural
4. **Exposing implementation**: URLs tied to database structure
5. **No API versioning**: Breaking changes break clients

## Quick Reference

```go
// Chi router setup
r := chi.NewRouter()
r.Get("/resource", handler)
r.Post("/resource", handler)
r.Get("/resource/{id}", handler)
r.Put("/resource/{id}", handler)
r.Delete("/resource/{id}", handler)

// Get URL param
id := chi.URLParam(r, "id")

// JSON response
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(data)

// Error response
http.Error(w, "error message", http.StatusBadRequest)

// Status code
w.WriteHeader(http.StatusCreated)
```

## Next Steps

1. Run the REST API example
2. Test all endpoints with curl
3. Study the response patterns
4. Move to **15-dependency-injection-fx** for Uber FX

## Further Reading

- [REST API Tutorial](https://restfulapi.net/)
- [Chi Router](https://github.com/go-chi/chi)
- [HTTP Status Codes](https://httpstatuses.com/)
- [Best Practices for REST API Design](https://stackoverflow.blog/2020/03/02/best-practices-for-rest-api-design/)
