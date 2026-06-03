BINARY := nona
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)

.PHONY: build test clean

build:
	go build -ldflags "-X main.version=$(VERSION)" -o $(BINARY) ./cmd/nona

test:
	go test ./...

clean:
	rm -f $(BINARY)
