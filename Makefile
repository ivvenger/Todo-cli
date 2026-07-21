BINARY := todo

.PHONY: build run test lint fmt vet clean

build:
	go build -o $(BINARY) .

run:
	go run . $(ARGS)

test:
	go test ./... -v

cover:
	go test ./... -cover

lint:
	golangci-lint run

fmt:
	go fmt ./...

vet:
	go vet ./...

clean:
	rm -f $(BINARY)