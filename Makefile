.PHONY: test bench benchfile fmt

all: test bench

test:
	go test -cover ./...

bench:
	go test -v ./_examples/performance_test.go -benchtime=1s

fmt:
	go fmt ./...