# Build the CLI binary
build:
	go build -o aoc-go ./cmd/aoc-go

# Install locally
install: build
	mv aoc-go $(GOPATH)/bin/

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -f aoc-go

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run

# Build for all platforms
build-all:
	GOOS=darwin GOARCH=amd64 go build -o dist/aoc-go_darwin_amd64 ./cmd/aoc-go
	GOOS=darwin GOARCH=arm64 go build -o dist/aoc-go_darwin_arm64 ./cmd/aoc-go
	GOOS=linux GOARCH=amd64 go build -o dist/aoc-go_linux_amd64 ./cmd/aoc-go
	GOOS=linux GOARCH=arm64 go build -o dist/aoc-go_linux_arm64 ./cmd/aoc-go
	GOOS=windows GOARCH=amd64 go build -o dist/aoc-go_windows_amd64.exe ./cmd/aoc-go

# Help target
help:
	@echo "aoc-go - Advent of Code CLI tool"
	@echo ""
	@echo "Available targets:"
	@echo "  make build        Build the CLI binary"
	@echo "  make install      Install the binary to GOPATH/bin"
	@echo "  make test         Run tests"
	@echo "  make clean        Remove build artifacts"
	@echo "  make fmt          Format code"
	@echo "  make lint         Run linter"
	@echo "  make build-all    Build for all platforms"
	@echo ""

.PHONY: build install test clean fmt lint build-all help
