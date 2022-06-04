build:
	CGO_ENABLED=0 GOOS=linux go build -o pdf2text .

release: build
	docker buildx build --platform linux/amd64 -t haxii/pdf2text -f Dockerfile .

start:
	docker run -i --rm -p 5000:5000 --tmpfs /tmp haxii/pdf2text
