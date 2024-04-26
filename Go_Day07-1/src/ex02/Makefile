.PHONY:fmt vet build

all: build

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build && ./logoGen

doc:
	godoc -http=localhost:8080 & wget -r -np -N -E -p -k http://localhost:8080/pkg/findMinCoins/ || true && mv localhost\:8080/ docs/ && zip -r docs.zip docs/

clean:
	go clean & rm -rf docs/ docs.zip