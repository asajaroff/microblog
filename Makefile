run: dependencies
	go run cmd/main.go

build: dependencies
	go build cmd/main.go

dependencies:
	go mod tidy
