GO_BIN := $(shell which go)
CMD_PARAMS := "CGO_ENABLED=0 GOOS=linux"
CGO_ENABLED ?= 0
GOOS = linux

format:
	@gofmt -s -w .

run: dependencies format
	@go run cmd/server/main.go

build: dependencies format
	@echo Params:
	@echo ${CMD_PARAMS} ${CGO_ENABLED}
	${GO_BIN} build -o microblog-api cmd/server/main.go

dependencies:
	@go mod tidy
