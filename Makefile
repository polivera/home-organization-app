

# Set variables
DB_USER ?= root
DB_PASS ?= secret
DB_HOST ?= localhost
DB_PORT ?= 3306
DB_NAME ?= homeorg

conStr := "mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}"

test:
	@echo

# Install required packages
install:
	go mod vendor

# Migration with migration tool
migration-install-tool:
	go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Create a new migration
migration-create:
	migrate create -ext sql -dir ./migrations -seq ${MIG_NAME}

migration-up:
	migrate -path ./migrations -database ${conStr} up
