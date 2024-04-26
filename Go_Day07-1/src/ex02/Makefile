.PHONY:fmt vet build

all: build

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build && ./logoGen

doc:
	godoc -http=localhost:8080 & wget -r -np -N -E -p -k http://localhost:8080/pkg/findMinCoins/

clean:
	go clean & rm -rf localhost\:8080/