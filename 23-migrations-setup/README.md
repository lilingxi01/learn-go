# Database Migrations

Production-ready database migration strategies with golang-migrate.

## Prerequisites

- Completed [22-gorm-introduction](../22-gorm-introduction/)

## What You'll Learn

- Why migrations matter
- golang-migrate library
- Creating migrations
- Running migrations (up/down)
- Migration best practices
- Version control for schemas

## Why Migrations?

- Track schema changes over time
- Reproducible database setup
- Safe rollbacks
- Team collaboration
- Production deployments

## Running

```bash
cd 23-migrations-setup
go mod download

# Run migrations
make migrate-up

# Rollback
make migrate-down

# Create new migration
make migrate-create NAME=add_users_table
```

## Next Steps

Move to **24-gorm-crud-operations** for CRUD patterns
