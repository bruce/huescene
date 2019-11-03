SHELL := /bin/bash

export GO111MODULE=on

.PHONY: build
build: huescene

huescene:
	CGO_ENABLED=0 go build ./cmd/huescene

.PHONY: clean
clean:
	$(RM) huescene
