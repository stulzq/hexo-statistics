LDFLAGS := -s -w
VERSION := "v0.1.0"
APPNAME := "hexo-statistics"
BIN_NAME := "hexo-stat"

build:
	@env CGO_ENABLED=0 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/hexo-stat ./cmd

run:
	@go run ./cmd

build-linux:
	@GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/linux-amd64/$(BIN_NAME) ./cmd

build-darwin:
	@GOOS=darwin GOARCH=amd64 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/darwin-amd64/$(BIN_NAME) ./cmd

clean:
	@rm -rf bin/*

.PHONY: run build build-linux build-darwin
