# Getting Started with Go

Welcome to your Go journey! This lesson introduces you to the Go programming language and gets you writing your first program.

## Prerequisites

- A computer with internet access
- Basic understanding of programming concepts (variables, functions)

## What You'll Learn

- What Go is and why it's popular
- How to install Go
- How to write and run your first Go program
- Basic Go project structure
- Go's philosophy and design principles

## Introduction to Go

Go (often called Golang) is an open-source programming language created by Google in 2007 and released in 2009. It was designed by Robert Griesemer, Rob Pike, and Ken Thompson.

### Why Go?

Go was created to solve real problems at Google:

- **Fast compilation**: Go compiles quickly to machine code
- **Simplicity**: Easy to learn, easy to read
- **Concurrency**: Built-in support for concurrent programming
- **Performance**: Fast execution, efficient memory usage
- **Strong standard library**: Rich set of built-in packages
- **Cross-platform**: Compile for different OS easily

### Where is Go Used?

- **Cloud services**: Docker, Kubernetes, Terraform
- **Web services**: APIs, microservices
- **DevOps tools**: Command-line tools, infrastructure
- **Networking**: Proxies, load balancers
- **Databases**: CockroachDB, InfluxDB

## Installing Go

### Download Go

1. Visit [https://go.dev/dl/](https://go.dev/dl/)
2. Download the installer for your operating system
3. Run the installer

### Verify Installation

```bash
go version
```

You should see something like: `go version go1.21.0 darwin/amd64`

### Set Up Your Workspace

Go uses a workspace directory structure. By default:

- `~/go` on macOS/Linux
- `%USERPROFILE%\go` on Windows

Check your workspace:

```bash
go env GOPATH
```

## Your First Go Program

Look at `hello.go` in this folder - it's your first Go program!

### Program Structure

Every Go program has:

1. **Package declaration**: `package main` - the entry point
2. **Imports**: External packages you use
3. **Functions**: Code that does work
4. **main() function**: Where execution begins

### Running the Program

```bash
# Run directly
go run hello.go

# Or build and run
go build hello.go
./hello
```

### Understanding the Code

```go
package main  // This is an executable program
```

- `package main` means this is an executable (not a library)
- Every executable Go program must have a `main` package

```go
import "fmt"  // Import the format package
```

- `import` brings in code from other packages
- `fmt` is the format package (for printing, formatting)

```go
func main() {
    // Code here
}
```

- `func` declares a function
- `main()` is the entry point - execution starts here
- `{}` contain the function body

```go
fmt.Println("Hello, World!")
```

- Calls the `Println` function from the `fmt` package
- Prints text and adds a newline

## Challenge Time! 🎯

Now it's your turn! Open `challenge/challenge.go` and try to solve the challenge.

**Task**: Modify the program to:

1. Print a personalized greeting with your name
2. Print your favorite programming language
3. Print why you want to learn Go

Try to solve it yourself before looking at `challenge-solution/challenge-solution.go`!

## Go Commands to Know

```bash
# Run a program
go run main.go

# Build an executable
go build main.go

# Format your code (do this often!)
go fmt main.go

# Get help
go help

# Check Go environment
go env
```

## Go File Naming Conventions

- Use lowercase letters
- Separate words with underscores: `my_package.go`
- Or use camelCase: `myPackage.go`
- Test files end with `_test.go`

## Next Steps

After completing this lesson:

1. Try the challenge
2. Experiment with `fmt.Println()` - print different messages
3. Read the solution and compare with yours
4. Move on to **02-basic-syntax** to learn about variables

## Further Reading

- [A Tour of Go](https://go.dev/tour/) - Interactive Go tutorial
- [Effective Go](https://go.dev/doc/effective_go) - Official style guide
- [Go by Example](https://gobyexample.com/) - Learn by examples
- [Why Go?](https://go.dev/solutions/) - Use cases and case studies

## Common Mistakes

1. **Forgetting package declaration**: Every `.go` file needs `package <name>`
2. **Wrong package name**: Executables must use `package main`
3. **No main() function**: Executables need `func main()`
4. **Unused imports**: Go won't compile if you import but don't use a package

## Quick Reference

```go
// This is a complete, runnable Go program
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

---

**Congratulations!** 🎉 You've written your first Go program. Keep going!
