# Control Flow in Go

Master if/else statements, switch cases, and loops in Go.

## Prerequisites

- Completed [02-basic-syntax](../02-basic-syntax/)
- Understanding of variables and operators

## What You'll Learn

- If/else statements
- Switch statements
- For loops (the only loop in Go!)
- Loop control (break, continue)
- Range loops
- Infinite loops
- Nested control structures

## If Statements

Go's if statements don't need parentheses around conditions:

```go
if x > 10 {
    fmt.Println("x is greater than 10")
}
```

### If-Else

```go
if x > 10 {
    fmt.Println("Greater than 10")
} else {
    fmt.Println("10 or less")
}
```

### If-Else-If

```go
if score >= 90 {
    fmt.Println("A")
} else if score >= 80 {
    fmt.Println("B")
} else if score >= 70 {
    fmt.Println("C")
} else {
    fmt.Println("D or F")
}
```

### If with Short Statement

You can declare variables in the if statement:

```go
if age := getAge(); age >= 18 {
    fmt.Println("Adult")
} else {
    fmt.Println("Minor")
}
// age is only available inside the if/else block
```

## Switch Statements

Switch in Go is cleaner than in many languages - no need for break!

### Basic Switch

```go
switch day {
case "Monday":
    fmt.Println("Start of week")
case "Friday":
    fmt.Println("TGIF!")
case "Saturday", "Sunday":
    fmt.Println("Weekend!")
default:
    fmt.Println("Midweek")
}
```

### Switch with Expression

```go
switch {
case score >= 90:
    fmt.Println("A")
case score >= 80:
    fmt.Println("B")
default:
    fmt.Println("C or below")
}
```

### Switch with Short Statement

```go
switch num := getValue(); num {
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two")
}
```

## For Loops

Go has only one loop keyword: `for`. But it's very flexible!

### Traditional For Loop

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

### While-Style Loop

```go
count := 0
for count < 5 {
    fmt.Println(count)
    count++
}
```

### Infinite Loop

```go
for {
    // Loop forever (until break)
    if condition {
        break
    }
}
```

### Range Loop

```go
// Loop over slice
numbers := []int{1, 2, 3, 4, 5}
for index, value := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}

// Loop over string (runes)
for i, char := range "Hello" {
    fmt.Printf("%d: %c\n", i, char)
}
```

## Loop Control

### Break

Exit the loop immediately:

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break // Exit when i is 5
    }
    fmt.Println(i)
}
```

### Continue

Skip to the next iteration:

```go
for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue // Skip even numbers
    }
    fmt.Println(i) // Only prints odd numbers
}
```

## Running the Examples

```bash
go run conditionals.go
go run loops.go
```

## Challenge: FizzBuzz with Custom Rules 🎯

The classic FizzBuzz challenge with a twist!

**Standard FizzBuzz**:
- Print numbers from 1 to 100
- For multiples of 3, print "Fizz"
- For multiples of 5, print "Buzz"
- For multiples of both 3 and 5, print "FizzBuzz"
- Otherwise, print the number

**Your Task**: Implement FizzBuzz with these additional rules:
- For multiples of 7, print "Boom"
- Handle combinations (e.g., 15 is "FizzBuzz", 21 is "FizzBoom", 35 is "BuzzBoom")

Try it yourself before checking `challenge-solution.go`!

## Best Practices

1. **Omit parentheses**: `if x > 10` not `if (x > 10)`
2. **Use switch for multiple conditions**: More readable than if-else chains
3. **Use range for slices/arrays**: More idiomatic
4. **Use `_` to ignore values**: `for _, value := range slice`
5. **Keep it simple**: Avoid deeply nested conditionals

## Common Patterns

### Early Return

```go
func divide(a, b float64) float64 {
    if b == 0 {
        return 0 // Early return on error case
    }
    return a / b
}
```

### Guard Clauses

```go
func processUser(user User) {
    if user == nil {
        return
    }
    if !user.IsActive {
        return
    }
    // Process user
}
```

## Common Mistakes

1. **Adding parentheses around conditions**: Not needed in Go
2. **Forgetting braces**: Even for single-line blocks, use `{}`
3. **Using break in switch**: Not needed (auto-breaks)
4. **Wrong range values**: Remember it's `index, value` not `value, index`

## Quick Reference

```go
// If statement
if condition {
    // code
}

// If-else
if condition {
    // code
} else {
    // code
}

// Switch
switch value {
case 1:
    // code
case 2:
    // code
default:
    // code
}

// For loop
for i := 0; i < 10; i++ {
    // code
}

// Range loop
for index, value := range collection {
    // code
}

// While-style
for condition {
    // code
}

// Infinite
for {
    // code
}
```

## Next Steps

1. Complete the FizzBuzz challenge
2. Try different loop variations
3. Experiment with switch statements
4. Move to **04-functions** to learn about functions

## Further Reading

- [A Tour of Go - Flow Control](https://go.dev/tour/flowcontrol)
- [Effective Go - Control Structures](https://go.dev/doc/effective_go#control-structures)
- [Go by Example - If/Else](https://gobyexample.com/if-else)
- [Go by Example - For](https://gobyexample.com/for)

