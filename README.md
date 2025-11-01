# Learn Go - Interactive Tutorial

A comprehensive, hands-on tutorial for learning Go from basics to production-ready services.

## 📚 What You'll Learn

- **Fundamentals** (01-08): Syntax, data structures, functions, interfaces, error handling
- **Intermediate** (09-13): Packages, concurrency, testing, file I/O, HTTP
- **Production APIs** (14-19): REST, Uber FX, middleware, validation, config
- **Database** (20-26): PostgreSQL, GORM, migrations, relationships
- **Complete Service** (27): Full production API with FX + GORM + PostgreSQL
- **Best Practices** (28-30): Code organization, documentation, standards
- **Monorepo** (31-33): Go workspaces, CI/CD pipelines
- **Advanced** (34-36): Docker, performance, production patterns

## ⚡ Quick Start

### Prerequisites

- Go 1.21+ ([download](https://go.dev/dl/))
- PostgreSQL (see [folder 20](/20-postgresql-setup/))
- Git and a code editor

### Get Started

```bash
# Clone and create a learning branch
git clone <repo-url>
cd learn-go
git checkout -b my-learning

# Start with the basics
cd 01-getting-started
go run hello.go
```

## 🚀 How to Use

### Running Code

**Standalone files (01-13):**
```bash
cd 02-basic-syntax
go run variables.go
```

**Modules (14-36):**
```bash
cd 14-rest-api-fundamentals
go mod download
go run .
```

**Tests:**
```bash
cd 11-testing
go test -v ./...
```

### Challenges (01-08)

Each beginner folder includes practice challenges:
1. Read the `README.md`
2. Try `challenge.go` yourself
3. Compare with `challenge-solution.go`

## 📖 Learning Path

### Beginner
- **01-getting-started** - Hello World and setup
- **02-basic-syntax** - Variables, types, constants
- **03-control-flow** - If/else, loops, switch
- **04-functions** - Functions, closures, defer
- **05-data-structures** - Slices, maps, structs
- **06-pointers** - Pointers and references
- **07-interfaces** - Interfaces and polymorphism
- **08-error-handling** - Error patterns

### Intermediate
- **09-packages-modules** - Go modules and packages
- **10-concurrency** - Goroutines and channels
- **11-testing** - Tests and benchmarks
- **12-file-io** - File operations and JSON
- **13-http-basics** - HTTP servers and clients

### Production API
- **14-rest-api-fundamentals** - REST with chi router
- **15-dependency-injection-fx** - Uber FX basics
- **16-api-with-fx-structure** - API architecture with FX
- **17-middleware-with-fx** - Middleware patterns
- **18-request-validation** - Input validation
- **19-configuration-logging** - Config and structured logging

### Database
- **20-postgresql-setup** - PostgreSQL installation
- **21-database-basics** - SQL and database/sql
- **22-gorm-introduction** - GORM ORM
- **23-migrations-setup** - Database migrations
- **24-gorm-crud-operations** - CRUD operations
- **25-gorm-relationships** - Model relationships
- **26-gorm-advanced** - Transactions and hooks

### Integration
- **27-production-api-service** - Complete production API

### Best Practices
- **28-code-organization** - Project structure
- **29-documentation-standards** - Go documentation
- **30-go-standards-resources** - Community standards

### Monorepo
- **31-monorepo-introduction** - Monorepo concepts
- **32-monorepo-local-setup** - Go workspaces
- **33-monorepo-advanced** - CI/CD pipelines

### Advanced
- **34-docker-containerization** - Docker best practices
- **35-performance-optimization** - Profiling and optimization
- **36-production-patterns** - Production-ready patterns

## 🤖 Cursor AI Integration

This course is designed for use with [Cursor](https://cursor.com), an AI-powered code editor.

### Interactive Learning

Ask questions directly in your editor (Cmd+L or Ctrl+L):
- "How do I handle errors in Go?"
- "Explain the FX dependency injection in folder 15"
- "I'm stuck on the challenge in folder 03"

Cursor's AI searches this repo, the web, and official Go docs to provide contextual answers.

### Expand the Course

Add new topics while maintaining quality:
```
Ask: "Add folder 37-grpc following the same pattern"
```

Cursor follows the guidelines in `.cursor/rules/` to ensure consistency.

### Branch Strategy

Create separate branches for experiments:
```bash
git checkout -b experiment/new-features
git checkout -b practice/challenges
```

This keeps the original course clean for reference.

## 🛠️ Technology Stack

- **Language**: Go 1.21+
- **Router**: [chi](https://github.com/go-chi/chi)
- **DI Framework**: [Uber FX](https://github.com/uber-go/fx)
- **ORM**: [GORM](https://gorm.io/)
- **Database**: PostgreSQL
- **Migrations**: [golang-migrate](https://github.com/golang-migrate/migrate)
- **Logging**: [zerolog](https://github.com/rs/zerolog)
- **Config**: [viper](https://github.com/spf13/viper)
- **Validation**: [go-playground/validator](https://github.com/go-playground/validator)

## 🤝 Contributing

Contributions welcome! See `.cursor/rules/guideline.mdc` for standards.

To add a topic:
1. Create numbered folder (e.g., `37-new-topic/`)
2. Include `README.md` and code examples
3. Follow existing patterns
4. Update this README

## 📚 Resources

### Official
- [Go Documentation](https://go.dev/doc/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go by Example](https://gobyexample.com/)

### Style Guides
- [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)
- [Google Go Style Guide](https://google.github.io/styleguide/go/)

### Books
- "The Go Programming Language" by Donovan & Kernighan
- "Learning Go" by Jon Bodner

## 🌟 Acknowledgments

**Created by**: [Lingxi Li](https://lingxi.li/) - Built with Cursor AI for learning production Go development. This course was created to support professional development while exploring an interactive, customizable learning methodology.

**Inspired by**:
- Official Go tutorials
- Production Go patterns from leading tech companies
- The Go community

---

**Ready to start?** → [01-getting-started](/01-getting-started/)
