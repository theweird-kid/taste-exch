build:
	@go build -o bin/main cmd/main.go
test:
	@go test -v ./...
run: build
	@./bin/main
migrate:
	@cd ./cmd/migrations/schema && goose postgres postgres://postgres:postgres@localhost:5432/taste_db?sslmode=disable up
