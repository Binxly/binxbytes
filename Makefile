# Makefile for building Go Lambda binary for AWS SAM

.PHONY: build clean help

help:
	@echo "Available commands:"
	@echo "  make run    - Run dev server locally"
	@echo "  make build  - Build the Go binary for Lambda"
	@echo "  make clean  - Remove the built binary"

run:
	go run main.go -dev

build:
	GOOS=linux GOARCH=arm64 go build -o bootstrap

clean:
	rm -f bootstrap
