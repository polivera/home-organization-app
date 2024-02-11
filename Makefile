#!make
include .env

conStr := "mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}"

# Install required packages
install:
	go mod vendor

# Migration with migration tool
install-tools:
	go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/rinchsan/gosimports/cmd/gosimports@latest
	go install go.uber.org/mock/mockgen@latest

generate:
	go generate ./...

test-unit:
	go test -tags unit ./...

# Create a new migration
migrate-create:
	migrate create -ext sql -dir ./migrations -seq ${MIG_NAME}

migrate-up:
	@migrate -path ./migrations -database ${conStr} up ${STEP}

migrate-down:
	@migrate -path ./migrations -database ${conStr} down ${STEP}
