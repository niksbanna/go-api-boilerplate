.PHONY: help build run test clean migrate-up migrate-down migrate-create lint fmt

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## Build the application
	@echo "Building application..."
	@go build -o bin/api cmd/api/main.go

run: ## Run the application
	@echo "Running application..."
	@go run cmd/api/main.go

test: ## Run tests
	@echo "Running tests..."
	@go test -v -race -coverprofile=coverage.out ./...

test-coverage: test ## Run tests with coverage report
	@go tool cover -html=coverage.out

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -rf bin/
	@rm -f coverage.out

migrate-up: ## Run database migrations up
	@echo "Running migrations up..."
	@migrate -path migrations -database "${DB_URL}" up

migrate-down: ## Run database migrations down
	@echo "Running migrations down..."
	@migrate -path migrations -database "${DB_URL}" down

migrate-create: ## Create a new migration file (usage: make migrate-create name=migration_name)
	@echo "Creating migration: $(name)"
	@migrate create -ext sql -dir migrations -seq $(name)

lint: ## Run linter
	@echo "Running linter..."
	@golangci-lint run

fmt: ## Format code
	@echo "Formatting code..."
	@go fmt ./...
	@gofmt -s -w .

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy

install-tools: ## Install development tools
	@echo "Installing tools..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
