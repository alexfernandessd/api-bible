.PHONY: usage build build-webserver test run-webserver

OK_COLOR=\033[32;01m
NO_COLOR=\033[0m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

GO := go
GO_LINTER := golint
DOCKER := docker
DOCKER_COMPOSE := docker-compose
BUILDOS ?= $(shell go env GOHOSTOS)
BUILDARCH ?= amd64
ECHOFLAGS ?=
ROOT_DIR := $(realpath .)

PKGS = $(shell $(GO) list ./...)

ENVFLAGS ?= CGO_ENABLED=0
BUILDENV ?= GOOS=$(BUILDOS) GOARCH=$(BUILDARCH)

BIN_WEBSERVER := api-bible

CREATE_LOCAL_ENV := $(shell if [ ! -f "$(ROOT_DIR)/.env" ]; then cp $(ROOT_DIR)/.env.example $(ROOT_DIR)/.env; fi)
LOCAL_VARIABLES ?= $(shell while read -r line; do printf "$$line" | sed 's/ /\\ /g' | awk '{print}'; done < $(ROOT_DIR)/.env)

## test: run unit tests
test:
	@echo $(ECHOFLAGS) "$(OK_COLOR)==> Running tests with envs:[$(LOCAL_VARIABLES)]...$(NO_COLOR)"
	@$(LOCAL_VARIABLES) $(ENVFLAGS) $(GO) test $(GOFLAGS) $(PKGS)

## run: run
run:
	@echo $(ECHOFLAGS) "$(OK_COLOR)==> Running webserver with envs:[$(LOCAL_VARIABLES)]...$(NO_COLOR)"
	@$(LOCAL_VARIABLES) bin/$(BUILDOS)_$(BUILDARCH)/$(BIN_WEBSERVER) $(args)

## build: build
build:
	@echo $(ECHOFLAGS) "$(OK_COLOR)==> Building binary ($(BUILDOS)/$(BUILDARCH)/$(BIN_WEBSERVER))...$(NO_COLOR)"
	@echo $(ECHOFLAGS) $(ENVFLAGS) $(BUILDENV) $(GO) build -v $(BUILDFLAGS) -o bin/$(BUILDOS)_$(BUILDARCH)/$(BIN_WEBSERVER) ./cmd
	@$(ENVFLAGS) $(BUILDENV) $(GO) build -v -o bin/$(BUILDOS)_$(BUILDARCH)/$(BIN_WEBSERVER) ./cmd/api-bible
