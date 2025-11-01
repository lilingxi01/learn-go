# Data Structures in Go

Master Go's fundamental data structures: arrays, slices, maps, and structs.

## Prerequisites

- Completed [04-functions](../04-functions/)
- Understanding of variables and functions

## What You'll Learn

- Arrays vs Slices
- Slice operations and patterns
- Maps (hash tables)
- Structs and methods
- Pointer vs value receivers
- Composition

## Arrays

Fixed-size, same-type elements:

```go
var arr [5]int // Array of 5 integers
arr[0] = 10

// Initialize with values
numbers := [3]int{1, 2, 3}

// Let compiler count
names := [...]string{"Alice", "Bob", "Carol"}
```

Arrays are rarely used directly in Go - use slices instead!

## Slices

Dynamic-size, flexible:

```go
// Create slice
var s []int // nil slice

// Make slice with initial length and capacity
s = make([]int, 5)      // length 5, capacity 5
s = make([]int, 5, 10)  // length 5, capacity 10

// Slice literal
fruits := []string{"apple", "banana", "cherry"}

// Append elements
fruits = append(fruits, "date")

// Slice operations
sub := fruits[1:3] // Elements at index 1 and 2
```

### Key Slice Operations

```go
len(slice)    // Length
cap(slice)    // Capacity
append(slice, items...)  // Add elements
copy(dst, src)          // Copy elements
```

## Maps

Key-value pairs (hash table):

```go
// Create map
var m map[string]int // nil map

// Make map
m = make(map[string]int)

// Map literal
ages := map[string]int{
    "Alice": 25,
    "Bob": 30,
}

// Operations
ages["Carol"] = 28    // Add/update
age := ages["Alice"]  // Get value
delete(ages, "Bob")   // Delete entry

// Check existence
age, exists := ages["David"]
if exists {
    fmt.Println("Age:", age)
}
```

## Structs

Custom data types:

```go
type Person struct {
    Name string
    Age  int
    Email string
}

// Create struct
p := Person{
    Name: "Alice",
    Age: 25,
    Email: "alice@example.com",
}

// Access fields
fmt.Println(p.Name)
p.Age = 26
```

### Methods on Structs

```go
// Value receiver
func (p Person) Greet() string {
    return "Hello, " + p.Name
}

// Pointer receiver (can modify struct)
func (p *Person) HaveBirthday() {
    p.Age++
}
```

## Running the Examples

```bash
go run slices.go
go run maps.go
go run structs.go
```

## Challenge: Contact Book 📇

Build a contact management system!

**Requirements**:

1. Create a Contact struct with: Name, Phone, Email
2. Implement functions to:
   - Add a contact
   - Find contact by name
   - List all contacts
   - Delete a contact
3. Use a slice to store contacts
4. Bonus: Use a map for faster lookups

Try it yourself before checking `challenge-solution/challenge-solution.go`!

## Best Practices

1. **Prefer slices over arrays**: More flexible
2. **Check map existence**: Use two-value assignment
3. **Use pointers for large structs**: Avoid copying
4. **Initialize maps**: `make(map[K]V)` before use
5. **Export fields**: Capitalize for public access

## Common Patterns

### Slice Patterns

```go
// Remove element at index i
s = append(s[:i], s[i+1:]...)

// Filter slice
filtered := []int{}
for _, v := range slice {
    if condition(v) {
        filtered = append(filtered, v)
    }
}
```

### Map Patterns

```go
// Increment map value
m[key]++  // Works even if key doesn't exist (0 + 1)

// Map as set
set := make(map[string]bool)
set["item"] = true
if set["item"] {
    // Item exists in set
}
```

## Common Mistakes

1. **Slice capacity confusion**: Understand len vs cap
2. **Modifying slice in range**: Changes may not persist
3. **Nil map operations**: Must initialize with make()
4. **Passing large structs by value**: Use pointers
5. **Modifying slice during iteration**: Can cause issues

## Quick Reference

```go
// Slice
s := []int{1, 2, 3}
s = append(s, 4)
sub := s[1:3]

// Map
m := make(map[string]int)
m["key"] = 42
val, ok := m["key"]
delete(m, "key")

// Struct
type Person struct {
    Name string
    Age  int
}
p := Person{Name: "Alice", Age: 25}

// Method
func (p *Person) Method() {
    // ...
}
```

## Next Steps

1. Complete the contact book challenge
2. Experiment with slice operations
3. Practice map lookups
4. Move to **06-pointers** to understand pointers deeply

## Further Reading

- [A Tour of Go - Slices](https://go.dev/tour/moretypes/7)
- [Go Slices: usage and internals](https://go.dev/blog/slices-intro)
- [Effective Go - Maps](https://go.dev/doc/effective_go#maps)
- [Effective Go - Composite Literals](https://go.dev/doc/effective_go#composite_literals)
