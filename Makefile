# sourced by https://github.com/octomation/makefiles

.DEFAULT_GOAL = test-with-coverage
GIT_HOOKS     = post-merge pre-commit pre-push
GO_VERSIONS   = 1.15
GO111MODULE   = on

AT    := @
ARCH  := $(shell uname -m | tr '[:upper:]' '[:lower:]')
OS    := $(shell uname -s | tr '[:upper:]' '[:lower:]')
DATE  := $(shell date +%Y-%m-%dT%T%Z)
SHELL := /usr/bin/env bash -euo pipefail -c

make-verbose:
	$(eval AT :=)
	@echo > /dev/null
.PHONY: make-verbose

todo:
	@grep \
		--exclude=Makefile \
		--exclude-dir={bin,components,node_modules,vendor} \
		--color \
		--text \
		-nRo -E ' TODO:.*|SkipNow' . || true
.PHONY: todo

rmdir:
	$(AT) for dir in `git ls-files --others --exclude-standard --directory`; do \
		find $${dir%%/} -depth -type d -empty | xargs rmdir; \
	done
.PHONY: rmdir

COMMIT  := $(shell git rev-parse --verify HEAD)
RELEASE := $(shell git describe --tags 2>/dev/null | rev | cut -d - -f3- | rev)

ifdef GIT_HOOKS

hooks: unhook
	$(AT) for hook in $(GIT_HOOKS); do cp githooks/$$hook .git/hooks/; done
.PHONY: hooks

unhook:
	@ls .git/hooks | grep -v .sample | sed 's|.*|.git/hooks/&|' | xargs rm -f || true
.PHONY: unhook

define hook_tpl
$(1):
	@githooks/$(1)
.PHONY: $(1)
endef

render_hook_tpl = $(eval $(call hook_tpl,$(hook)))
$(foreach hook,$(GIT_HOOKS),$(render_hook_tpl))

endif

git-check:
	$(AT) git diff --exit-code >/dev/null
	$(AT) git diff --cached --exit-code >/dev/null
	$(AT) ! git ls-files --others --exclude-standard | grep -q ^
.PHONY: git-check

GOBIN       ?= $(PWD)/bin/$(OS)/$(ARCH)
GOFLAGS     ?= -mod=
GOPRIVATE   ?= go.octolab.net
GOPROXY     ?= direct
GOTEST      ?= $(GOBIN)/testit
GOTESTFLAGS ?=
GOTRACEBACK ?= all
LOCAL       ?= $(MODULE)
MODULE      ?= `go list -m $(GOFLAGS)`
PACKAGES    ?= `go list $(GOFLAGS) ./...`
PATHS       ?= $(shell echo $(PACKAGES) | sed -e "s|$(MODULE)/||g" | sed -e "s|$(MODULE)|$(PWD)/*.go|g")
TIMEOUT     ?= 1s

ifeq (, $(wildcard $(GOTEST)))
	GOTEST = $(shell command -v testit)
endif
ifeq (, $(GOTEST))
	GOTEST = go test
else
	GOTEST := $(GOTEST) go --colored
endif

ifeq (, $(PACKAGES))
	PACKAGES = $(MODULE)
endif

ifeq (, $(PATHS))
	PATHS = .
endif

export GOBIN       := $(GOBIN)
export GOFLAGS     := $(GOFLAGS)
export GOPRIVATE   := $(GOPRIVATE)
export GOPROXY     := $(GOPROXY)
export GOTRACEBACK := $(GOTRACEBACK)

go-env:
	@echo "GO111MODULE: $(strip `go env GO111MODULE`)"
	@echo "GOBIN:       $(strip `go env GOBIN`)"
	@echo "GOFLAGS:     $(strip `go env GOFLAGS`)"
	@echo "GOPRIVATE:   $(strip `go env GOPRIVATE`)"
	@echo "GOPROXY:     $(strip `go env GOPROXY`)"
	@echo "GOTEST:      $(GOTEST)"
	@echo "GOTESTFLAGS: $(GOTESTFLAGS)"
	@echo "GOTRACEBACK: $(GOTRACEBACK)"
	@echo "LOCAL:       $(LOCAL)"
	@echo "MODULE:      $(MODULE)"
	@echo "PACKAGES:    $(PACKAGES)"
	@echo "PATHS:       $(strip $(PATHS))"
	@echo "TIMEOUT:     $(TIMEOUT)"
