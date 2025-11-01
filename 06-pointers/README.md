# Pointers in Go

Understand pointers, when to use them, and common patterns.

## Prerequisites

- Completed [05-data-structures](../05-data-structures/)
- Understanding of variables and structs

## What You'll Learn

- What pointers are and why they matter
- Pointer operators (\* and &)
- Passing by value vs passing by reference
- Pointer receivers vs value receivers
- When to use pointers
- Common pointer patterns

## What is a Pointer?

A pointer holds the memory address of a value.

```go
var x int = 42
var p *int = &x  // p points to x

fmt.Println(x)   // 42
fmt.Println(p)   // 0xc0000140a8 (memory address)
fmt.Println(*p)  // 42 (dereference pointer)
```

## Pointer Operators

- `&` - Address-of operator: Gets memory address
- `*` - Dereference operator: Gets value at address

```go
x := 10
p := &x    // p = address of x
*p = 20    // Set value at address to 20
// x is now 20
```

## Passing by Value vs Reference

### By Value (copies the data)

```go
func modifyValue(x int) {
    x = 100  // Only modifies the copy
}

num := 10
modifyValue(num)
fmt.Println(num)  // Still 10
```

### By Reference (modifies original)

```go
func modifyPointer(x *int) {
    *x = 100  // Modifies original
}

num := 10
modifyPointer(&num)
fmt.Println(num)  // Now 100
```

## Pointers with Structs

```go
type Person struct {
    Name string
    Age  int
}

// Value receiver - cannot modify
func (p Person) CantModify() {
    p.Age++  // Only modifies copy
}

// Pointer receiver - can modify
func (p *Person) CanModify() {
    p.Age++  // Modifies original
}
```

## When to Use Pointers

1. **Need to modify**: When function needs to modify argument
2. **Large structs**: Avoid copying large data structures
3. **Optional values**: Use `*T` for optional fields
4. **Sharing data**: Multiple parts need same data

## Running the Examples

```bash
go run pointers.go
```

## Challenge: Swap and Modify 🔄

Practice pointer manipulation!

**Requirements**:

1. Write a `swap` function that swaps two integers using pointers
2. Write a `modifyPerson` function that updates a Person struct
3. Write a `doubleValues` function that doubles all numbers in a slice
4. Demonstrate the difference between value and pointer receivers

Try it yourself before checking `challenge-solution.go`!

## Best Practices

1. **Use pointers for large structs**: Avoid expensive copies
2. **Use pointers to modify**: When function needs to change argument
3. **Receiver choice**: Use pointer receivers for modification or large types
4. **Nil checks**: Always check if pointer is nil before dereferencing
5. **Don't over-use**: Simple types (int, string) usually don't need pointers

## Common Patterns

### Nil Pointer Check

```go
func process(p *Person) {
    if p == nil {
        return
    }
    // Use p
}
```

### Pointer to Interface

```go
var i interface{} = (*Person)(nil)
```

## Common Mistakes

1. **Dereferencing nil pointer**: Causes panic
2. **Forgetting & or \***: Type mismatch errors
3. **Unnecessary pointers**: Using pointers for simple types
4. **Pointer to loop variable**: Common gotcha in closures

## Quick Reference

```go
// Create pointer
x := 42
p := &x        // Pointer to x
fmt.Println(*p) // Dereference: 42

// Modify via pointer
*p = 100       // x is now 100

// Function with pointer
func modify(p *int) {
    *p = 200
}
modify(&x)     // x is now 200

// Struct pointer
type T struct{ val int }
t := &T{val: 10}  // Pointer to struct
t.val = 20        // Auto-dereference
```

## Next Steps

1. Complete the swap and modify challenge
2. Practice with pointer receivers
3. Experiment with nil pointers
4. Move to **07-interfaces** for polymorphism

## Further Reading

- [A Tour of Go - Pointers](https://go.dev/tour/moretypes/1)
- [Effective Go - Pointers vs Values](https://go.dev/doc/effective_go#pointers_vs_values)
- [Go FAQ - Pointers](https://go.dev/doc/faq#Pointers)
