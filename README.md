# Learn Go - Interactive Production Service Tutorial

A comprehensive, hands-on tutorial for learning Go from basics to building production-ready services with APIs, databases, and deployment pipelines.

## 📚 What You'll Learn

This repository takes you from Go beginner to building production services, covering:

- **Go Fundamentals** (Folders 01-08): Syntax, data structures, functions, interfaces, error handling
- **Intermediate Go** (Folders 09-13): Packages, concurrency, testing, file I/O, HTTP basics
- **Production API with Uber FX** (Folders 14-19): REST APIs, dependency injection, middleware, validation
- **Database Integration** (Folders 20-26): PostgreSQL, GORM ORM, migrations, relationships
- **Complete Production Service** (Folder 27): Full-stack API with FX + GORM + PostgreSQL
- **Go Best Practices** (Folders 28-30): Code organization, documentation standards, community resources
- **Monorepo Architecture** (Folders 31-33): Go workspaces, shared packages, CI/CD pipelines
- **Advanced Topics** (Folders 34-36): Docker, performance optimization, production patterns

## 🎯 Who This Is For

- Developers learning Go for backend services
- Engineers transitioning from other languages
- Teams adopting Go for production APIs
- Anyone wanting to build scalable services with Go

## 🤖 Built with Cursor AI - Interactive Learning

