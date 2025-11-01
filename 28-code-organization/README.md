# Go Code Organization

Learn production repository structure and code organization patterns.

## Prerequisites

- Completed [27-production-api-service](../27-production-api-service/)

## What You'll Learn

- Standard Go project layout
- `/cmd`, `/internal`, `/pkg` directories
- Package organization principles
- Module boundaries
- Visibility and encapsulation

## Standard Go Project Layout

```
myproject/
├── cmd/                    # Main applications
│   ├── api/               # API server
│   │   └── main.go
│   └── worker/            # Background worker
│       └── main.go
├── internal/               # Private application code
│   ├── handlers/          # HTTP handlers
│   ├── services/          # Business logic
│   ├── models/            # Data models
│   └── database/          # Database code
├── pkg/                    # Public libraries
│   └── api/               # Public API client
├── api/                    # API definitions (OpenAPI, Protocol Buffers)
├── web/                    # Web assets
├── configs/                # Configuration files
├── scripts/                # Build/deployment scripts
├── test/                   # Additional test data
├── docs/                   # Documentation
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## Directory Purposes

### `/cmd`
- Main applications for this project
- Directory name matches executable name
- Keep code in `main.go` minimal
- Import and use `/internal` and `/pkg`

```
cmd/
├── api/          # API server entry point
├── worker/       # Background worker entry point
└── cli/          # CLI tool entry point
```

### `/internal`
- Private application code
- Cannot be imported by external projects
- Core business logic here
- Not exported outside your module

```
internal/
├── handlers/     # HTTP request handlers
├── services/     # Business logic services
├── repository/   # Data access layer
├── models/       # Domain models
├── middleware/   # HTTP middleware
└── database/     # Database setup
```

### `/pkg`
- Public libraries that external projects can import
- Use sparingly
- Well-documented public APIs
- Stable interfaces

```
pkg/
├── api/          # API client library
├── models/       # Shared models
└── utils/        # Utility functions
```

## Package Organization Principles

### 1. Package by Feature (Preferred)

```
internal/
├── users/
│   ├── handler.go
│   ├── service.go
│   ├── repository.go
│   └── model.go
├── posts/
│   ├── handler.go
│   ├── service.go
│   ├── repository.go
│   └── model.go
```

### 2. Package by Layer

```
internal/
├── handlers/
│   ├── user_handler.go
│   └── post_handler.go
├── services/
│   ├── user_service.go
│   └── post_service.go
├── repository/
│   ├── user_repo.go
│   └── post_repo.go
```

## Best Practices

1. **Use `/internal` for private code**: Enforces encapsulation
2. **Keep `main.go` small**: Just wiring, no logic
3. **One package per directory**: Clear organization
4. **Avoid deep nesting**: Max 3-4 levels
5. **Use descriptive names**: `userservice` not `us`

## Common Patterns

### Clean Architecture

```
internal/
├── domain/        # Business entities (independent)
├── usecase/       # Business logic
├── repository/    # Data access (interface)
└── delivery/      # HTTP/gRPC handlers
```

### Hexagonal Architecture

```
internal/
├── core/          # Business logic (independent)
├── ports/         # Interfaces
└── adapters/      # Implementations (HTTP, DB, etc.)
```

## Anti-Patterns

❌ **Don't:**
- Put everything in `main` package
- Create `utils` or `helpers` packages (too generic)
- Have circular dependencies
- Export everything
- Mix business logic with HTTP handlers

✅ **Do:**
- Use meaningful package names
- Keep packages focused
- Use `/internal` for private code
- Clear separation of concerns
- Dependency injection

## Example Project

See the `example-project/` folder for a complete demonstration.

## Next Steps

1. Review the example project structure
2. Refactor one of your projects using this layout
3. Move to **29-documentation-standards** for Go documentation

## Further Reading

- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
- [GopherCon 2018: How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0)
- [Package Oriented Design](https://www.ardanlabs.com/blog/2017/02/package-oriented-design.html)

