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

DARWIN_BINARY_NAME=$(BINARY_NAME)-darwin-amd64
LINUX_BINARY_NAME=$(BINARY_NAME)-linux-amd64
WINDOWS_BINARY_NAME=$(BINARY_NAME)-windows-amd64.exe

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

build-all-os:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(DARWIN_BINARY_NAME) -v $(MAIN_FILE)
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(LINUX_BINARY_NAME) -v $(MAIN_FILE)
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(WINDOWS_BINARY_NAME) -v $(MAIN_FILE)
