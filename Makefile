.PHONY: help format format-go format-md check install

# Default target
help:
	@echo "Available commands:"
	@echo "  make format       - Format all Go and Markdown files"
	@echo "  make format-go    - Format only Go files"
	@echo "  make format-md    - Format only Markdown files"
	@echo "  make check        - Check formatting without changing files"
	@echo "  make install      - Install formatting tools"

# Install formatting tools
install:
	@echo "Installing formatting tools..."
	@which gofmt > /dev/null || echo "gofmt comes with Go installation"
	@which prettier > /dev/null || npm install -g prettier
	@echo "✓ Tools installed"

# Format all files
format: format-go format-md

# Format Go files
format-go:
	@echo "Formatting Go files..."
	@gofmt -s -w .
	@echo "✓ Go files formatted"

# Format Markdown files
format-md:
	@echo "Formatting Markdown files..."
	@prettier --write "**/*.md"
	@echo "✓ Markdown files formatted"

# Check formatting
check:
	@echo "Checking Go files..."
	@gofmt -l . | grep . && echo "❌ Some Go files need formatting" || echo "✓ Go files OK"
	@echo "Checking Markdown files..."
	@prettier --check "**/*.md" || echo "❌ Some Markdown files need formatting"

