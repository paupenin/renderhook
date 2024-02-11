# Makefile for a monorepo with Go (backend) and SvelteKit (frontend)

# Backend variables
BE_DIR := backend
BE_BUILD_DIR := $(BE_DIR)/dist

# Frontend variables
FE_DIR := site
FE_BUILD_DIR := $(FE_DIR)/.next

# Common tasks
.PHONY: all clean dev

all: clean be fe
dev: be-dev fe-dev
test: be-test fe-test

# Backend tasks
.PHONY: be be-install be-dev be-build be-test

be: be-install be-build be-test

be-install:
	@echo "Installing dependencies for the Go backend..."
	cd $(BE_DIR) && go mod download

be-dev:
	@echo "Starting the Go backend in development mode..."
	cd $(BE_DIR) && gow run cmd/main.go

be-build:
	@echo "Building the Go backend..."
	cd $(BE_DIR) && go build -o dist/main cmd/main.go

be-test:
	@echo "Running tests for the Go backend..."
	cd $(BE_DIR) && go test ./...

# Frontend tasks
.PHONY: fe fe-build fe-test fe-dev

fe: fe-install fe-build fe-test

fe-install:
	@echo "Installing dependencies for the NextJS site..."
	cd $(FE_DIR) && bun install

fe-dev:
	@echo "Starting the NextJS site in development mode..."
	cd $(FE_DIR) && bun run dev

fe-build:
	@echo "Building the NextJS site..."
	cd $(FE_DIR) && bun run build

fe-test:
	@echo "Running tests for the NextJS site..."
	cd $(FE_DIR) && bun run test

# Clean task
clean:
	@echo "Cleaning up..."
	rm -Rf $(BE_BUILD_DIR)
	rm -Rf $(FE_BUILD_DIR)
