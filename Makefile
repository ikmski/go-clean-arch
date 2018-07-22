# meta
NAME := go-clean-arch
VERSION := 0.1.0
REVISION := $(shell git rev-parse --short HEAD)
DOCKER_IMAGE := ${NAME}

SOURCES := $(shell find . -type f -name "*.go")
GOFILES := $(shell find . -type f -name "*.go" -not -name '*_test.go' -not -path "./vendor/*")

LDFLAGS := -s -w -X 'main.version=$(VERSION)' -X 'main.revision=$(REVISION)'

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
	go run $(GOFILES)

.PHONY: build
## build
build: bin/$(NAME)

bin/$(NAME): $(SOURCES)
	go build \
		-a -v \
		-tags netgo \
		-installsuffix netgo \
		-ldflags "$(LDFLAGS)" \
		-o $@

.PHONY: install
## install
install: install-deps
	go install \
		-a -v \
		-tags netgo \
		-installsuffix netgo \
		-ldflags "$(LDFLAGS)" \
		.

.PHONY: clean
## clean
clean:
	go clean -i ./...
	rm -rf bin/*

.PHONY: docker-build
## docker build
docker-build:
	docker build --rm -t ${DOCKER_IMAGE} .

.PHONY: docker-run
## docker run
docker-run:
	docker run --rm ${DOCKER_IMAGE}

.PHONY: help
## show help
help:
	@make2help $(MAKEFILE_LIST)

