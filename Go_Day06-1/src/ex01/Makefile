.PHONY:fmt vet build

all: build

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build && ./myArticles

clean:
	go clean