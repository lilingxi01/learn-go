# Interfaces in Go

Master Go's most powerful feature: interfaces for polymorphism and abstraction.

## Prerequisites

- Completed [06-pointers](../06-pointers/)
- Understanding of structs and methods

## What You'll Learn

- What interfaces are
- Implicit interface implementation
- Empty interface (interface{})
- Type assertions and type switches
- Common standard library interfaces
- Interface composition

## What are Interfaces?

Interfaces define behavior - a set of method signatures:

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

Any type that implements these methods satisfies the interface.

## Implicit Implementation

Go doesn't require explicit "implements" declarations:

```go
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// Rectangle automatically implements Shape!
```

## Empty Interface

`interface{}` can hold any type:

```go
var i interface{}
i = 42           // int
i = "hello"      // string
i = []int{1,2,3} // slice
```

Go 1.18+: Use `any` as alias for `interface{}`:

```go
var i any = "anything"
```

## Type Assertions

Extract the underlying value:

```go
var i interface{} = "hello"

// Type assertion
s := i.(string)
fmt.Println(s) // "hello"

// Safe type assertion
s, ok := i.(string)
if ok {
    fmt.Println(s)
}
```

## Type Switches

Check type of interface value:

```go
func describe(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Integer: %d\n", v)
    case string:
        fmt.Printf("String: %s\n", v)
    default:
        fmt.Printf("Unknown type\n")
    }
}
```

## Common Standard Library Interfaces

### fmt.Stringer

```go
type Stringer interface {
    String() string
}
```

### io.Reader and io.Writer

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
```

### error

```go
type error interface {
    Error() string
}
```

## Running the Examples

```bash
go run interfaces.go
go run type-assertions.go
```

## Challenge: Shape Calculator 📐

Build a shape calculator using interfaces!

**Requirements**:

1. Create a Shape interface with methods: Area(), Perimeter()
2. Implement for: Rectangle, Circle, Triangle
3. Create a function that accepts Shape interface
4. Calculate total area of multiple shapes
5. Use type assertions to print shape-specific info

Try it yourself before checking `challenge-solution.go`!

## Best Practices

1. **Small interfaces**: Prefer single-method interfaces
2. **Accept interfaces, return structs**: Flexible input, concrete output
3. **Define interfaces where used**: Not where implemented
4. **Use standard library interfaces**: io.Reader, fmt.Stringer, etc.
5. **Avoid interface pollution**: Don't create interfaces "just in case"

## Common Patterns

### Interface Segregation

```go
// Good: Small, focused interfaces
type Reader interface {
    Read() []byte
}

type Writer interface {
    Write([]byte)
}

// Compose when needed
type ReadWriter interface {
    Reader
    Writer
}
```

### Dependency Injection

```go
type Database interface {
    Query(string) []Record
}

type Service struct {
    db Database  // Accepts any implementation
}
```

## Common Mistakes

1. **Nil interface confusion**: `nil` interface ≠ interface with `nil` value
2. **Over-abstracting**: Creating interfaces too early
3. **Large interfaces**: Too many methods
4. **Forgetting pointer receivers**: Type with pointer receiver doesn't implement interface for value type

## Quick Reference

```go
// Define interface
type Shape interface {
    Area() float64
}

// Implement (implicit)
type Circle struct{ Radius float64 }
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

// Use interface
func PrintArea(s Shape) {
    fmt.Println(s.Area())
}

// Type assertion
var i interface{} = "hello"
s := i.(string)

// Type switch
switch v := i.(type) {
case string:
    // v is string
case int:
    // v is int
}
```

## Next Steps

1. Complete the shape calculator challenge
2. Implement fmt.Stringer for your types
3. Practice type assertions
4. Move to **08-error-handling** for robust error handling

## Further Reading

- [A Tour of Go - Interfaces](https://go.dev/tour/methods/9)
- [Effective Go - Interfaces](https://go.dev/doc/effective_go#interfaces)
- [Go Proverbs - Interfaces](https://go-proverbs.github.io/)
- [Interface Design in Go](https://rakyll.org/interface-pollution/)
