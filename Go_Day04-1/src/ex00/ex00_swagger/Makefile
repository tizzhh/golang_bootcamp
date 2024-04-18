.PHONY:fmt vet build

all: run

run: vet
	go run cmd/candy-server-server/main.go --port 3333

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

clean:
	go clean