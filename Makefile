# Binary names
BINARY_NAME=stats
CRON_BINARY=cron

# Directories
CMD_DIR=./cmd
STATS_CMD=$(CMD_DIR)/stats
CRON_CMD=$(CMD_DIR)/cron
MIGRATIONS_DIR=internal/infra/db/migrations

# Go vars
GO=go

# Tools
WIRE=go run github.com/google/wire/cmd/wire

# Docker Compose
DC=docker compose

PHONY: db-migrate wire-migrate

all: build

## 🏗 Build binaries
build:
	$(GO) build -o bin/$(BINARY_NAME) $(STATS_CMD)
	$(GO) build -o bin/$(CRON_BINARY) $(CRON_CMD)

## 🚀 Run the API server
run:
	$(GO) run $(STATS_CMD)

dev: run

## 🔁 Run cron job (manually)
cron:
	$(GO) run $(CRON_CMD)

## ✅ Run tests
test:
	$(GO) test ./...

## 🧵 Generate DI (Wire)
wire:
	$(WIRE) ./...

## 🗃 Run DB migrations
db-migrate:
	$(GO) run ./cmd/db migrate

db-rollback:
	$(GO) run ./cmd/db rollback

db-drop:
	$(GO) run ./cmd/db drop

## 🧽 Tidy Go modules
tidy:
	$(GO) mod tidy

## 🧹 Clean binaries
clean:
	rm -rf bin