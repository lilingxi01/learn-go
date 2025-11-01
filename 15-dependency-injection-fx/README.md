# Dependency Injection with Uber FX

Learn dependency injection patterns using Uber's FX framework.

## Prerequisites

- Completed [14-rest-api-fundamentals](../14-rest-api-fundamentals/)
- Understanding of interfaces and structs

## What You'll Learn

- What is dependency injection (DI)
- Why use DI in production services
- Uber FX basics
- Providing and consuming dependencies
- Application lifecycle hooks
- Module organization
- Testing with DI

## What is Dependency Injection?

Dependency Injection is a design pattern where objects receive their dependencies from external sources rather than creating them internally.

### Without DI (Tight Coupling)

```go
type UserService struct {
    db *Database
}

func NewUserService() *UserService {
    return &UserService{
        db: NewDatabase(), // Tightly coupled!
    }
}
```

### With DI (Loose Coupling)

```go
type UserService struct {
    db Database
}

func NewUserService(db Database) *UserService {
    return &UserService{
        db: db, // Injected dependency
    }
}
```

## Why Use Uber FX?

1. **Automatic dependency resolution**: FX wires dependencies automatically
2. **Lifecycle management**: Graceful startup and shutdown
3. **Testing friendly**: Easy to inject mocks
4. **Type-safe**: Compile-time dependency checking
5. **Production proven**: Used at Uber and many companies

## FX Basics

### Simple FX Application

```go
package main

import (
    "go.uber.org/fx"
)

func main() {
    fx.New(
        fx.Provide(NewLogger),
        fx.Provide(NewServer),
        fx.Invoke(func(*Server) {}),
    ).Run()
}
```

### Providing Dependencies

```go
// Constructor function
func NewLogger() *Logger {
    return &Logger{}
}

// Register with FX
fx.Provide(NewLogger)
```

### Consuming Dependencies

```go
// FX automatically injects logger
func NewServer(logger *Logger) *Server {
    return &Server{logger: logger}
}
```

### Lifecycle Hooks

```go
func NewServer(lc fx.Lifecycle, logger *Logger) *Server {
    server := &Server{logger: logger}

    lc.Append(fx.Hook{
        OnStart: func(ctx context.Context) error {
            // Start server
            return nil
        },
        OnStop: func(ctx context.Context) error {
            // Graceful shutdown
            return nil
        },
    })

    return server
}
```

## FX Modules

Group related dependencies:

```go
var DatabaseModule = fx.Module("database",
    fx.Provide(NewConnection),
    fx.Provide(NewRepository),
)

fx.New(
    DatabaseModule,
    // Other modules
)
```

## Running the Example

```bash
cd 15-dependency-injection-fx

# Install dependencies
go mod download

# Run the application
go run .

# Application will start and demonstrate FX lifecycle
```

## Best Practices

1. **One constructor per type**: Keep it simple
2. **Use interfaces for dependencies**: Flexibility
3. **Constructor naming**: Use `New` prefix
4. **Lifecycle hooks for resources**: Servers, DB connections
5. **Module organization**: Group related providers
6. **Avoid circular dependencies**: Design your modules carefully

## FX Patterns

### Optional Dependencies

```go
type Params struct {
    fx.In
    Logger *Logger `optional:"true"`
}

func NewServer(p Params) *Server {
    // Use p.Logger if available
}
```

### Named Dependencies

```go
type Params struct {
    fx.In
    ReadDB  *Database `name:"read"`
    WriteDB *Database `name:"write"`
}
```

### Value Groups

```go
fx.Provide(
    fx.Annotate(
        NewHandler1,
        fx.ResultTags(`group:"handlers"`),
    ),
)

type Params struct {
    fx.In
    Handlers []Handler `group:"handlers"`
}
```

## Testing with FX

```go
func TestService(t *testing.T) {
    var service *Service

    app := fxtest.New(t,
        fx.Provide(NewMockDB),
        fx.Provide(NewService),
        fx.Populate(&service),
    )
    defer app.RequireStart().RequireStop()

    // Test service
}
```

## Common Mistakes

1. **Circular dependencies**: A depends on B, B depends on A
2. **Too many dependencies**: Keep constructors focused
3. **Not using lifecycle hooks**: Servers don't shut down gracefully
4. **Over-engineering**: FX for simple apps is overkill
5. **Missing error returns**: Constructors should return errors

## Quick Reference

```go
// Basic FX app
fx.New(
    fx.Provide(NewDep),
    fx.Invoke(Run),
).Run()

// Provider
func NewService(dep *Dependency) *Service {
    return &Service{dep: dep}
}

// Lifecycle
func NewServer(lc fx.Lifecycle) *Server {
    lc.Append(fx.Hook{
        OnStart: func(ctx context.Context) error {
            // Start
        },
        OnStop: func(ctx context.Context) error {
            // Stop
        },
    })
    return server
}

// Module
var Module = fx.Options(
    fx.Provide(NewA),
    fx.Provide(NewB),
)
```

## Next Steps

1. Run the FX example
2. Study the dependency graph
3. Experiment with lifecycle hooks
4. Move to **16-api-with-fx-structure** for complete API with FX

## Further Reading

- [Uber FX Documentation](https://uber-go.github.io/fx/)
- [Dependency Injection in Go](https://blog.drewolson.org/dependency-injection-in-go)
- [FX GitHub](https://github.com/uber-go/fx)
- [Inversion of Control](https://martinfowler.com/articles/injection.html)
