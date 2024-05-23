# Go parameters
GOCMD=go
TEMPL=templ
BUILD_DIR=./tmp
GOCMS_DIR=./cmd/cmsgo

# Name of the binary
BINARY_NAME=cmsgo

all: build test

build:
	$(TEMPL) generate
	$(GOCMD) build -v -o $(BUILD_DIR)/$(BINARY_NAME) $(GOCMS_DIR)

test:
	$(GOCMD) test -v ./...

clean:
	$(GOCMD) clean
	rm -f $(BUILD_DIR)/$(BINARY_NAME)

run:
	$(GOCMD) run -o $(BINARY_NAME) -v ./...
	./$(BUILD_DIR)/$(BINARY_NAME)

.PHONY: all build test clean run


