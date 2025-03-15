# Makefile

# Variables
DB_URL=$(DATABASE_URL)
MIGRATE_CMD=go run cmd/migrate.go
MIGRATIONS_DIR=db/migrations

# Targets
.PHONY: init-db migrate-db new-migration

# Initialize the database
init-db:
	@echo "Initializing the database..."
	@$(MIGRATE_CMD) init

# Migrate the database
migrate-db:
	@echo "Migrating the database..."
	@$(MIGRATE_CMD) migrate

# Create new migration files=
new-migration:
	deno run --allow-read --allow-write ./scripts/create_migration.ts

# Help
help:
	@echo "Usage:"
	@echo "  make init-db         Initialize the database"
	@echo "  make migrate-db      Migrate the database"
	@echo "  make new-migration   Create new migration files"
	@echo "  make help            Show this help message"