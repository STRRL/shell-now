# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Shell Now is a Go CLI tool that creates instant public web terminals using ttyd and Cloudflare Quick Tunnels. The tool automatically downloads required dependencies (ttyd and cloudflared) and sets up a secure web-based shell accessible via a public URL.

## Build Commands

- `make build` or `go build -o out/shell-now ./cmd/shell-now/` - Build the binary
- `DEBUG=1 ./out/shell-now` - Run with debug logging enabled
- `docker build -t shell-now .` - Build Docker image
- `docker run cr.strrl.dev/strrl/shell-now:latest` - Run via Docker

## Architecture

The codebase follows a simple CLI pattern with platform-specific implementations:

### Core Components
- `cmd/shell-now/main.go` - CLI entry point using Cobra, handles signals and context
- `pkg/bootstrap.go` - Main orchestration logic that coordinates ttyd and cloudflared processes
- Platform-specific preparation files (`prepare_linux.go`, `prepare_darwin.go`, `prepare.go`) that handle dependency installation

### Key Architecture Points
- Uses build tags for platform-specific code (`//go:build !darwin && !linux`)
- Linux/ARM64 support with automatic binary downloads to `~/.local/bin/`
- macOS requires manual installation via Homebrew
- Concurrent goroutines for ttyd and cloudflared with WaitGroup synchronization
- Context-based cancellation for graceful shutdown on SIGINT/SIGTERM
- Random port allocation for ttyd using `net.Listen("tcp", ":0")`
- Random 6-digit password generation for web terminal authentication

### Dependencies
- Uses `github.com/spf13/cobra` for CLI framework
- Downloads ttyd v1.7.7 and cloudflared 2025.5.0 on Linux
- Expects `ttyd` and `cloudflared` via Homebrew on macOS

## Platform Considerations

The tool has different behavior per platform:
- **Linux**: Automatically downloads binaries to `~/.local/bin/` if not in PATH
- **macOS**: Requires manual installation via `brew install ttyd cloudflared`
- **Other platforms**: Requires manual installation with informational messages only