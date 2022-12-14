BINARY := ./main

.PHONY: all
all: clean mod build

.PHONY: clean
clean:
	@go clean $(BINARY)

.PHONY: mod
mod:
	@go mod download

.PHONY: build
build:
	@go build -o main ./...

# run with live reload
.PHONY: dev
dev:
	docker compose up

.PHONY: lint
lint:
	@go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run --fix ./...

.PHONY: gen-api
gen-api:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.11.0
	oapi-codegen --config config/oapi-codegen/server.yaml ./docs/openapi.yaml