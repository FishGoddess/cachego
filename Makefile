.PHONY: test bench fmt

all: test bench

test:
	go test -cover ./...

bench:
	go test -v ./_examples/performance_test.go

fmt:
	go fmt ./...