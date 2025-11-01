# Functions in Go

Master Go functions, from basics to advanced patterns.

## Prerequisites

- Completed [03-control-flow](../03-control-flow/)
- Understanding of variables and control structures

## What You'll Learn

- Function declaration and syntax
- Parameters and return values
- Multiple return values
- Named return values
- Variadic functions
- Anonymous functions and closures
- Defer statement
- Function as values

## Basic Function Syntax

```go
func functionName(param1 type1, param2 type2) returnType {
    // function body
    return value
}
```

## Functions with Multiple Return Values

One of Go's best features:

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

## Named Return Values

```go
func calculate(a, b int) (sum int, product int) {
    sum = a + b
    product = a * b
    return // naked return
}
```

## Variadic Functions

Accept any number of arguments:

```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Call with any number of args
result := sum(1, 2, 3, 4, 5)
```

## Anonymous Functions

Functions without names:

```go
// Immediate execution
func() {
    fmt.Println("Anonymous function")
}()

// Assign to variable
add := func(a, b int) int {
    return a + b
}
result := add(3, 4)
```

## Closures

Functions that reference variables from outside:

```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

c := counter()
fmt.Println(c()) // 1
fmt.Println(c()) // 2
```

## Defer Statement

Defer execution until surrounding function returns:

```go
func example() {
    defer fmt.Println("World")
    fmt.Println("Hello")
}
// Outputs: Hello\nWorld
```

Common use: cleanup operations

```go
func readFile(filename string) {
    f, err := os.Open(filename)
    if err != nil {
        return
    }
    defer f.Close() // Ensure file is closed
    // Read file...
}
```

## Running the Examples

```bash
go run functions.go
go run closures.go
```

## Challenge: Calculator with Error Handling 🧮

Build a calculator that performs basic operations with proper error handling.

**Requirements**:

1. Create functions for: add, subtract, multiply, divide
2. All functions should return (result, error)
3. Handle division by zero
4. Create a main calculator function that takes operation type
5. Test with various inputs

Try it yourself before checking `challenge-solution/challenge-solution.go`!

## Best Practices

1. **Keep functions small**: One job per function
2. **Use multiple returns**: Return (value, error) pattern
3. **Name return values for clarity**: In complex functions
4. **Use defer for cleanup**: Files, connections, locks
5. **Descriptive names**: Function names should be verbs
6. **Exported functions**: Start with capital letter (public)

## Common Patterns

### Error Handling Pattern

```go
func doSomething() error {
    result, err := operation()
    if err != nil {
        return err
    }
    // use result
    return nil
}
```

### Options Pattern

```go
type Options struct {
    Timeout int
    Retries int
}

func NewServer(opts Options) *Server {
    // ...
}
```

## Common Mistakes

1. **Ignoring errors**: Always check returned errors
2. **Too many parameters**: Use structs for >3 parameters
3. **Forgetting defer order**: Last defer executes first (LIFO)
4. **Shadowing variables**: Be careful with := in if blocks

## Quick Reference

```go
// Basic function
func greet(name string) string {
    return "Hello, " + name
}

// Multiple returns
func swap(a, b int) (int, int) {
    return b, a
}

// Named returns
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

// Variadic
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

// Defer
func example() {
    defer fmt.Println("Last")
    fmt.Println("First")
}
```

## Next Steps

1. Complete the calculator challenge
2. Experiment with closures
3. Practice defer with file operations
4. Move to **05-data-structures** to learn about slices and maps

## Further Reading

- [A Tour of Go - Functions](https://go.dev/tour/moretypes)
- [Effective Go - Functions](https://go.dev/doc/effective_go#functions)
- [Go by Example - Functions](https://gobyexample.com/functions)
- [Go by Example - Closures](https://gobyexample.com/closures)
