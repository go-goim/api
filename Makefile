SHELL:=/usr/bin/env bash

.DEFAULT_GOAL:=help

Srv ?= push
BinPath ?= bin/service.goim.$(Srv)
CmdPath ?= apps/$(Srv)/cmd/main.go
CfgPath ?= apps/$(Srv)/configs
IMAGE ?= goim/$(Srv)
VERSION ?= $(shell git describe --exact-match --tags 2> /dev/null || git rev-parse --abbrev-ref HEAD)

## env
export ROCKETMQ_GO_LOG_LEVEL=warn

## jwt
export JWT_SECRET="goim"

## enable config center
export ENABLE_CONFIG_CENTER=true

##################################################
# Development                                    #
##################################################

##@ Development

.PHONY: lint
lint: ## Run go lint against code.
	golangci-lint run ./... -v

.PHONY: vet
vet: ## Run go vet against code.
	go vet -v ./...

.PHONEY: test
test: ## Run test against code.
	go test -v ./...

##################################################
# Generate                                          #
##################################################

##@ Generate

.PHONY: gen-protoc
gen-protoc: ## Run protoc command to generate pb code.
	# call gen_proto.sh
	./gen_proto.sh .

.PHONY: tools-install
tools-install: ## Install tools.
	go get -u github.com/golang/protobuf/protoc-gen-go

.PHONY: generate
generate: ## generate code by run go generate
	go generate ./...

##################################################
# General                                        #
##################################################

##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)