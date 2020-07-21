    # Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=$(GOCMD)fmt
MAIN= cmd/main.go
BINARY_NAME=fuego

all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN)

test:
	$(GOTEST) -v ./...

run:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN)
	./$(BINARY_NAME)

fmt:
	$(GOFMT) -w .

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v