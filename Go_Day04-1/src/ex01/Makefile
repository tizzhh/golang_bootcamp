.PHONY:fmt vet build

all: run

run: vet
	go run cmd/candy-server-server/main.go  --tls-port 3333 --tls-key candy.tld/key.pem --tls-certificate candy.tld/cert.pem --tls-ca minica.pem --host candy.tld

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

clean:
	go clean