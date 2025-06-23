.PHONY: build
build:
	go build -ldflags "-X main.version=$(shell git describe --tags --always --dirty) -X main.commit=$(shell git rev-parse HEAD) -X main.buildTime=$(shell date -u +%Y-%m-%dT%H:%M:%SZ)" -o out/shell-now ./cmd/shell-now/