.PHONY: go-env

go-verbose:
	$(eval GOTESTFLAGS := -v)
	@echo > /dev/null
.PHONY: go-verbose

deps-check:
	@go mod verify
	@if command -v egg > /dev/null; then \
		egg deps check license; \
		egg deps check version; \
	fi
.PHONY: deps-check

deps-clean:
	@go clean -modcache
.PHONY: deps-clean

deps-fetch:
	@go mod download
	@if [[ "`go env GOFLAGS`" =~ -mod=vendor ]]; then go mod vendor; fi
.PHONY: deps-fetch

deps-tidy:
	@go mod tidy
	@if [[ "`go env GOFLAGS`" =~ -mod=vendor ]]; then go mod vendor; fi
.PHONY: deps-tidy

deps-update: selector = '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}'
deps-update:
	$(AT) if command -v egg > /dev/null; then \
		packages="`egg deps list | tr ' ' '\n' | sed -e 's|$$|/...@latest|'`"; \
	else \
		packages="`go list -f $(selector) -m -mod=readonly all | sed -e 's|$$|/...@latest|'`"; \
	fi; \
	if [[ "$$packages" = "/...@latest" ]]; then exit; fi; \
	for package in $$packages; do go get $$package; done

	$(AT) if [ -z "$(AT)" ]; then MAKE="$(MAKE) verbose"; else MAKE="$(MAKE)"; fi; \
	$$MAKE deps-tidy
.PHONY: deps-update

GODOC_HOST ?= localhost:6060

