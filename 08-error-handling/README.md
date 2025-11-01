# Error Handling in Go

Master Go's explicit error handling patterns and best practices.

## Prerequisites

- Completed [07-interfaces](../07-interfaces/)
- Understanding of interfaces and functions

## What You'll Learn

- The error interface
- Creating and returning errors
- Error wrapping (Go 1.13+)
- Custom error types
- Panic and recover
- Error handling patterns
- Best practices

## The error Interface

```go
type error interface {
    Error() string
}
```

Any type with an `Error()` method satisfies the error interface.

## Creating Errors

### Using errors.New()

```go
import "errors"

func doSomething() error {
    return errors.New("something went wrong")
}
```

### Using fmt.Errorf()

```go
import "fmt"

func doSomething(id int) error {
    return fmt.Errorf("failed to process id %d", id)
}
```

## Error Handling Pattern

Always check errors:

```go
result, err := doSomething()
if err != nil {
    // Handle error
    return err
}
// Use result
```

## Error Wrapping (Go 1.13+)

Add context to errors:

```go
func outer() error {
    err := inner()
    if err != nil {
        return fmt.Errorf("outer failed: %w", err)
    }
    return nil
}
```

Unwrap errors:

```go
errors.Is(err, target)      // Check if err is target
errors.As(err, &target)     // Extract specific error type
errors.Unwrap(err)          // Get wrapped error
```

## Custom Error Types

```go
type ValidationError struct {
    Field string
    Value interface{}
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("invalid value %v for field %s",
        e.Value, e.Field)
}
```

## Panic and Recover

### Panic

Stops normal execution:

```go
panic("something terrible happened")
```

### Recover

Catch panics (only in deferred functions):

```go
func safe() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered:", r)
        }
    }()
    // Code that might panic
}
```

## Running the Examples

```bash
go run errors.go
go run custom-errors.go
```

## Challenge: Safe Division with Validation 🛡️

Build a robust calculator with comprehensive error handling!

**Requirements**:

1. Create custom error types for different failure modes
2. Implement safe division with proper error handling
3. Add input validation
4. Use error wrapping for context
5. Demonstrate panic/recover

Try it yourself before checking `challenge-solution/challenge-solution.go`!

## Best Practices

1. **Always check errors**: Don't ignore returned errors
2. **Add context**: Wrap errors with additional information
3. **Custom errors for specific cases**: When you need to handle differently
4. **Return early**: Check errors and return immediately
5. **Don't panic**: Use panic only for truly exceptional situations
6. **Log and return**: Don't do both (usually)

## Common Patterns

### Sentinel Errors

```go
var ErrNotFound = errors.New("not found")

if err == ErrNotFound {
    // Handle not found
}
```

### Error Types

```go
type TemporaryError interface {
    Temporary() bool
}

if te, ok := err.(TemporaryError); ok && te.Temporary() {
    // Retry
}
```

### Error Aggregation

```go
type MultiError []error

func (m MultiError) Error() string {
    // Combine error messages
}
```

## Common Mistakes

1. **Ignoring errors**: `result, _ := function()`
2. **Not adding context**: Returning raw errors
3. **Over-using panic**: panic for non-exceptional cases
4. **Not checking nil**: Calling methods on nil error
5. **Swallowing errors**: Logging but not returning

## Quick Reference

```go
// Create error
err := errors.New("error message")
err := fmt.Errorf("error: %s", details)

// Check error
if err != nil {
    return err
}

// Wrap error
return fmt.Errorf("context: %w", err)

// Check wrapped error
if errors.Is(err, ErrNotFound) {
    // Handle
}

// Extract error type
var pe *PathError
if errors.As(err, &pe) {
    // Use pe
}

// Custom error
type MyError struct{}
func (e MyError) Error() string {
    return "my error"
}

// Panic and recover
defer func() {
    if r := recover(); r != nil {
        // Handle panic
    }
}()
```

## Next Steps

1. Complete the safe division challenge
2. Create custom error types
3. Practice error wrapping
4. Move to **09-packages-modules** to learn about project organization

## Further Reading

- [Go Blog - Error Handling](https://go.dev/blog/error-handling-and-go)
- [Go Blog - Working with Errors](https://go.dev/blog/go1.13-errors)
- [Effective Go - Errors](https://go.dev/doc/effective_go#errors)
- [Errors are values](https://go.dev/blog/errors-are-values)
