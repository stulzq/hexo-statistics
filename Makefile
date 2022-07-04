LDFLAGS := -s -w

build: build-linux-amd64 build-darwin-amd64

run:
	@go run ./cmd

build-linux-amd64:
	@GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/linux-amd64/hexo-stat ./cmd
	@cp conf/config.yml bin/linux-amd64

build-darwin-amd64:
	@GOOS=darwin GOARCH=amd64 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/darwin-amd64/hexo-stat ./cmd
	@cp conf/config.yml bin/darwin-amd64

clean:
	@rm -f /bin/hexo-stat

.PHONY: run build
