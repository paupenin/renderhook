# Makefile for a monorepo with Go (backend) and SvelteKit (frontend)

# Backend variables
BE_DIR := backend
BE_BUILD_DIR := $(BE_DIR)/bin

# Frontend variables
FE_DIR := frontend
FE_BUILD_DIR := $(FE_DIR)/build

# Common tasks
.PHONY: all clean dev

all: be fe
dev: be-dev fe-dev
test: be-test fe-test

# Backend tasks
.PHONY: be be-build be-test be-dev

be: be-build be-test
be-dev:
	@echo "Starting the Go backend in development mode..."
	cd $(BE_DIR) && gow run cmd/main.go

be-build:
	@echo "Building the Go backend..."
	cd $(BE_DIR) && go build -o $(BE_BUILD_DIR)/app cmd/main.go

be-test:
	@echo "Running tests for the Go backend..."
	cd $(BE_DIR) && go test ./...

# Frontend tasks
.PHONY: fe fe-build fe-test fe-dev

fe: fe-build fe-test
fe-dev:
	@echo "Starting the SvelteKit frontend in development mode..."
	cd $(FE_DIR) && npm run dev

fe-build:
	@echo "Building the SvelteKit frontend..."
	cd $(FE_DIR) && npm run build

fe-test:
	@echo "Running tests for the SvelteKit frontend..."
	cd $(FE_DIR) && npm test

# Clean task
clean:
	@echo "Cleaning up..."
	rm -rf $(BE_BUILD_DIR)/*
	rm -rf $(FE_BUILD_DIR)/*
