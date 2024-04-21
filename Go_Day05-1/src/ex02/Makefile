.PHONY:fmt vet build

all: build

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build

test: vet
	go test -v present_heap/present_heap_test.go

clean:
	go clean