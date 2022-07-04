LDFLAGS := -s -w
VERSION := "v0.1.0"
APPNAME := "hexo-statistics"

build: clean build-linux-amd64 build-darwin-amd64 build-windows-amd64

run:
	@go run ./cmd

build-linux-amd64:
	@rm -rf bin/output
	@GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/output/hexo-stat ./cmd
	@mkdir bin/output/conf/
	@cp conf/config.yml bin/output/conf/
	@cd bin && tar -zvcf $(APPNAME)-$(VERSION)-linux-amd64.tar.gz output
	@echo "linux-amd64 build and package"

build-darwin-amd64:
	@rm -rf bin/output
	@GOOS=darwin GOARCH=amd64 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/output/hexo-stat ./cmd
	@mkdir bin/output/conf/
	@cp conf/config.yml bin/output/conf/
	@cd bin && tar -zvcf $(APPNAME)-$(VERSION)-darwin-amd64.tar.gz output
	@echo "darwin-amd64 build and package"

build-windows-amd64:
	@rm -rf bin/output
	@GOOS=windows GOARCH=amd64 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/output/hexo-stat ./cmd
	@mkdir bin/output/conf/
	@cp conf/config.yml bin/output/conf/
	@cd bin && tar -zvcf $(APPNAME)-$(VERSION)-windows-amd64.tar.gz output
	@echo "windows-amd64 build and package"

docker:
	@rm -rf bin/$(APPNAME)-docker
	@GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/$(APPNAME)-docker/hexo-stat ./cmd
	@mkdir bin/$(APPNAME)-docker/conf
	@cp conf/config.yml bin/$(APPNAME)-docker/conf/
	@docker build -t stulzq/hexo-statistics:$(VERSION) .

clean:
	@rm -rf bin/*

.PHONY: run build
