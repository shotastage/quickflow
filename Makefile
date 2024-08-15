# Go parameters
APP_NAME := quickflow
BIN_DIR := ./bin
SRC_DIRS := $(shell find . -type d -name '*')
PKG_LIST := $(shell go list ./... | grep -v /vendor/)

# Default task
.PHONY: all
all: build

# Run the application
.PHONY: run
run:
	@echo "Running the application..."
	go run main.go

# Build the application
.PHONY: build
build:
	@echo "Building the application..."
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) main.go

# Install dependencies
.PHONY: deps
deps:
	@echo "Installing dependencies..."
	go mod tidy

# Clean up generated files
.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -rf $(BIN_DIR)

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	go test -v $(PKG_LIST)

# Run only unit tests
.PHONY: unit-test
unit-test:
	@echo "Running unit tests..."
	go test -v ./test/unit/...

# Run only integration tests
.PHONY: integration-test
integration-test:
	@echo "Running integration tests..."
	go test -v ./test/integration/...

# Run lint check
.PHONY: lint
lint:
	@echo "Running lint..."
	golangci-lint run

# Run migrations (requires golang-migrate)
.PHONY: migrate
migrate:
	@echo "Running database migrations..."
	migrate -path ./migrations -database $(DATABASE_URL) up

# Generate Swagger documentation (assuming swag is installed)
.PHONY: swagger
swagger:
	@echo "Generating Swagger documentation..."
	swag init -g main.go

# Help
.PHONY: help
help:
	@echo "Usage:"
	@echo "  make run                Run the application"
	@echo "  make build              Build the application"
	@echo "  make deps               Install dependencies"
	@echo "  make clean              Clean up generated files"
	@echo "  make test               Run all tests"
	@echo "  make unit-test          Run only unit tests"
	@echo "  make integration-test   Run only integration tests"
	@echo "  make lint               Run lint checks"
	@echo "  make migrate            Run database migrations"
	@echo "  make swagger            Generate Swagger documentation"
	@echo "  make help               Show this help message"
