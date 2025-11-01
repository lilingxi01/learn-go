# Docker Containerization for Go

Build optimized Docker images for Go applications.

## Prerequisites

- Completed [33-monorepo-advanced](../33-monorepo-advanced/)
- Docker installed

## What You'll Learn

- Multi-stage Docker builds
- Optimization techniques
- Security best practices
- Docker Compose for local development
- Image size reduction

## Multi-Stage Build

```dockerfile
# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o api cmd/api/main.go

# Runtime stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/api .
EXPOSE 8080
CMD ["./api"]
```

## Best Practices

1. **Multi-stage builds**: Smaller images
2. **Use alpine**: Minimal base image
3. **Layer caching**: Order matters
4. **`.dockerignore`**: Exclude unnecessary files
5. **Non-root user**: Security
6. **Static compilation**: CGO_ENABLED=0

## Running

```bash
cd 34-docker-containerization

# Build image
docker build -t myapp:latest .

# Run container
docker run -p 8080:8080 myapp:latest

# With docker-compose
docker-compose up
```

## Next Steps

Move to **35-performance-optimization** for profiling and optimization

## Further Reading

- [Docker Best Practices](https://docs.docker.com/develop/dev-best-practices/)
- [Dockerizing Go Apps](https://docs.docker.com/language/golang/)

