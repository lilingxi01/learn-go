# PostgreSQL Setup

Complete guide to setting up PostgreSQL for Go development.

## Prerequisites

- Completed [19-configuration-logging](../19-configuration-logging/)
- Basic SQL knowledge helpful but not required

## What You'll Learn

- PostgreSQL installation options
- Supabase local development (recommended)
- Docker PostgreSQL setup
- Connection strings
- Database creation and management
- psql basics

## Installation Options

### Option 1: Supabase Local (Recommended)

Supabase provides PostgreSQL with additional features for local development.

**Install Supabase CLI:**

```bash
# macOS
brew install supabase/tap/supabase

# Linux/WSL
brew install supabase/tap/supabase

# Or with npm
npm install -g supabase
```

**Initialize Supabase:**

```bash
cd your-project
supabase init
supabase start
```

This starts:
- PostgreSQL on port 54322
- Studio UI on http://localhost:54323
- All credentials shown in terminal

**Connection String:**
```
postgresql://postgres:postgres@localhost:54322/postgres
```

### Option 2: Docker PostgreSQL

**Using Docker Compose:**

Create `docker-compose.yml`:

```yaml
version: '3.8'

services:
  postgres:
    image: postgres:16-alpine
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: tutorial
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data

volumes:
  postgres_data:
```

**Start:**

```bash
docker-compose up -d
```

**Connection String:**
```
postgresql://postgres:postgres@localhost:5432/tutorial
```

### Option 3: Local PostgreSQL Installation

**macOS:**

```bash
brew install postgresql@16
brew services start postgresql@16
```

**Ubuntu/Debian:**

```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
sudo systemctl start postgresql
```

**Windows:**

Download from [postgresql.org](https://www.postgresql.org/download/windows/)

**Connection String:**
```
postgresql://postgres:yourpassword@localhost:5432/tutorial
```

## Creating a Database

### Using psql

```bash
# Connect to PostgreSQL
psql -U postgres

# Create database
CREATE DATABASE tutorial;

# List databases
\l

# Connect to database
\c tutorial

# List tables
\dt

# Quit
\q
```

### Programmatically from Go

```go
import (
    "database/sql"
    _ "github.com/lib/pq"
)

db, err := sql.Open("postgres", "postgresql://postgres:postgres@localhost:5432/postgres")
if err != nil {
    log.Fatal(err)
}
defer db.Close()

_, err = db.Exec("CREATE DATABASE tutorial")
```

## Connection String Format

```
postgresql://[user]:[password]@[host]:[port]/[database]?[parameters]

# Examples:
postgresql://postgres:postgres@localhost:5432/mydb
postgresql://user:pass@localhost:5432/mydb?sslmode=disable
postgresql://user:pass@db.example.com:5432/prod?sslmode=require
```

## Testing Connection

```bash
# Using psql
psql postgresql://postgres:postgres@localhost:5432/tutorial

# Using Go (see connection.go example)
go run connection.go
```

## Supabase Features

Supabase provides additional features:

- **Studio UI**: Visual database management
- **Auth**: Built-in authentication
- **Storage**: File storage
- **Realtime**: Database change subscriptions
- **Edge Functions**: Serverless functions

Access Studio at: http://localhost:54323

## Best Practices

1. **Use environment variables**: Never hardcode credentials
2. **Connection pooling**: Reuse database connections
3. **SSL in production**: Always use `sslmode=require` in production
4. **Separate databases**: dev, test, production
5. **Regular backups**: Especially in production

## Common Commands

```bash
# Supabase
supabase start          # Start local Supabase
supabase stop           # Stop Supabase
supabase status         # Check status
supabase db reset       # Reset database

# Docker
docker-compose up -d    # Start PostgreSQL
docker-compose down     # Stop PostgreSQL
docker-compose logs     # View logs

# PostgreSQL
psql -U postgres        # Connect as postgres user
createdb tutorial       # Create database
dropdb tutorial         # Delete database
pg_dump tutorial > backup.sql  # Backup
psql tutorial < backup.sql     # Restore
```

## Troubleshooting

### Can't Connect

1. Check PostgreSQL is running: `ps aux | grep postgres`
2. Check port is correct: 5432 (local) or 54322 (Supabase)
3. Verify credentials
4. Check firewall settings

### Permission Denied

```bash
# Reset PostgreSQL password (macOS/Linux)
sudo -u postgres psql
ALTER USER postgres PASSWORD 'newpassword';
```

## Next Steps

1. Choose your setup method (Supabase recommended)
2. Verify connection with psql
3. Run the connection example
4. Move to **21-database-basics** for SQL and database/sql package

## Further Reading

- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [Supabase Documentation](https://supabase.com/docs)
- [PostgreSQL Tutorial](https://www.postgresqltutorial.com/)
- [pgAdmin](https://www.pgadmin.org/) - GUI tool for PostgreSQL

