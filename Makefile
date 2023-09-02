APP_BINARY := ./bin/app
GO_SOURCES := $(shell find . -iname "*.go")

.DEFAULT_GOAL := build
.PHONY: css build run clean

css:
	@npm run css

build: $(GO_SOURCES)
	@go build -o $(APP_BINARY) ./cmd/app

run: build
	@$(APP_BINARY)

clean:
	@rm -rf $(APP_BINARY)
