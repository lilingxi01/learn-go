# Monorepo Introduction

Understand monorepo architecture and when to use it.

## Prerequisites

- Completed [30-go-standards-resources](../30-go-standards-resources/)

## What You'll Learn

- What is a monorepo
- Benefits and challenges
- Monorepo vs polyrepo
- When to use monorepos
- Go workspace feature
- Real-world examples

## What is a Monorepo?

A monorepo (monolithic repository) stores multiple projects/services in a single repository.

### Monorepo Structure

```
monorepo/
├── go.work           # Workspace file
├── shared/           # Shared libraries
│   ├── logger/
│   ├── models/
│   └── utils/
├── services/
│   ├── api/          # API service
│   ├── worker/       # Background worker
│   └── admin/        # Admin service
└── tools/            # Development tools
```

## Monorepo vs Polyrepo

### Monorepo

✅ **Advantages:**

- Atomic changes across services
- Shared code reuse
- Consistent tooling
- Easier refactoring
- Single source of truth

❌ **Challenges:**

- Larger repository
- CI/CD complexity
- Access control
- Tooling requirements

### Polyrepo

✅ **Advantages:**

- Smaller repositories
- Independent deployments
- Clearer ownership
- Simpler CI/CD

❌ **Challenges:**

- Code duplication
- Version management
- Cross-repo changes difficult
- Dependency hell

## When to Use Monorepo

Use monorepo when:

- Multiple services share code
- Team collaborates across services
- Atomic cross-service changes needed
- Consistent versioning required

Use polyrepo when:

- Services are truly independent
- Different teams/organizations
- Different release cycles
- Clear service boundaries

## Go Workspaces

Go 1.18+ introduces workspaces for monorepo support:

```bash
# Initialize workspace
go work init

# Add modules
go work use ./shared
go work use ./services/api
go work use ./services/worker
```

Creates `go.work`:

```
go 1.21

use (
    ./shared
    ./services/api
    ./services/worker
)
```

## Real-World Examples

Companies using Go monorepos:

- **Uber**: Microservices in monorepo
- **Monzo**: Financial services
- **Cloudflare**: Edge services

## Benefits in Production

1. **Shared Libraries**: Common logging, auth, models
2. **Atomic Updates**: Change API and client together
3. **Consistent Standards**: Same linters, formatters
4. **Simplified Dependencies**: No version conflicts
5. **Better Refactoring**: IDE refactors across services

## Challenges to Consider

1. **Build Time**: More code to compile
2. **CI/CD**: Need selective builds/tests
3. **Access Control**: Everyone sees everything (unless using git features)
4. **Merge Conflicts**: More frequent with more developers
5. **Repository Size**: Can grow large over time

## Solutions

- **Selective Testing**: Only test changed services
- **Caching**: Build cache for unchanged code
- **Code Owners**: Define ownership boundaries
- **Branch Protection**: Require reviews for specific paths
- **Clear Structure**: Good organization prevents chaos

## Next Steps

1. Understand the tradeoffs
2. Review the comparison charts
3. Move to **32-monorepo-local-setup** for hands-on implementation

## Further Reading

- [Go Workspaces](https://go.dev/doc/tutorial/workspaces)
- [Monorepo at Scale](https://semaphoreci.com/blog/what-is-monorepo)
- [Managing Go Dependencies in Monorepos](https://golang.howtos.io/managing-go-dependencies-in-monorepos/)
