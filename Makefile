LDFLAGS := -s -w

build: build-linux-amd64 build-darwin-amd64

run:
	@go run ./cmd

build-linux-amd64:
	@GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/linux-amd64/hexo-stat ./cmd
	@mkdir bin/linux-amd64/conf/
	@cp conf/config.yml bin/linux-amd64/conf/

build-darwin-amd64:
	@GOOS=darwin GOARCH=amd64 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/darwin-amd64/hexo-stat ./cmd
	@mkdir bin/darwin-amd64/conf/
	@cp conf/config.yml bin/darwin-amd64/conf/

clean:
	@rm -rf bin/*

.PHONY: run build
