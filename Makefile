test:
	go test -v -cover ./...
bench:
	go test -v ./_examples/performance_test.go