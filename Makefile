SHELL:=/usr/bin/env bash

.DEFAULT_GOAL:=help

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
	./gen_proto.sh

.PHONY: protoc

COLOR := "\e[1;36m%s\e[0m"

PROTO_ROOT := .
PROTO_IGNORE := ./third_party
PROTO_FILES = $(shell find $(PROTO_ROOT) -path $(PROTO_IGNORE) -prune -o -type f -name "*.proto" -print)
PROTO_DIRS = $(sort $(dir $(PROTO_FILES)))
PROTO_IMPORTS := --proto_path=. --proto_path=./third_party
PROTO_OPTIONS := --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. --validate_out=lang=go,paths=source_relative:.

define NEWLINE


endef

protoc: ## Generate pb code.
	@printf $(COLOR) "Generating protobuf code..."
# Run protoc command to generate pb code.
	$(foreach PROTO_DIR,$(PROTO_DIRS),\
		protoc $(PROTO_IMPORTS) $(PROTO_OPTIONS) $(PROTO_DIR)*.proto \
	$(NEWLINE))
	@printf $(COLOR) "Done."

.PHONY: tools-install
tools-install: ## Install tools.
	go get -u github.com/golang/protobuf/protoc-gen-go

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
