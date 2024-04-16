# Makefile for housefinder webAPP

#FILE PATHS
BEPATH = ./backend
FEPATH = ./frontend

# GO LANG VARS
GOCMD = go
GOBUILD = $(GOCMD) build
GOTEST = $(GOCMD) test
BINARY_NAME = housefinder

#BUN VARS


## BACKEND TASKS ##

# Build task
build:
	@$(GOBUILD) -o $(BEPATH)/bin/$(BINARY_NAME) $(BEPATH)

# Test task
test:
	@$(GOTEST) -v $(BEPATH)

run:build
	@$(BEPATH)/bin/$(BINARY_NAME)

## FRONT END TASKS ##