go-docs:
	@(sleep 2 && open http://$(GODOC_HOST)/pkg/$(LOCAL)/) &
	@godoc -http=$(GODOC_HOST)
.PHONY: go-docs

go-fmt:
	@if command -v goimports > /dev/null; then \
		goimports -local $(LOCAL) -ungroup -w $(PATHS); \
	else \
		gofmt -s -w $(PATHS); \
	fi
.PHONY: go-fmt

go-generate:
	@go generate $(PACKAGES)
.PHONY: go-generate

go-pkg:
	@open https://pkg.go.dev/$(MODULE)@$(RELEASE)
.PHONY: go-pkg

lint:
	@golangci-lint run ./...
	@looppointer ./...
.PHONY: lint

test:
	@$(GOTEST) -race -timeout $(TIMEOUT) $(GOTESTFLAGS) $(PACKAGES)
.PHONY: test

test-clean:
	@go clean -testcache
.PHONY: test-clean

test-quick:
	@$(GOTEST) -timeout $(TIMEOUT) $(GOTESTFLAGS) $(PACKAGES)
.PHONY: test-quick

test-with-coverage:
	@$(GOTEST) \
		-cover \
		-covermode atomic \
		-coverprofile c.out \
		-race \
		-timeout $(TIMEOUT) \
		$(GOTESTFLAGS) \
		$(PACKAGES)
.PHONY: test-with-coverage

test-with-coverage-report: test-with-coverage
	@go tool cover -html c.out
.PHONY: test-with-coverage-report

test-integration: GOTAGS = integration
test-integration:
	@$(GOTEST) \
		-cover \
		-covermode atomic \
		-coverprofile integration.out \
		-race \
		-tags $(GOTAGS) \
		$(GOTESTFLAGS) \
		./...
.PHONY: test-integration

test-integration-quick: GOTAGS = integration
test-integration-quick:
	@$(GOTEST) -tags $(GOTAGS) $(GOTESTFLAGS) ./...
.PHONY: test-integration-quick

test-integration-report: test-integration
	@go tool cover -html integration.out
.PHONY: test-integration-report

BINARY  ?= $(GOBIN)/$(shell basename $(MAIN))
LDFLAGS ?= -ldflags "-s -w -X main.commit=$(COMMIT) -X main.date=$(DATE)"
MAIN    ?= $(MODULE)

build-env:
	@echo "DATE:        $(DATE)"
	@echo "COMMIT:      $(COMMIT)"
	@echo "RELEASE:     $(RELEASE)"
	@echo "BINARY:      $(BINARY)"
	@echo "LDFLAGS:     $(LDFLAGS)"
	@echo "MAIN:        $(MAIN)"
.PHONY: build-env

build:
	@go build -o $(BINARY) $(LDFLAGS) $(MAIN)
.PHONY: build

build-with-race:
	@go build -race -o $(BINARY) $(LDFLAGS) $(MAIN)
.PHONY: build-with-race

build-clean:
	@rm -f $(BINARY)
.PHONY: build-clean

install:
	@go install $(LDFLAGS) $(MAIN)
.PHONY: install

install-clean:
	@go clean -cache
.PHONY: install-clean

dist-check:
	@goreleaser --snapshot --skip-publish --rm-dist
.PHONY: dist-check

dist-dump:
	@godownloader .goreleaser.yml > bin/install
.PHONY: dist-dump

TOOLFLAGS ?= -mod=

tools-env:
	@echo "GOBIN:       `go env GOBIN`"
	@echo "TOOLFLAGS:   $(TOOLFLAGS)"
.PHONY: tools-env

tools-fetch: GOFLAGS = $(TOOLFLAGS)
tools-fetch:
	$(AT) cd tools; \
	go mod download; \
	if [[ "`go env GOFLAGS`" =~ -mod=vendor ]]; then go mod vendor; fi
.PHONY: tools-fetch

tools-tidy: GOFLAGS = $(TOOLFLAGS)
tools-tidy:
	$(AT) cd tools; \
	go mod tidy; \
	if [[ "`go env GOFLAGS`" =~ -mod=vendor ]]; then go mod vendor; fi
.PHONY: tools-tidy

tools-install: GOFLAGS = $(TOOLFLAGS)
tools-install: GOTAGS = tools
tools-install: tools-fetch
	$(AT) cd tools; \
	go generate -tags $(GOTAGS) tools.go
.PHONY: tools-install

tools-update: GOFLAGS = $(TOOLFLAGS)
tools-update: selector = '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}'
tools-update:
	$(AT) cd tools; \
	if command -v egg > /dev/null; then \
		packages="`egg deps list | tr ' ' '\n' | sed -e 's|$$|/...@latest|'`"; \
	else \
		packages="`go list -f $(selector) -m -mod=readonly all | sed -e 's|$$|/...@latest|'`"; \
	fi; \
	if [[ "$$packages" = "/...@latest" ]]; then exit; fi; \
	for package in $$packages; do go get $$package; done

	$(AT) if [ -z "$(AT)" ]; then MAKE="$(MAKE) verbose"; else MAKE="$(MAKE)"; fi; \
	$$MAKE tools-tidy tools-install
.PHONY: tools-update

ifdef GO_VERSIONS

define go_tpl
go$(1):
	@docker run \
		--rm -it \
		-v $(PWD):/src \
		-w /src \
		golang:$(1) bash
.PHONY: go$(1)
endef

render_go_tpl = $(eval $(call go_tpl,$(version)))
$(foreach version,$(GO_VERSIONS),$(render_go_tpl))

endif


export PATH := $(GOBIN):$(PATH)

init: deps test lint hooks
	@git config core.autocrlf input
.PHONY: init

clean: build-clean deps-clean install-clean test-clean
.PHONY: clean

deps: deps-fetch tools-install
.PHONY: deps

env: go-env build-env tools-env
env:
	@echo "PATH:        $(PATH)"
.PHONY: env

format: go-fmt
.PHONY: format

generate: go-generate format
.PHONY: generate

refresh: deps-tidy update deps generate test build
.PHONY: refresh

update: deps-update tools-update
.PHONY: update

verbose: make-verbose go-verbose
.PHONY: verbose

verify: deps-check generate test lint git-check
.PHONY: verify
