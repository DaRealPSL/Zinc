# Binaries output directory
BIN_DIR = bin

# Compiler and tools
ZN = $(BIN_DIR)/zn
ZNFMT = $(BIN_DIR)/znfmt
ZNVET = $(BIN_DIR)/znvet
ZNDOC = $(BIN_DIR)/zndoc

# Default target
.PHONY: all
all: build

# Build all binaries
.PHONY: build
build: $(ZN) $(ZNFMT) $(ZNVET) $(ZNDOC)
	@echo "All tools built."

# Build compiler
$(ZN):
	@mkdir -p $(BIN_DIR)
	go build -o $(ZN) ./cmd/zn

# Build formatter
$(ZNFMT):
	@mkdir -p $(BIN_DIR)
	go build -o $(ZNFMT) ./cmd/znfmt

# Build linter
$(ZNVET):
	@mkdir -p $(BIN_DIR)
	go build -o $(ZNVET) ./cmd/znvet

# Build documentation generator
$(ZNDOC):
	@mkdir -p $(BIN_DIR)
	go build -o $(ZNDOC) ./cmd/zndoc

# Format source code
.PHONY: format
format:
	@echo "Formatting Zinc and Go code..."
	$(ZNFMT) fmt .
	go fmt ./...

# Lint source code
.PHONY: lint
lint:
	@echo "Linting Zinc and Go code..."
	go vet ./...
	$(ZNVET) .

# Run tests
.PHONY: test
test:
	@echo "Running Go tests..."
	go test ./internal/... ./pkg/...
	@echo "Running Zinc end-to-end tests..."
	$(ZN) test examples

# Run benchmarks
.PHONY: bench
bench:
	@echo "Running Go benchmarks..."
	go test -bench=. ./benchmarks/go
	@echo "Running Zinc benchmarks..."
	$(ZN) bench benchmarks/zn

# Bootstrap the compiler via script
.PHONY: bootstrap
bootstrap:
	@echo "Bootstrapping Zinc compiler..."
	scripts/bootstrap.sh

# Release pipeline
.PHONY: release
release:
	@echo "Releasing Zinc..."
	scripts/release.sh

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -rf $(BIN_DIR) dist
