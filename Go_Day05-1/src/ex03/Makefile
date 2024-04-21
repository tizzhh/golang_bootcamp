.PHONY:fmt vet build

all: build

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build

test: vet
	go test -v knapsack/knapsack_test.go

clean:
	go clean