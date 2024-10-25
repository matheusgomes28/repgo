# Go parameters
GOCMD=go
BUILD_DIR=./build
REPGO_DIR=./cmd/repgo

# Name of the binary
BINARY_NAME=repgo

all: build test

build:
	GIN_MODE=release $(GOCMD) build -ldflags "-s" -v -o $(BUILD_DIR)/$(BINARY_NAME) $(REPGO_DIR)

test:
	$(GOCMD) test -v ./...

clean:
	$(GOCMD) clean
	rm -rf $(BUILD_DIR)

.PHONY: all build test clean
