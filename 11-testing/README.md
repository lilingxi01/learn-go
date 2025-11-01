# Testing in Go

Master Go's built-in testing framework and best practices.

## Prerequisites

- Completed [10-concurrency](../10-concurrency/)
- Understanding of functions and error handling

## What You'll Learn

- Writing test functions
- Table-driven tests
- Test coverage
- Benchmarking
- Example tests
- Test helpers and subtests
- Parallel tests
- Test fixtures

## Testing Basics

### Test File Naming

- Test files end with `_test.go`
- Test functions start with `Test`
- Example: `math.go` → `math_test.go`

### Test Function Signature

```go
func TestFunctionName(t *testing.T) {
    // Test code
}
```

## Running Tests

```bash
go test                    # Run tests in current directory
go test ./...              # Run all tests recursively
go test -v                 # Verbose output
go test -run TestName      # Run specific test
go test -cover             # Show coverage
go test -coverprofile=c.out # Generate coverage file
go tool cover -html=c.out  # View coverage in browser
```

## Table-Driven Tests

The Go idiom for comprehensive testing:

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name string
        a, b int
        want int
    }{
        {"positive", 2, 3, 5},
        {"negative", -2, -3, -5},
        {"zero", 0, 5, 5},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Add(tt.a, tt.b)
            if got != tt.want {
                t.Errorf("Add(%d, %d) = %d; want %d",
                    tt.a, tt.b, got, tt.want)
            }
        })
    }
}
```

## Test Assertions

Go's testing package is minimal. Use these patterns:

```go
// Basic assertion
if got != want {
    t.Errorf("got %v, want %v", got, want)
}

// Fatal error (stops test)
if err != nil {
    t.Fatalf("unexpected error: %v", err)
}

// Skip test
if condition {
    t.Skip("skipping test because...")
}
```

## Benchmarking

Test performance:

```go
func BenchmarkFunction(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Function()
    }
}
```

Run benchmarks:

```bash
go test -bench=.
go test -bench=. -benchmem  # Include memory stats
```

## Example Tests

Testable documentation:

```go
func ExampleAdd() {
    fmt.Println(Add(2, 3))
    // Output: 5
}
```

## Test Helpers

```go
func setup(t *testing.T) (*testState, func()) {
    t.Helper()  // Mark as helper for better error messages
    
    // Setup code
    state := &testState{}
    
    // Return cleanup function
    return state, func() {
        // Cleanup code
    }
}
```

## Parallel Tests

```go
func TestParallel(t *testing.T) {
    t.Parallel()  // Run in parallel with other parallel tests
    // Test code
}
```

## Running the Examples

```bash
cd 11-testing

# Run tests
go test -v

# Run specific test
go test -v -run TestAdd

# Run benchmarks
go test -bench=.

# Check coverage
go test -cover
```

## Best Practices

1. **Use table-driven tests**: Cover multiple cases
2. **Test edge cases**: Empty inputs, nil, zero values
3. **Test errors**: Ensure errors returned correctly
4. **Keep tests fast**: Use mocks for slow operations
5. **One assertion per test**: Or use subtests
6. **Test public API**: Don't test private functions directly
7. **Use t.Helper()**: For test helper functions

## Common Patterns

### Testing Errors

```go
func TestFunction(t *testing.T) {
    _, err := Function()
    if err == nil {
        t.Error("expected error, got nil")
    }
}
```

### Testing Panic

```go
func TestPanic(t *testing.T) {
    defer func() {
        if r := recover(); r == nil {
            t.Error("expected panic")
        }
    }()
    FunctionThatPanics()
}
```

### Comparing Structs

```go
got := Function()
want := ExpectedStruct{...}

if !reflect.DeepEqual(got, want) {
    t.Errorf("got %+v, want %+v", got, want)
}
```

## Test Coverage

Aim for high coverage, but don't obsess over 100%:

```bash
go test -cover
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

Good coverage metrics:
- **70-80%**: Acceptable
- **80-90%**: Good
- **90%+**: Excellent

## Common Mistakes

1. **Not using table-driven tests**: Missing edge cases
2. **Testing implementation**: Test behavior, not internals
3. **Slow tests**: Mock external dependencies
4. **No parallel tests**: Missing speedup opportunities
5. **Poor test names**: Use descriptive names

## Quick Reference

```go
// Basic test
func TestFunction(t *testing.T) {
    got := Function()
    want := expected
    if got != want {
        t.Errorf("got %v, want %v", got, want)
    }
}

// Table-driven test
tests := []struct {
    name string
    input int
    want int
}{
    {"case1", 1, 2},
}

for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        // Test
    })
}

// Benchmark
func BenchmarkFn(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Fn()
    }
}

// Example
func ExampleFn() {
    fmt.Println(Fn())
    // Output: expected
}
```

## Next Steps

1. Run the test examples
2. Write tests for previous lessons
3. Experiment with benchmarks
4. Move to **12-file-io** for file operations

## Further Reading

- [Go Testing Package](https://pkg.go.dev/testing)
- [Table Driven Tests](https://github.com/golang/go/wiki/TableDrivenTests)
- [Go Testing Best Practices](https://go.dev/blog/examples)
- [Advanced Testing with Go](https://www.youtube.com/watch?v=8hQG7QlcLBk)

