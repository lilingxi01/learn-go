# Go Standards and Resources

Comprehensive guide to Go coding standards and where to find them.

## Prerequisites

- Completed [29-documentation-standards](../29-documentation-standards/)

## Official Go Resources

### Essential Reading

1. **[Effective Go](https://go.dev/doc/effective_go)** - THE definitive style guide
   - Formatting
   - Commentary
   - Names
   - Control structures
   - Data structures
   - Concurrency

2. **[Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)** - Common review feedback
   - Don't panic
   - Error strings
   - Handle errors
   - Naming conventions
   - Package names

3. **[Go FAQ](https://go.dev/doc/faq)** - Frequently asked questions
   - Design decisions
   - Types and allocations
   - Concurrency
   - Performance

## Community Style Guides

### Uber Go Style Guide

[https://github.com/uber-go/guide/blob/master/style.md](https://github.com/uber-go/guide/blob/master/style.md)

Key points:

- Prefer interfaces over concrete types
- Use functional options pattern
- Avoid global variables
- Structured logging
- Error wrapping

### Google Go Style Guide

[https://google.github.io/styleguide/go/](https://google.github.io/styleguide/go/)

Emphasizes:

- Readability
- Simplicity
- Consistency
- Error handling
- Testing

## Code Quality Tools

### golangci-lint

The gold standard linter (aggregates many linters):

```bash
# Install
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run
golangci-lint run

# With auto-fix
golangci-lint run --fix
```

**Configuration** (`.golangci.yml`):

```yaml
linters:
  enable:
    - gofmt
    - govet
    - errcheck
    - staticcheck
    - unused
    - gosimple
    - structcheck
    - ineffassign
```

### staticcheck

Excellent static analysis:

```bash
go install honnef.co/go/tools/cmd/staticcheck@latest
staticcheck ./...
```

### Other Tools

```bash
# Race detector
go test -race ./...

# Code coverage
go test -cover ./...

# Benchmarks
go test -bench=. ./...

# Format
gofmt -s -w .

# Simplify
gofumpt -w .
```

## Naming Conventions

### Packages

- lowercase
- short, single word
- no underscores
- match directory name

```go
✅ package user
✅ package http
❌ package user_service
❌ package UserService
```

### Variables

- camelCase for locals
- Short names in short scopes
- Descriptive names in larger scopes

```go
✅ i, j, k          // Loop counters
✅ r, w             // Reader, Writer
✅ userService      // Descriptive
❌ usr              // Too abbreviated
❌ UserService      // Exported unnecessarily
```

### Functions

- PascalCase for exported
- camelCase for unexported
- Verbs for functions
- Nouns for types

```go
✅ func GetUser()
✅ func validateInput()
❌ func get_user()
❌ func User()  // Confusing
```

### Constants

- PascalCase or SCREAMING_SNAKE_CASE
- Exported if needed by other packages

```go
✅ const MaxRetries = 3
✅ const MAX_SIZE = 100
✅ const DefaultTimeout = 30 * time.Second
```

## Error Handling Patterns

### Standard Pattern

```go
if err != nil {
    return fmt.Errorf("context: %w", err)
}
```

### Error Types

```go
var ErrNotFound = errors.New("not found")

type ValidationError struct {
    Field string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed: %s", e.Field)
}
```

## Testing Conventions

```go
// TestFunctionName tests the FunctionName function
func TestFunctionName(t *testing.T) {
    tests := []struct {
        name string
        // test fields
    }{
        {name: "case1"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // test code
        })
    }
}
```

## Code Review Checklist

Before committing:

- [ ] `gofmt -w .` (formatted)
- [ ] `go vet ./...` (no issues)
- [ ] `golangci-lint run` (passes)
- [ ] `go test ./...` (all pass)
- [ ] All errors checked
- [ ] No magic numbers
- [ ] Godoc on exports
- [ ] Meaningful names

## Quick Wins

1. **Run gofmt**: `gofmt -w .`
2. **Run go vet**: `go vet ./...`
3. **Check unused**: `go mod tidy`
4. **Run linter**: `golangci-lint run`
5. **Add godoc**: Document all exports

## Resources by Topic

### Concurrency

- [Go Concurrency Patterns](https://go.dev/talks/2012/concurrency.slide)
- [Advanced Concurrency Patterns](https://go.dev/talks/2013/advconc.slide)

### Testing

- [Table Driven Tests](https://github.com/golang/go/wiki/TableDrivenTests)
- [Testing Best Practices](https://go.dev/doc/tutorial/add-a-test)

### Performance

- [Profiling Go Programs](https://go.dev/blog/pprof)
- [High Performance Go Workshop](https://dave.cheney.net/high-performance-go-workshop/dotgo-paris.html)

### Security

- [Go Security Policy](https://go.dev/security)
- [Writing Secure Go Code](https://github.com/OWASP/Go-SCP)

## Community Resources

- [Awesome Go](https://github.com/avelino/awesome-go) - Curated list
- [Go Forum](https://forum.golangbridge.org/) - Community discussion
- [r/golang](https://reddit.com/r/golang) - Reddit community
- [Gophers Slack](https://gophers.slack.com/) - Real-time chat

## Next Steps

1. Bookmark the essential resources
2. Configure golangci-lint for your projects
3. Review your code against these standards
4. Move to **31-monorepo-introduction** for monorepo patterns

## Further Reading

- [Effective Go](https://go.dev/doc/effective_go)
- [Go Proverbs](https://go-proverbs.github.io/)
- [Go at Google](https://go.dev/talks/2012/splash.article)
