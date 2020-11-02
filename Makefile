ifndef VERBOSE
.SILENT:
endif

binary="commit"

.PHONY: run build

run: build ## Build and run binary without arguments
	./$(binary)

build-linux-amd64: ## Build binary
	env GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o $(binary)-linux-amd64

build-darwin-amd64: ## Build binary
	env GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o $(binary)-darwin-amd64

build-windows-amd64: ## Build binary
	env GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o $(binary)-windows-amd64

debug-build: ## Build binary
	go build -o $(binary)

test: ## Run tests
	go test ./...

dependencies:
	go get ./...

cover: ## Run test-coverage and open in browser
	go test -v -covermode=count -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

quick-cover: ## Run simple coverage
	go test -cover ./...

fmt: ## Format source-tree
	gofmt -l -s -w .

help: ## Print all available make-commands
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
