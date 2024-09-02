# Variables
GOOSE_CMD=goose
MIGRATIONS_DIR=./db/migrations
DB_DSN=your_database_dsn_here  # Replace with your actual DSN

# Default target
.PHONY: all
all: migrate-up

# Create a new migration
# Usage: make new-migration name=<migration_name>
.PHONY: new-migration
new-migration:
	goose create $(name) sql

# Apply all up migrations
.PHONY: migrate-up
migrate-up:
	goose -dir ./db/migrations postgres "$(DB_DSN)" up

# Rollback the last migration
.PHONY: migrate-down
migrate-down:
	goose -dir ./db/migrations postgres "$(DB_DSN)" down
