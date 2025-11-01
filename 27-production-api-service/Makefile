.PHONY: run build test migrate-up migrate-down docker-up docker-down

# Run the application
run:
	go run cmd/api/main.go

# Build the application
build:
	go build -o bin/api cmd/api/main.go

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# Database migrations
DB_URL=postgresql://postgres:postgres@localhost:5432/tutorial?sslmode=disable

migrate-up:
	migrate -path migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" down

migrate-create:
	@if [ -z "$(NAME)" ]; then \
		echo "Usage: make migrate-create NAME=migration_name"; \
		exit 1; \
	fi
	migrate create -ext sql -dir migrations -seq $(NAME)

# Docker
docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

# Install tools
install-tools:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Lint
lint:
	golangci-lint run

# Format code
fmt:
	gofmt -s -w .
	go mod tidy

# Help
help:
	@echo "Available commands:"
	@echo "  make run           - Run the application"
	@echo "  make build         - Build the application"
	@echo "  make test          - Run tests"
	@echo "  make migrate-up    - Run database migrations"
	@echo "  make migrate-down  - Rollback migrations"
	@echo "  make docker-up     - Start PostgreSQL with Docker"
	@echo "  make docker-down   - Stop PostgreSQL"

