# Complete Production API Service

A production-ready API service integrating Uber FX, chi router, GORM, and PostgreSQL.

## Prerequisites

- Completed [26-gorm-advanced](../26-gorm-advanced/)
- PostgreSQL running locally

## What You'll Learn

- Complete production API architecture
- FX dependency injection throughout
- GORM with PostgreSQL integration
- Migration management
- Configuration management
- Structured logging
- Error handling patterns
- Testing strategies

## Project Structure

```
27-production-api-service/
├── cmd/
│   └── api/
│       └── main.go           # Entry point
├── internal/
│   ├── config/               # Configuration
│   ├── database/             # DB connection & migrations
│   ├── models/               # GORM models
│   ├── handlers/             # HTTP handlers
│   ├── services/             # Business logic
│   └── middleware/           # Middleware
├── migrations/               # SQL migrations
├── go.mod
├── go.sum
├── Makefile
├── config.example.yaml
└── README.md
```

## Features

- ✅ Uber FX dependency injection
- ✅ Chi router with middleware
- ✅ GORM ORM with PostgreSQL
- ✅ Database migrations (golang-migrate)
- ✅ Configuration management (Viper)
- ✅ Structured logging (zerolog)
- ✅ Request validation
- ✅ Error handling
- ✅ Graceful shutdown

## Running

```bash
cd 27-production-api-service

# Setup database (Docker)
docker-compose up -d

# Or use Supabase
supabase start

# Run migrations
make migrate-up

# Start server
go run cmd/api/main.go

# Test API
curl http://localhost:8080/api/users
```

## API Endpoints

```
GET    /api/health           - Health check
GET    /api/users            - List users
POST   /api/users            - Create user
GET    /api/users/{id}       - Get user
PUT    /api/users/{id}       - Update user
DELETE /api/users/{id}       - Delete user
GET    /api/users/{id}/posts - Get user's posts
```

## Next Steps

1. Run the complete service
2. Study the architecture
3. Add your own features
4. Move to **28-code-organization** for Go best practices

## Further Reading

- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

