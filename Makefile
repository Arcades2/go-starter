.PHONY: run build clean test lint fmt tidy check deps up down logs restart

# --- Variables ---
APP_NAME        := app
GO_CMD          := go
AIR_CMD         := air
DOCKER_COMPOSE  := docker compose
DB_CONTAINER    := postgres


# --- Docker commands ---

run:
	@$(DOCKER_COMPOSE) up --build

restart:
	@$(DOCKER_COMPOSE) restart

down:
	@$(DOCKER_COMPOSE) down

logs:
	@$(DOCKER_COMPOSE) logs -f $(APP_NAME)


# --- Go commands ---

build:
	@echo "🔧 Building the application..."
	@$(GO_CMD) build -o bin/$(APP_NAME)
	@echo "✅ Build complete: bin/$(APP_NAME)"

clean:
	@echo "🧹 Cleaning up build artifacts..."
	@rm -rf bin/
	@echo "✅ Cleanup complete"

test:
	@echo "🧪 Running tests..."
	@$(GO_CMD) test -v ./...
	@echo "✅ Tests completed"


# --- Code quality ---

fmt:
	@echo "🎨 Formatting code..."
	@goimports -w .
	@gofumpt -w .
	@echo "✅ Code formatted"

tidy:
	@echo "🧩 Checking dependencies..."
	@$(GO_CMD) mod tidy
	@echo "✅ Dependencies up to date"

lint:
	@echo "🔍 Running static analysis with golangci-lint..."
	@golangci-lint run ./...
	@echo "✅ Lint OK"

check: fmt tidy lint test
	@echo "🚀 Code ready to commit!"


# --- Development tools ---

deps:
	@echo "⬇️ Installing development tools..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install mvdan.cc/gofumpt@latest
	go install golang.org/x/tools/cmd/goimports@latest
	@echo "✅ Tools installed"