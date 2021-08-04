generate: ## Run go generate
	go generate

lint: ## Lint code
	golangci-lint run

test: ## Test packages
	go test -count=1 -cover -coverprofile=coverage.out -v ./...

coverage: ## Test coverage with default output
	go tool cover -func=coverage.out

coverage-html: ## Test coverage with html output
	go tool cover -html=coverage.html

clean: ## Clean project
	rm -Rf ./bin
	rm -Rf coverage.out

build: clean ## Build local binary
	mkdir -p ./bin
	go build -o ./bin ./cmd/clamp

build-image: ## Build local image
	docker build -t ghcr.io/julienbreux/clamp:latest .

run: build ## Run local binary
	./bin/clamp

run-container: ## Run prepared local container
	docker run --rm julienbreux/clamp:latest

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: generate lint test coverage coverage-html clean build build-image run run-container help
