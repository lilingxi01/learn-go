# Packages and Modules in Go

Learn how to organize Go code with packages and modules.

## Prerequisites

- Completed [08-error-handling](../08-error-handling/)
- Understanding of Go basics

## What You'll Learn

- What packages are
- Creating and importing packages
- Package visibility (exported vs unexported)
- Go modules (go.mod)
- Dependency management
- Module organization best practices

## Packages

A package is a collection of Go source files in the same directory.

### Package Declaration

Every `.go` file starts with a package declaration:

```go
package main    // Executable
package mypackage  // Library
```

### Importing Packages

```go
import "fmt"
import "math"

// Or grouped
import (
    "fmt"
    "math"
)
```

### Package Naming

- Lowercase, single word
- Match directory name
- Descriptive but concise

## Visibility

Go uses capitalization for visibility:

```go
// Exported (public) - starts with capital letter
func PublicFunction() {}
type PublicStruct struct {}

// Unexported (private) - starts with lowercase
func privateFunction() {}
type privateStruct struct {}
```

## Go Modules

Modules are collections of packages with dependency tracking.

### Creating a Module

```bash
go mod init example.com/myproject
```

This creates `go.mod`:

```go
module example.com/myproject

go 1.21
```

### Adding Dependencies

```bash
go get github.com/user/package
```

Or just import and run:

```bash
go mod tidy
```

### go.mod Structure

```go
module example.com/myproject

go 1.21

require (
    github.com/user/package v1.2.3
    github.com/another/lib v2.0.0
)
```

## Project Structure

```
myproject/
├── go.mod
├── go.sum
├── main.go
├── internal/         # Private packages
│   └── utils/
│       └── helpers.go
├── pkg/             # Public packages
│   └── api/
│       └── client.go
└── cmd/             # Executables
    └── server/
        └── main.go
```

## Running the Example

```bash
cd 09-packages-modules

# Initialize module
go mod init example.com/tutorial

# Run main
go run main.go
```

## Common Commands

```bash
go mod init <module>   # Initialize module
go mod tidy            # Add missing, remove unused deps
go mod download        # Download dependencies
go mod vendor          # Copy dependencies to vendor/
go list -m all         # List all dependencies
go get <package>       # Add/update dependency
go get -u              # Update all dependencies
```

## Best Practices

1. **One package per directory**: All files in directory same package
2. **Use internal/ for private code**: Can't be imported by external projects
3. **Keep packages focused**: Single responsibility
4. **Avoid circular dependencies**: Package A imports B, B imports A (bad!)
5. **Use go.mod in root**: One module per repository (usually)

## Package Organization Patterns

### Flat Structure (Small Projects)

```
myapp/
├── go.mod
├── main.go
├── handlers.go
├── models.go
└── utils.go
```

### Standard Layout (Medium/Large)

```
myapp/
├── go.mod
├── cmd/
│   └── myapp/
│       └── main.go
├── internal/
│   ├── handlers/
│   ├── models/
│   └── database/
└── pkg/
    └── api/
```

## Common Mistakes

1. **Circular imports**: Causes compile error
2. **Wrong package name**: Must match directory (except main)
3. **Mixing package names**: All files in directory must use same package
4. **Not running go mod tidy**: Dependencies out of sync
5. **Importing internal/**: Can't import another project's internal/

## Quick Reference

```bash
# Create module
go mod init example.com/myapp

# Add dependency
go get github.com/user/package

# Update dependencies
go get -u
go mod tidy

# Vendor dependencies
go mod vendor
```

## Next Steps

1. Create your own package
2. Practice importing local packages
3. Work with external dependencies
4. Move to **10-concurrency** to learn goroutines and channels

## Further Reading

- [Go Modules Reference](https://go.dev/ref/mod)
- [How to Write Go Code](https://go.dev/doc/code)
- [Package Names](https://go.dev/blog/package-names)
- [Standard Package Layout](https://github.com/golang-standards/project-layout)
