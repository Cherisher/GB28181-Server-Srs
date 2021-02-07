.PHONY: build clean tool lint help

all: build

build:
	@go build -tags='srs' -v -o gb28181-proxy-srs 

tool:
	go vet ./...; true
	gofmt -w .

lint:
	golint ./...

clean:
	rm -rf gb28181-proxy-srs
	go clean -i .

help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cached files"