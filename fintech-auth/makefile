build:
	go build -o bin/main cmd/api/main.go

run:
	go run cmd/api/main.go

build_image: 
	go build -o bin/main cmd/api/main.go
	docker build -t fintech-auth .
