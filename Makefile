.PHONY: test bench fmt

all: test bench

test:
	go test -cover ./...

bench:
	go test -v -bench=. -benchtime=1s ./_examples/performance_test.go

fmt:
	go fmt ./...