This entire course was generated using [Cursor](https://cursor.com), an AI-powered code editor, following Meta principal engineer standards. This means you can:

### 💬 Ask Questions Directly in Your Editor

Cursor's AI can answer your Go questions with **semantic search** across this entire repository:

- **Ask about concepts**: "How do I handle errors in Go?"
- **Get help with challenges**: "I'm stuck on the FizzBuzz challenge in folder 03"
- **Understand code**: "Explain how the UserHandler works in folder 27"
- **Debug issues**: "Why am I getting this error in my goroutine?"

Cursor will search this repo, the web, and official Go documentation to provide contextual answers.

### 🔨 Easily Expand This Course

Want to add new topics? Cursor can help you extend this course while maintaining the same quality and style:

```
Ask Cursor: "Add a new folder 37-graphql-apis following the same 
pattern as other folders, with challenges and production examples"
```

Cursor will:
- ✅ Follow the established naming convention (padded numbers)
- ✅ Match the documentation style and structure
- ✅ Apply the same code quality standards (godoc, no magic numbers, error handling)
- ✅ Create challenges and solutions (for appropriate topics)
- ✅ Include working code examples and tests

The `.cursor/rules/` directory contains guidelines that ensure consistency:
- `guideline.mdc` - Code quality and structure standards
- `mentor-guide.mdc` - How to answer questions like a principal engineer

### 🚀 How to Use Cursor with This Repo

1. **Clone this repository**
   ```bash
   git clone <repo-url>
   cd learn-go
   ```

2. **Create a learning branch** (recommended)
   ```bash
   git checkout -b my-learning
   ```
   
   **Why?** This keeps the original course clean so you can:
   - Experiment freely without fear of breaking examples
   - Reset to original state anytime: `git checkout main`
   - Compare your solutions with originals
   - Try different approaches on separate branches

3. **Open in Cursor** (or install Cursor extension in VS Code)

4. **Start asking questions** using Cursor's chat (Cmd+L or Ctrl+L)

5. **Get contextual help** based on your current file and progress

**Example questions:**
- "What's the difference between buffered and unbuffered channels?"
- "Show me how to implement JWT authentication following this course's patterns"
- "Help me understand the FX dependency injection in folder 15"
- "Create a new folder for gRPC following the existing structure"

The AI understands the entire course structure and can guide you through any topic!

**Pro tip:** Create separate branches for experiments:
```bash
git checkout -b experiment/graphql-integration
git checkout -b practice/challenges
git checkout -b project/my-api
```

This way, you always have the clean original course to reference!

## ⚡ Quick Start

### Prerequisites

Before starting, ensure you have:

- **Go 1.21 or later** - [Download here](https://go.dev/dl/)
- **PostgreSQL** - See [folder 20](/20-postgresql-setup/) for setup options:
  - Local PostgreSQL installation
  - Supabase local development (recommended)
  - Docker PostgreSQL
- **Docker** (optional) - For containerization lessons
- **Git** - For version control
- A code editor (VS Code with Go extension recommended)

### Verify Your Setup

```bash
# Check Go installation
go version  # Should show 1.21 or later

# Check PostgreSQL (if installed locally)
psql --version
```

## 🚀 How to Use This Tutorial

### Structure

Each folder is numbered (01, 02, etc.) and contains:
- `README.md` - Concept explanations and instructions
- Playground files (`.go`) - Runnable code examples
- **Challenge files** (beginner sections) - Practice problems
- **Solution files** - Explained solutions to challenges
- **Complete modules** (advanced sections) - Full Go projects

### Running the Code

#### Standalone Files (Folders 01-13)

```bash
# Navigate to a folder
cd 01-getting-started

# Run a Go file
go run hello.go
```

#### Challenges (Folders 01-08)

1. Read the `README.md` to understand the concept
2. Try solving `challenge.go` yourself
3. Run your solution: `go run challenge.go`
4. Compare with `challenge-solution.go`

#### Complete Modules (Folders 14+)

```bash
# Navigate to module folder
cd 14-rest-api-fundamentals

# Install dependencies
go mod download

# Run the module
go run .

# Or run main.go directly
go run main.go
```

#### Running Tests

```bash
# Run tests in a module
go test ./...

# Run tests with coverage
go test -cover ./...

# Run benchmarks
go test -bench=. ./...
```

## 📖 Learning Path

### Beginner Path (Start Here!)

1. **01-getting-started** - Install Go and write Hello World
2. **02-basic-syntax** - Variables, types, constants
3. **03-control-flow** - If/else, loops, switch
4. **04-functions** - Functions, closures, defer
5. **05-data-structures** - Slices, maps, structs
6. **06-pointers** - Pointer basics
7. **07-interfaces** - Interface patterns
8. **08-error-handling** - Error handling in Go

### Intermediate Path

9. **09-packages-modules** - Go modules and packages
10. **10-concurrency** - Goroutines and channels
11. **11-testing** - Writing tests and benchmarks
12. **12-file-io** - File operations and JSON
13. **13-http-basics** - Basic HTTP servers

### Production API Path

14. **14-rest-api-fundamentals** - REST with chi router
15. **15-dependency-injection-fx** - Uber FX introduction
16. **16-api-with-fx-structure** - Structuring APIs with FX
17. **17-middleware-with-fx** - Middleware patterns
18. **18-request-validation** - Input validation
19. **19-configuration-logging** - Config and logging

### Database Path

20. **20-postgresql-setup** - PostgreSQL installation
21. **21-database-basics** - SQL and database/sql
22. **22-gorm-introduction** - GORM ORM basics
23. **23-migrations-setup** - Database migrations
24. **24-gorm-crud-operations** - CRUD with GORM
25. **25-gorm-relationships** - Model relationships
26. **26-gorm-advanced** - Transactions, hooks

### Production Service

27. **27-production-api-service** - Complete production API

### Best Practices

28. **28-code-organization** - Project structure
29. **29-documentation-standards** - Go documentation
30. **30-go-standards-resources** - Community standards

### Monorepo

31. **31-monorepo-introduction** - Monorepo concepts
32. **32-monorepo-local-setup** - Go workspaces
33. **33-monorepo-advanced** - CI/CD for monorepo

### Advanced Topics

34. **34-docker-containerization** - Docker best practices
35. **35-performance-optimization** - Profiling and optimization
36. **36-production-patterns** - Production-ready patterns

## 🗂️ Technology Stack

This tutorial uses production-grade tools and libraries:

- **Language**: Go 1.21+
- **HTTP Router**: [chi](https://github.com/go-chi/chi)
- **Dependency Injection**: [Uber FX](https://github.com/uber-go/fx)
- **ORM**: [GORM](https://gorm.io/)
- **Database**: PostgreSQL
- **Migrations**: [golang-migrate](https://github.com/golang-migrate/migrate)
- **Logging**: [zerolog](https://github.com/rs/zerolog)
- **Configuration**: [viper](https://github.com/spf13/viper)
- **Validation**: [go-playground/validator](https://github.com/go-playground/validator)

## 🤝 Contributing

We welcome contributions! Here's how to add new topics:

### Adding a New Topic

1. Create a new numbered folder (e.g., `37-new-topic/`)
2. Follow the naming convention with padded numbers
3. Include a `README.md` with:
   - Prerequisites
   - What You'll Learn
   - Running the Code
   - Further Reading
4. Add runnable code examples
5. For beginner topics, include challenges and solutions
6. Update this main README.md

### Code Standards

- Follow [Effective Go](https://go.dev/doc/effective_go)
- Use `gofmt` for formatting
- Add comments explaining patterns
- Include error handling examples
- Write tests for non-trivial code

## 📚 Additional Resources

### Official Go Resources

- [Go Documentation](https://go.dev/doc/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go by Example](https://gobyexample.com/)
- [Go Playground](https://go.dev/play/)
- [Go Tour](https://go.dev/tour/)

### Community Resources

- [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)
- [Google Go Style Guide](https://google.github.io/styleguide/go/)
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
- [Awesome Go](https://github.com/avelino/awesome-go)

### Books

- "The Go Programming Language" by Donovan & Kernighan
- "Learning Go" by Jon Bodner
- "Let's Go" by Alex Edwards

## 🔍 Getting Help

- Check the `README.md` in each folder for specific guidance
- Refer to [Folder 30](/30-go-standards-resources/) for Go standards
- Join the [Go community](https://go.dev/help/)
- Ask questions on [Stack Overflow](https://stackoverflow.com/questions/tagged/go)

## 📝 License

This tutorial is open source and available for educational purposes.

## 🌟 Acknowledgments

Built with inspiration from:
- [Official Go tutorials](https://go.dev/doc/tutorial/)
- Production Go services at leading tech companies
- The amazing Go community

---

**Ready to start?** Head to [01-getting-started](/01-getting-started/) and begin your Go journey!

