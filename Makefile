.PHONY: build test vet clean run

BINARY=fakerkit

build:
	CGO_ENABLED=0 go build -o $(BINARY) ./cmd/fakerkit

test:
	go test ./... -v

vet:
	go vet ./...

run: build
	./$(BINARY)

clean:
	rm -f $(BINARY)
