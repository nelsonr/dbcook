.PHONY: test build lint clean

test:
	go test ./internal

test-verbose:
	go test ./internal -v | grep -v "^.*RUN"

build:
	go build

install:
	go install

lint:
	golangci-lint run

clean:
	rm *.sql *.db
