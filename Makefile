# Project-specific variables
BINARIES ?=	tvtime
OS=$(shell uname -s) 

# Common variables
SOURCES :=	$(shell find . -type f -name "*.go")
COMMANDS :=	$(shell go list ./... | grep -v /vendor/ | grep /cmd/)
PACKAGES :=	$(shell go list ./... | grep -v /vendor/ | grep -v /cmd/)
GO ?=		go

.PHONY: setup
setup:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh
ifeq ($(OS), Darwin)
	brew install dep
else
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
endif
	dep ensure -vendor-only

all:	build

.PHONY: build
build:	$(BINARIES)

.PHONY: install
install:
	$(GO) install ./cmd/gotty-client

$(BINARIES): $(SOURCES)
	$(GO) build -i -o $@ ./cmd/$@

.PHONY: lint
lint: 
	./bin/golangci-lint run --enable-all ./...

.PHONY: ci
ci: lint build

.PHONY: clean
clean:
	rm -f $(BINARIES)
