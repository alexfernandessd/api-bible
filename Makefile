.PHONY: usage build build-webserver test run-webserver

OK_COLOR=\033[32;01m
NO_COLOR=\033[0m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

# Build Flags
BUILD_DATE = $(shell date -u --rfc-3339=seconds)
BUILD_HASH ?= $(shell git rev-parse --short HEAD)
APP_VERSION ?= undefined
BUILD_NUMBER ?= dev

NOW = $(shell date -u '+%Y%m%d%I%M%S')

GO := go
GO_LINTER := golint
DOCKER := docker
DOCKER_COMPOSE := docker-compose
BUILDOS ?= $(shell go env GOHOSTOS)
BUILDARCH ?= amd64
GOFLAGS ?=
ECHOFLAGS ?=
ROOT_DIR := $(realpath .)

ENVFLAGS ?= CGO_ENABLED=0
BUILDENV ?= GOOS=$(BUILDOS) GOARCH=$(BUILDARCH)
EXTLDFLAGS ?= -extldflags "-lm -lstdc++ -static"

BIN_WEBSERVER := api-bible

CREATE_LOCAL_ENV := $(shell if [ ! -f "$(ROOT_DIR)/.env" ]; then cp $(ROOT_DIR)/.env.example $(ROOT_DIR)/.env; fi)
LOCAL_VARIABLES ?= $(shell for i in $(shell cat $(ROOT_DIR)/.env); do echo -n "$$i "; done)

## run: run
run:
	@echo $(ECHOFLAGS) "$(OK_COLOR)==> Running webserver with envs:[$(LOCAL_VARIABLES)]...$(NO_COLOR)"
	@$(LOCAL_VARIABLES) bin/$(BUILDOS)_$(BUILDARCH)/$(BIN_WEBSERVER) $(args)

## build: build
build:
	@echo $(ECHOFLAGS) "$(OK_COLOR)==> Building binary ($(BUILDOS)/$(BUILDARCH)/$(BIN_WEBSERVER))...$(NO_COLOR)"
	@echo $(ECHOFLAGS) $(ENVFLAGS) $(BUILDENV) $(GO) build -v $(BUILDFLAGS) -o bin/$(BUILDOS)_$(BUILDARCH)/$(BIN_WEBSERVER) ./cmd
	@$(ENVFLAGS) $(BUILDENV) $(GO) build -v -o bin/$(BUILDOS)_$(BUILDARCH)/$(BIN_WEBSERVER) ./cmd