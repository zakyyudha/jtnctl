# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get
BINARY_NAME = jtnctl

# This should be set to the directory containing your Go source files
SRC_DIR = ./main.go

# Build for all supported platforms
all: clean build-linux build-windows build-darwin

# Clean the workspace
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Build for Linux
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) $(SRC_DIR)

# Build for Windows
build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME).exe $(SRC_DIR)

# Build for macOS (Darwin)
build-darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)_darwin $(SRC_DIR)

# Install dependencies (if needed)
get:
	$(GOGET) ./...

# Run tests (if you have tests)
test:
	$(GOTEST) -v ./...

# Run the built binary (example)
run:
	./$(BINARY_NAME)

.PHONY: clean build-linux build-windows build-darwin get test run
