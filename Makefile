# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=eifel
BINARY_UNIX=$(BINARY_NAME)_unix

# All target
all: test build

# Build the project
build:
	$(GOBUILD) -o $(BINARY_NAME) ./cmd/$(BINARY_NAME)

# Run tests
test:
	$(GOTEST) -v ./...

# Clean build files
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

# Install dependencies
deps:
	$(GOGET) -u ./...

# Format the code
fmt:
	$(GOCMD) fmt ./...

# Run the project
run:
	$(GOBUILD) -o $(BINARY_NAME) ./cmd/$(BINARY_NAME)
	./$(BINARY_NAME) $(ARGS)

# Cross compile for Linux
build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) ./cmd/$(BINARY_NAME)