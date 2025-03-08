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

# Create new migration files
new-migration:
	@read -p "Enter migration name: " name; \
	version=$$(date +%Y%m%d%H%M%S); \
	up_file="$(MIGRATIONS_DIR)/$${version}_$${name}.up.sql"; \
	down_file="$(MIGRATIONS_DIR)/$${version}_$${name}.down.sql"; \
	touch $$up_file $$down_file; \
	@echo "Created migration files: $$up_file and $$down_file"

# Help
help:
	@echo "Usage:"
	@echo "  make init-db         Initialize the database"
	@echo "  make migrate-db      Migrate the database"
	@echo "  make new-migration   Create new migration files"
	@echo "  make help            Show this help message"