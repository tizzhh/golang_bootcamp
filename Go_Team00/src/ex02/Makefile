.PHONY:fmt vet build

all: build

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

test: vet
	go test -v ./...

build: vet
	go build && ./transmitterGRPCWriteToDB

clean:
	go clean