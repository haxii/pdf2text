VERSION=$(shell git describe --abbrev=0 --tags)
BUILD=$(shell git rev-parse --short HEAD)

build:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o ./bin/pdf2text ./cmd/pdf2text-server/main.go
	upx -8 ./bin/pdf2text

dep:
	mkdir -p ./bin
clean:
	rm -rf ./bin

release: build
	docker buildx build --platform linux/amd64 -t haxii/pdf2text:$(VERSION) -f Dockerfile .

start:
	docker run -i --rm -p 5000:80 --tmpfs /tmp haxii/pdf2text
