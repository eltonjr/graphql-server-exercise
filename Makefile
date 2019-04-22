
.DEFAULT_GOAL:=help

.PHONY: help
help: ## prints this pretty message
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

.PHONY: build
build: ## builds the project using the 'go' command
	CGO_ENABLED=0 go build -v -a -installsuffix cgo ./...

.PHONY: run
run: ## runs the project
	go run main.go
