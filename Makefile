

# Set variables
DB_USER ?= root
DB_PASS ?= secret
DB_HOST ?= localhost
DB_PORT ?= 3306
DB_NAME ?= homeorg

conStr := "mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}"

test:
	@echo $(conStr)


# Install required packages
install:
	go mod vendor

# Migration with migration tool
migration-install-tool:
	go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migration-up:
	migrate -source file://migrations -database mysql://root:secret@tcp(192.168.0.152)/dbname?query up
