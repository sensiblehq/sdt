.PHONY: dev test

test:
	go test ./...
dev:
	go build -o bin/sdt ./cmd/sdt