# Basic Syntax in Go

Learn the fundamental building blocks of Go: variables, types, constants, and operators.

## Prerequisites

- Completed [01-getting-started](../01-getting-started/)
- Go installed and working

## What You'll Learn

- Variable declarations and initialization
- Go's type system
- Constants
- Basic operators (arithmetic, comparison, logical)
- Type inference
- Zero values
- Type conversion

## Variables in Go

### Declaration Methods

Go has several ways to declare variables:

```go
// Method 1: var with type
var name string = "Alice"

// Method 2: var with type inference
var age = 25

// Method 3: Short declaration (most common inside functions)
count := 10

// Method 4: Multiple variables
var x, y, z int
```

### The := Operator

The `:=` operator is a shorthand for declaring and initializing variables:
- Can only be used inside functions
- Type is inferred from the value
- Most common in Go code

## Go's Basic Types

### Numeric Types

```go
// Integers
int, int8, int16, int32, int64        // Signed integers
uint, uint8, uint16, uint32, uint64   // Unsigned integers

// Floating point
float32, float64

// Complex numbers
complex64, complex128
```

### Other Types

```go
string  // Text
bool    // true or false
byte    // Alias for uint8
rune    // Alias for int32 (represents a Unicode code point)
```

### Zero Values

Variables declared without an initial value get a "zero value":
- `0` for numeric types
- `false` for booleans
- `""` (empty string) for strings
- `nil` for pointers, functions, interfaces, slices, channels, and maps

## Constants

Constants are declared with the `const` keyword:

```go
const Pi = 3.14159
const Greeting = "Hello"

// Multiple constants
const (
    StatusOK = 200
    StatusNotFound = 404
)
```

Constants cannot be changed after declaration.

## Operators

### Arithmetic Operators

```go
+   // Addition
-   // Subtraction
*   // Multiplication
/   // Division
%   // Modulus (remainder)
++  // Increment (postfix only)
--  // Decrement (postfix only)
```

### Comparison Operators

```go
==  // Equal to
!=  // Not equal to
<   // Less than
<=  // Less than or equal to
>   // Greater than
>=  // Greater than or equal to
```

### Logical Operators

```go
&&  // Logical AND
||  // Logical OR
!   // Logical NOT
```

## Type Conversion

Go requires explicit type conversion:

```go
var i int = 42
var f float64 = float64(i)  // Convert int to float64
var u uint = uint(f)        // Convert float64 to uint
```

No automatic type conversion - this prevents bugs!

## Running the Examples

```bash
# Run the variables example
go run variables.go

# Run the constants example
go run constants.go
```

## Challenge: Temperature Converter 🌡️

Open `challenge.go` and create a program that converts temperatures between Celsius, Fahrenheit, and Kelvin.

**Requirements**:
1. Declare a temperature in Celsius as a variable
2. Convert it to Fahrenheit using the formula: F = C × 9/5 + 32
3. Convert it to Kelvin using the formula: K = C + 273.15
4. Print all three values

Try it yourself before checking `challenge-solution.go`!

## Best Practices

1. **Use `:=` for local variables**: It's more concise
2. **Use `var` for package-level variables**: More explicit
3. **Name variables clearly**: Use descriptive names
4. **Use `const` for unchanging values**: Better performance and clarity
5. **Follow naming conventions**:
   - camelCase for variables: `myVariable`
   - PascalCase for exported (public) names: `MyFunction`

## Common Mistakes

1. **Using := outside functions**: Won't compile
2. **Redeclaring with :=**: Variable already declared error
3. **Unused variables**: Go won't compile with unused variables
4. **Mixing types without conversion**: Type mismatch error

## Quick Reference

```go
package main

import "fmt"

func main() {
    // Variables
    var name string = "Go"
    age := 15  // Years since Go was released
    
    // Constants
    const Version = "1.21"
    
    // Multiple declarations
    var x, y int = 1, 2
    
    // Zero values
    var count int        // 0
    var active bool      // false
    var text string      // ""
    
    // Type conversion
    var i int = 42
    var f float64 = float64(i)
    
    fmt.Println(name, age, Version, x, y, count, active, text, f)
}
```

## Next Steps

After mastering variables and types:
1. Complete the temperature converter challenge
2. Experiment with different types
3. Try type conversions
4. Move to **03-control-flow** to learn about if/else and loops

## Further Reading

- [A Tour of Go - Basics](https://go.dev/tour/basics)
- [Go Spec - Types](https://go.dev/ref/spec#Types)
- [Effective Go - Names](https://go.dev/doc/effective_go#names)

