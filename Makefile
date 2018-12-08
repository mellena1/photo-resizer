# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOFMT=$(GOCMD) fmt
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

MAIN_FILE=cmd/photoresizer.go
BINARY_NAME=photo-resizer

all: test build

build: 
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN_FILE)

test: 
	$(GOTEST) -v ./...

fmt:
	$(GOFMT) ./...

clean: 
	$(GOCLEAN) $(MAIN_FILE)
	rm -f $(BINARY_NAME)

get:
	$(GOGET) ./...
