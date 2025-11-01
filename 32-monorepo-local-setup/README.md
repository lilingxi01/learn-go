# Monorepo Local Setup

Build a Go monorepo with workspaces and shared packages.

## Prerequisites

- Completed [31-monorepo-introduction](../31-monorepo-introduction/)

## What You'll Learn

- Creating go.work file
- Shared package development
- Local module dependencies
- Running multiple services
- Development workflow

## Structure

```
32-monorepo-local-setup/
├── go.work              # Workspace file
├── shared/              # Shared packages
│   ├── logger/
│   ├── models/
│   └── database/
├── services/
│   ├── api/             # API service
│   └── worker/          # Background worker
├── Makefile
└── docker-compose.yml
```

## Running

```bash
cd 32-monorepo-local-setup

# Start dependencies
docker-compose up -d

# Run API
cd services/api && go run main.go

# Run Worker
cd services/worker && go run main.go
```

## Next Steps

Move to **33-monorepo-advanced** for CI/CD and deployment

## Further Reading

- [Go Workspaces Tutorial](https://go.dev/doc/tutorial/workspaces)

