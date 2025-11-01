# Go Documentation Standards

Master Go documentation conventions and godoc.

## Prerequisites

- Completed [28-code-organization](../28-code-organization/)

## What You'll Learn

- Godoc format and conventions
- Package documentation
- Function/method documentation
- Example functions
- Documentation tools
- Best practices

## Godoc Format

### Package Documentation

```go
// Package mathutil provides mathematical utility functions
// for common operations.
//
// Example usage:
//
//	result := mathutil.Add(2, 3)
//	fmt.Println(result) // Output: 5
package mathutil
```

### Function Documentation

```go
// Add returns the sum of two integers.
//
// Example:
//
//	sum := Add(2, 3) // Returns 5
func Add(a, b int) int {
    return a + b
}
```

### Type Documentation

```go
// User represents a user in the system.
// It contains personal information and metadata.
type User struct {
    ID    int    // Unique identifier
    Name  string // Full name
    Email string // Email address
}
```

### Method Documentation

```go
// Save persists the user to the database.
// Returns an error if the operation fails.
func (u *User) Save() error {
    // Implementation
}
```

## Documentation Tools

### Viewing Documentation

```bash
# View package documentation
go doc mathutil

# View function documentation
go doc mathutil.Add

# View all documentation
go doc -all mathutil

# Start documentation server
godoc -http=:6060
# Visit http://localhost:6060
```

### Generating Documentation

```bash
# Install godoc (if needed)
go install golang.org/x/tools/cmd/godoc@latest

# Or use pkgsite (new official tool)
go install golang.org/x/pkgsite/cmd/pkgsite@latest
pkgsite
```

## Example Functions

Testable documentation examples:

```go
func ExampleAdd() {
    result := Add(2, 3)
    fmt.Println(result)
    // Output: 5
}

func ExampleAdd_negative() {
    result := Add(-2, -3)
    fmt.Println(result)
    // Output: -5
}
```

## Best Practices

1. **Start with function/type name**: `// Add returns...`
2. **Full sentences**: End with period
3. **No empty lines**: Between comment and declaration
4. **Include examples**: When helpful
5. **Explain why, not what**: Code shows what, docs explain why
6. **Keep it concise**: Brief but complete
7. **Use proper formatting**: Indented code blocks

## Documentation Checklist

- [ ] Package has package-level documentation
- [ ] All exported functions have godoc comments
- [ ] All exported types have comments
- [ ] Complex logic explained
- [ ] Examples provided for non-obvious usage
- [ ] No typos or grammar errors
- [ ] Comments are up-to-date with code

## Common Mistakes

1. **Missing package documentation**
2. **Not starting with element name**
3. **Empty lines between comment and code**
4. **Documenting obvious things**
5. **Out-of-date comments**

## Next Steps

1. Review the documented examples
2. Document one of your packages
3. Run `go doc` on your code
4. Move to **30-go-standards-resources** for community standards

## Further Reading

- [Effective Go - Commentary](https://go.dev/doc/effective_go#commentary)
- [Godoc: documenting Go code](https://go.dev/blog/godoc)
- [Go Doc Comments](https://go.dev/doc/comment)

