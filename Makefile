# Makefile for housefinder webAPP

#FILE PATHS
BEPATH = ./backend
FEPATH = ./frontend

# GO LANG VARS
GOCMD = go
GOBUILD = $(GOCMD) build
GOTEST = $(GOCMD) test
BINARY_NAME = reBank

#BUN VARS


## BACKEND TASKS ##

# Build task
build:
	$(GOBUILD) -o $(BEPATH)/bin/$(BINARY_NAME) $(BEPATH)

# Test task
test:
	@$(GOTEST) -v $(BEPATH)

run:build
	@$(BEPATH)/bin/$(BINARY_NAME) --seed true

pg: 
	@docker rm -f post-db
	@docker run --name post-db -e POSTGRES_PASSWORD=1234 -p 5432:5432 -d postgres

## FRONT END TASKS ##
