ARCH  := $(shell uname -m | tr '[:upper:]' '[:lower:]')
OS    := $(shell uname -s | tr '[:upper:]' '[:lower:]')
SHELL  = /usr/local/bin/bash -euo pipefail

export GOBIN := $(abspath $(PWD)/../..)/bin/$(OS)/$(ARCH)
export PATH  := $(GOBIN):$(PATH)

GOTESTFLAGS ?=

verbose:
	$(eval GOTESTFLAGS := -v)
	@echo > /dev/null
.PHONY: verbose

test:
	@go test $(GOTESTFLAGS) ./...
.PHONY: test

gotest:
	@gotest $(GOTESTFLAGS) ./...
.PHONY: gotest

panicparse:
	@go test $(GOTESTFLAGS) ./... |&panicparse
.PHONY: panicparse

gotest-fail-with-panicparse:
	@gotest $(GOTESTFLAGS) ./... |&panicparse
.PHONY: gotest-fail-with-panicparse

gotest-with-panicparse-fail:
	@script -q /dev/null gotest $(GOTESTFLAGS) ./... |&panicparse
.PHONY: gotest-with-panicparse-fail

testit-with-panicparse-fail:
	@script -q /dev/null testit go --colored $(GOTESTFLAGS) ./... |&panicparse
.PHONY: testit-with-panicparse-fail

testit:
	@testit go --colored --stacked $(GOTESTFLAGS) ./...
.PHONY: testit
