.PHONY: build help
PWD := $(shell pwd)
GOPATH := $(shell go env GOPATH)

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

VERSION ?= $(shell git describe --tags)

all: build

help: ## print this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' Makefile | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-40s\033[0m %s\n", $$1, $$2}'

build: ## build
	@CGO_ENABLED=1 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(PWD)/build/kopi

