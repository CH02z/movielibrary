# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=movielibrary
BINARY_UNIX=$(BINARY_NAME)_unix
 
all: build

build:
	$(GOBUILD) -o ./bin/$(BINARY_NAME) -v ./cmd/...
 
test:
	$(GOTEST) -v ./...

run:
	$(GOBUILD) -o ./bin/$(BINARY_NAME) -v ./cmd/...
	./bin/$(BINARY_NAME)
 
# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v