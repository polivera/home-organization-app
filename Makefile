# Install required packages
install:
	go mod vendor

# Migration with migration tool
migration-install-tool:
	go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migration-up:
	migrate -source file://migrations -database mysql://root:secret@tcp(192.168.0.152)/dbname?query up