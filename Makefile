#!make
include .env

conStr := "mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}"

test:
	@echo ${conStr}

# Install required packages
install:
	go mod vendor

# Migration with migration tool
migrate-tool-install:
	go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Create a new migration
migrate-create:
	migrate create -ext sql -dir ./migrations -seq ${MIG_NAME}

migrate-up:
	@migrate -path ./migrations -database ${conStr} up ${STEP}

migrate-down:
	@migrate -path ./migrations -database ${conStr} down ${STEP}