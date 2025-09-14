PHONY: run build clean test

APP_NAME = app
DOCKER_COMPOSE = docker compose
AIR_CMD = air
GO_CMD = go
DB_CONTAINER = postgres

run:
	@$(DOCKER_COMPOSE) up --build

build:
	@$(GO_CMD) build -o bin/$(APP_NAME)

clean:
	@rm -rf bin/

test:
	@$(GO_CMD) test -v ./...
