.PHONY: test build lint

test:
	go test ./internal

test-verbose:
	go test ./internal -v | grep -v "^.*RUN"

lint:
	golangci-lint run

build:
	go build
