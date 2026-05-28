BINARY := nona

.PHONY: build test clean

build:
	go build -o $(BINARY) ./cmd/nona

test:
	go test ./...

clean:
	rm -f $(BINARY)
