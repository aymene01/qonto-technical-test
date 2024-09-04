# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=api
BUILD_DIR=bin

.PHONY: all build clean test

all: build
build:
	@mkdir -p $(BUILD_DIR)
	@$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) -v cmd/

clean:
	@$(GOCLEAN)
	@rm -f $(BUILD_DIR)/$(BINARY_NAME)

dev:
	@go run cmd/main.go cmd/routes.go

test:
	@$(GOTEST) -v ./test/...
