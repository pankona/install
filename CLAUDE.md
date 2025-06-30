# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go CLI tool that installs development tools on Linux systems. The tool is designed for personal use by the author (pankona) to set up a development environment with specific tools and configurations.

## Architecture

The project follows a simple modular structure:
- `main.go`: Entry point with flag-based command selection for different tool installations
- Individual `.go` files for each tool installation (e.g., `asdf.go`, `docker.go`, `vim.go`)
- `util.go`: Contains shared utilities including `errContainer` for command execution and `withTempDir` for temporary directory operations
- `install.sh`: Bash script that downloads and runs the binary with predefined installation sequence

## Common Commands

### Build and Release
- `make dry-run`: Test release build without publishing
- `make release`: Create and publish release using goreleaser
- `make install-tools`: Install goreleaser dependency

### Development
- Standard Go commands apply: `go build`, `go run`, `go test`
- Binary is built as `install` in the root directory

### Usage
Each installation module is invoked via command-line flags:
- `./install -asdf`: Install asdf version manager
- `./install -apttools`: Install apt-based tools
- `./install -bashrc`: Install custom bashrc configuration
- `./install -golang`: Install Go toolchain
- And many others (see main.go for complete list)

## Key Patterns

### Installation Functions
Each tool has its own installation function following the pattern `install<ToolName>(ctx context.Context)`. These functions:
- Check if tool is already installed to avoid duplicates
- Use `errContainer` to chain command executions with error handling
- Often use `withTempDir` for temporary file operations

### Error Handling
The `errContainer` pattern allows chaining multiple shell commands while preserving the first error encountered.

### Shell Integration
The project installs a custom bashrc (`mybashrc`) that integrates with asdf and other tools.

## Code Style Guidelines

### Comments
- Comments should explain **why** or **why not**, not **what**
- Do not comment on code that is self-explanatory
- Comments should provide context that cannot be understood from reading the code alone
- Focus on business logic, design decisions, or non-obvious technical choices