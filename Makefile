# Project Variables
BINARY_NAME = aws-comparator
BUILD_DIR = bin

# Default target when `make` is run
.PHONY: all
all: build ## Build the project

# Help target to auto-generate a list of targets and their descriptions
.PHONY: help
help:
	@echo "Available commands:"
	@awk 'BEGIN {FS = ":.*##"} /^[a-zA-Z_-]+:.*?##/ {printf "  make %-10s - %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

# Build the binary
.PHONY: build
build: ## Build the binary and place it in the bin directory
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) main.go

# Run the project
.PHONY: run
run: build ## Build and run the project
	./$(BUILD_DIR)/$(BINARY_NAME)

# Test the project
.PHONY: test
test: ## Run all tests
	go test ./... -v

# Format the code using gofmt
.PHONY: fmt
fmt: ## Format Go code with gofmt
	gofmt -w .

# Lint the project
.PHONY: lint
lint: ## Run golangci-lint to lint the code
	golangci-lint run

# Clean the build directory
.PHONY: clean
clean: ## Remove the bin directory and built files
	rm -rf $(BUILD_DIR)
