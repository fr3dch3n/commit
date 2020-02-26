ifndef VERBOSE
.SILENT:
endif

binary="commit"

.PHONY: run build

run: build ## Build and run binary without arguments
	./$(binary)

build: ## Build binary
	go build -o $(binary)

test: ## Run tests
	go test -v ./...

cover: ## Run test-coverage and open in browser
	go test -v -covermode=count -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

help: ## Print all available make-commands
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
