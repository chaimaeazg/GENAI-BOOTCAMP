# Build
build:
	go build -o bin/server cmd/server/main.go

# Run
run:
	go run cmd/server/main.go

# Test
test:
	go test -v ./...

# Clean
clean:
	rm -rf bin/
	rm -rf logs/
	rm -f *.db

# Migrate
migrate:
	sqlite3 db/words.db < db/migrations/0001_initial_schema.sql
	sqlite3 db/words.db < db/migrations/0002_relationships.sql
	sleep 2
	curl -X POST http://localhost:8080/api/import-initial-data

# Lint
lint:
	golangci-lint run

# Add new target for running mage
mage:
	go run tools/mage/mage.go

.PHONY: build run test clean migrate lint mage
