all: test

build:
	@go build

test: build
	@go test
	@go generate ./examples
	@go test ./examples

.PHONY: build, test
