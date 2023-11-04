.PHONY: build
build:
	go build -o url-shortener src/cmd/app/main.go

.PHONY: test
test:
	go test -v -timeout 30s ./...

.DEFAULT_GOAL := build