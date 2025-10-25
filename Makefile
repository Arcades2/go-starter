.PHONY: run build clean test lint fmt tidy check deps up down logs restart

# --- Variables ---
APP_NAME        := app
GO_CMD          := go
AIR_CMD         := air
DOCKER_COMPOSE  := docker compose
DB_CONTAINER    := postgres

# --- Run targets ---

## Lancer l'application avec Docker (mode développement)
run:
	@$(DOCKER_COMPOSE) up --build

## Redémarrer le conteneur
restart:
	@$(DOCKER_COMPOSE) restart

## Arrêter et supprimer les conteneurs
down:
	@$(DOCKER_COMPOSE) down

## Voir les logs (suivi temps réel)
logs:
	@$(DOCKER_COMPOSE) logs -f $(APP_NAME)

# --- Go commands ---

## Compiler le binaire Go
build:
	@echo "🔧 Build de l'application..."
	@$(GO_CMD) build -o bin/$(APP_NAME)
	@echo "✅ Build terminé : bin/$(APP_NAME)"

## Nettoyer les artefacts
clean:
	@echo "🧹 Nettoyage des fichiers compilés..."
	@rm -rf bin/
	@echo "✅ Nettoyage terminé"

## Exécuter les tests
test:
	@echo "🧪 Lancement des tests..."
	@$(GO_CMD) test -v ./...
	@echo "✅ Tests terminés"

# --- Code quality ---

## Formatter le code avec goimports + gofumpt
fmt:
	@echo "🎨 Formatage du code..."
	@goimports -w .
	@gofumpt -w .
	@echo "✅ Formatage terminé"

## Vérifier les dépendances
tidy:
	@echo "🧩 Vérification des dépendances..."
	@$(GO_CMD) mod tidy
	@echo "✅ Modules Go à jour"

## Lancer le linter complet
lint:
	@echo "🔍 Analyse statique avec golangci-lint..."
	@golangci-lint run ./...
	@echo "✅ Lint OK"

## Tout vérifier avant commit
check: fmt tidy lint test
	@echo "🚀 Code prêt à être commit !"

# --- Dépendances de dev ---

## Installer les outils nécessaires
deps:
	@echo "⬇️ Installation des outils..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install mvdan.cc/gofumpt@latest
	go install golang.org/x/tools/cmd/goimports@latest
	@echo "✅ Outils installés"