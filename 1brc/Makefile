# Define the binary name
BINARY_NAME=1brc

# Define the main package
MAIN_PACKAGE=.

# Define the Go source files
SRC=$(shell find . -type f -name '*.go')

# Default target
all: build

# Build the binary
build: $(SRC)
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME) $(MAIN_PACKAGE)

# Clean up
clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod tidy

# Install on system in $GOBIN
install: build
	@echo "Installing on system in $(go env GOBIN)..."
	go install

# Run the application
run: build
	@echo "Running $(BINARY_NAME)..."
	@./$(go env GOBIN)/$(BINARY_NAME)

.PHONY: all build clean deps run