# Makefile for housefinder API

# Variables
GOCMD = go
GOBUILD = $(GOCMD) build
GOTEST = $(GOCMD) test
BINARY_NAME = housefinder

# Build task
build:
	@$(GOBUILD) -o bin/$(BINARY_NAME)

# Test task
test:
	@$(GOTEST) -v ./...

run:build
	@./bin/$(BINARY_NAME)