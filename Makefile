# meta
NAME := password
VERSION := 0.1.0
REVISION := $(shell git rev-parse --short HEAD)

PKG := github.com/ikmski/password

SOURCES := $(shell find . -type f -name "*.go")
CMDSRCS := $(shell find ./cmd -type f -name "*.go")
GOFILES := $(shell find . -type f -name "*.go" -not -path "./vendor/*")
LDFLAGS := -X 'main.version=$(VERSION)' -X 'main.revision=$(REVISION)'

.PHONY: all
## all
all: build

.PHONY: setup
## setup
setup:
	go get -u github.com/golang/lint
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/Songmu/make2help/cmd/make2help

.PHONY: install-deps
## install dependencies
install-deps: setup
	dep ensure

.PHONY: update-deps
## update dependencies
update-deps: setup
	dep ensure -update

.PHONY: test
## run tests
test:
	go test -v

.PHONY: lint
## lint
lint:
	go vet
	for pkg in $(GOFILES); do\
		golint --set_exit_status $$pkg || exit $$?; \
	done

.PHONY: run
## run
run:
	go run $(CMDSRCS)

.PHONY: build
## build
build: bin/$(NAME)

bin/$(NAME): $(SOURCES)
	go build \
		-a -v \
		-tags netgo \
		-installsuffix netgo \
		-ldflags "$(LDFLAGS)" \
		-o $@ \
		${PKG}/cmd/${NAME}

.PHONY: install
## install
install:
	go install $(LDFLAGS)

.PHONY: clean
## clean
clean:
	go clean -i ./...
	rm -rf bin/*

.PHONY: help
## show help
help:
	@make2help $(MAKEFILE_LIST)

