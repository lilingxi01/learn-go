# Monorepo Advanced - CI/CD and Deployment

Production CI/CD pipelines for Go monorepos.

## Prerequisites

- Completed [32-monorepo-local-setup](../32-monorepo-local-setup/)

## What You'll Learn

- Selective testing (only changed services)
- Selective building
- Multi-service deployment
- GitHub Actions for monorepo
- Docker multi-stage builds
- Change detection

## Structure

```
33-monorepo-advanced/
├── .github/
│   └── workflows/
│       ├── ci.yml              # CI pipeline
│       ├── deploy-api.yml      # API deployment
│       └── deploy-worker.yml   # Worker deployment
├── scripts/
│   ├── detect-changes.sh       # Detect changed services
│   └── build.sh                # Selective build
├── Dockerfile.api
├── Dockerfile.worker
└── Makefile
```

## CI/CD Strategy

1. **Detect Changes**: Which services changed?
2. **Selective Test**: Only test affected services
3. **Build**: Build only changed services
4. **Deploy**: Deploy only what changed

## Running

```bash
# Detect changes
./scripts/detect-changes.sh

# Build all services
make build-all

# Build specific service
make build-api
make build-worker
```

## Next Steps

Move to **34-docker-containerization** for Docker best practices

## Further Reading

- [GitHub Actions for Monorepos](https://github.blog/2021-11-29-using-github-actions-to-build-a-monorepo/)
- [Monorepo CI/CD](https://semaphoreci.com/blog/monorepo-cicd)

