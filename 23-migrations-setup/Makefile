.PHONY: migrate-up migrate-down migrate-create migrate-force

# Database connection
DB_URL=postgresql://postgres:postgres@localhost:5432/tutorial?sslmode=disable

# Run migrations up
migrate-up:
	migrate -path migrations -database "$(DB_URL)" up

# Run migrations down
migrate-down:
	migrate -path migrations -database "$(DB_URL)" down

# Create new migration
migrate-create:
	@if [ -z "$(NAME)" ]; then \
		echo "Usage: make migrate-create NAME=migration_name"; \
		exit 1; \
	fi
	migrate create -ext sql -dir migrations -seq $(NAME)

# Force migration version
migrate-force:
	@if [ -z "$(VERSION)" ]; then \
		echo "Usage: make migrate-force VERSION=version_number"; \
		exit 1; \
	fi
	migrate -path migrations -database "$(DB_URL)" force $(VERSION)

# Show migration version
migrate-version:
	migrate -path migrations -database "$(DB_URL)" version

