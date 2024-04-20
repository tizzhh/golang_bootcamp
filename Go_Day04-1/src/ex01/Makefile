.PHONY:fmt vet build

all: run

run: vet
	go run cmd/candy-server-server/main.go  --tls-port 3333 --tls-key cert/key.pem --tls-certificate cert/cert.pem --tls-ca cert/minica.pem --host candy.tld

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

clean:
	go clean