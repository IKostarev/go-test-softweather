.PHONY: build run clean

build:
	go build -o app ./cmd/substring/main.go

docker-build:
	docker build -t app .

run:
	go run ./cmd/substring/main.go

docker-run:
	docker run -p 8080:8080 app

clean:
	rm -f app
