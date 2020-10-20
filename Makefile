.PHONY: build test run wait clean

GO := go
BIN_WEBSERVER := cc
BUILD_PATH := ./cmd/clickcounter
ENVFLAGS = GO111MODULE=on CGO_ENABLED=0 GOOS=$(shell go env GOOS) GOARCH=$(shell go env GOARCH)

## usage: show available actions
usage: Makefile
	@echo  "to use make call:"
	@echo  "    make <action>"
	@echo  ""
	@echo  "list of available actions:"
	@if [ -x /usr/bin/column ]; \
	then \
		echo "$$(sed -n 's/^## /    /p' $< | column -t -s ':')"; \
	else \
		echo "$$(sed -n 's/^## /    /p' $<)"; \
	fi

## build: build server
build:
	@echo "==> Building binary (bin/$(BIN_WEBSERVER))..."
	$(ENVFLAGS) $(GO) build -v -o bin/$(BIN_WEBSERVER) $(BUILD_PATH)

## test: run unit tests
test:
	@echo  "==> Running tests with envs:"
	$(GO) test -v -race -cover

## run: run server
run:
	@echo "==> Running server with envs:"
	./bin/$(BIN_WEBSERVER) $(args)

## wait: used only to wait for database connections
wait:
	@echo "==> Waiting for other services to be ready..."
	bash scripts/tcp-port-wait.sh $(DATABASE_HOST) $(DATABASE_PORT)

## clean: clean local binaries
clean:
	@echo  "==> Running clean..."
	@rm -rf bin/$(BIN_WEBSERVER)
	@echo  "App clear! :)"