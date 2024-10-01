format:
	@gofmt -s -w .

run: dependencies format
	@go run cmd/server/main.go

build: dependencies format
	@go build cmd//server/main.go

dependencies:
	@go mod tidy
