
DB_DSN=your_database_dsn_here  # Replace with your actual DSN

new-migration:
	goose create $(name) sql

migrate-up:
	goose -dir ./db/migrations postgres "$(DB_DSN)" up

migrate-down:
	goose -dir ./db/migrations postgres "$(DB_DSN)" down

.PHONY: new-migration  migrate-down