# Load environment variables from .env file
include .env
export $(shell sed 's/=.*//' .env)

# Variables (optional, if you want to use them in the Makefile)
DSN := $(DSN)

# Targets
.PHONY: run
run:
	@echo "Running the application with DSN=$(DSN)"
	go run ./cmd

.PHONY: migrate
migrate:
	@echo "Running migrations with DSN=$(DSN)"
	@cd ./cmd/migrations/schema && goose postgres "$(DSN)" up

.PHONY: demigrate
demigrate:
	@echo "Running demigrations with DSN=$(DSN)"
	@cd ./cmd/migrations/schema && goose postgres "$(DSN)" down


.PHONY: test
test:
	@echo "Running tests with DSN=$(DSN)"
	go test ./...

.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -rf ./bin
