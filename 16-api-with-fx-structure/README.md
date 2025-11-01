# API with FX Structure

Build a production-structured API using Uber FX and chi router.

## Prerequisites

- Completed [15-dependency-injection-fx](../15-dependency-injection-fx/)
- Understanding of HTTP and FX basics

## What You'll Learn

- Structuring API with FX modules
- Separating concerns (config, server, handlers, routes)
- Dependency injection throughout the stack
- Lifecycle management for HTTP server
- Production-ready project structure

## Project Structure

```
16-api-with-fx-structure/
├── go.mod
├── main.go              # FX app initialization
├── config/
│   └── config.go        # Configuration module
├── server/
│   └── server.go        # HTTP server module
├── handlers/
│   └── user_handler.go  # HTTP handlers with DI
└── routes/
    └── routes.go        # Route registration
```

## Architecture

```
main.go
  ├─> Config Module
  ├─> Server Module (depends on Config, Routes)
  ├─> Handler Module (depends on Config)
  └─> Routes Module (depends on Handler)
```

## Running the Example

```bash
cd 16-api-with-fx-structure

go mod download
go run .

# Test endpoints
curl http://localhost:8080/api/users
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice","email":"alice@example.com"}'
```

## Best Practices

1. **Separate concerns**: Config, Server, Handlers, Routes
2. **Use interfaces**: For dependencies that might have multiple implementations
3. **Constructor injection**: All dependencies via constructors
4. **Lifecycle hooks**: For resources with startup/shutdown
5. **Module organization**: Group related providers

## Next Steps

1. Run and test the API
2. Study the FX dependency graph
3. Add your own handlers
4. Move to **17-middleware-with-fx** for middleware patterns

## Further Reading

- [FX Best Practices](https://uber-go.github.io/fx/best-practices.html)
- [Clean Architecture in Go](https://medium.com/@hatajoe/clean-architecture-in-go-4030f11ec1b1)
