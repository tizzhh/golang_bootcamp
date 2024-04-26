.PHONY:fmt vet build

all: build

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build && ./logoGen

test: vet
	go test -v

clean:
	go clean