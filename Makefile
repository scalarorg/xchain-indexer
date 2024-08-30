#!/usr/bin/make -f

BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
COMMIT := $(shell git log -1 --format='%H')

# don't override user values
ifeq (,$(VERSION))
  VERSION := $(shell git describe --tags)
  # if VERSION is empty, then populate it with branch's name and raw commit hash
  ifeq (,$(VERSION))
    VERSION := $(BRANCH)-$(COMMIT)
  endif
endif

TARGETPLATFORM := linux/amd64
ifeq ($(shell uname -m), arm64)
TARGETPLATFORM := linux/arm64
endif

ifndef $(PACKAGES)
	PACKAGES := ./main.go
endif
# default value, overide with: make -e FQCN="foo"
# FQCN = ghcr.io/scalarorg/xchains-indexer
FQCN = scalarorg/xchains-indexer
all: install

install: go.sum
	go install .

build:
	go build -o bin/xchains-indexer .

clean:
	rm -rf build

build-docker:
	docker build -t $(FQCN) -f ./Dockerfile \
	--build-arg TARGETPLATFORM=$(TARGETPLATFORM) \
	--build-arg PACKAGES=$(PACKAGES) .

build-docker-amd:
	docker build -t $(FQCN):$(VERSION) -f ./Dockerfile \
	--build-arg TARGETPLATFORM=linux/amd64 \
	--build-arg PACKAGES=$(PACKAGES) .

build-docker-arm:
	docker build -t $(FQCN):$(VERSION) -f ./Dockerfile \
	--build-arg TARGETPLATFORM=linux/arm64 \
	--build-arg PACKAGES=$(PACKAGES) .

.PHONY: lint
lint: ## Run golangci-linter
	golangci-lint run --out-format=tab

.PHONY: format
format: ## Formats the code with gofumpt
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/docs/*" | xargs gofumpt -w

###############################################################################
###                                Protobuf                                 ###
###############################################################################

proto-all: proto-update-deps proto-format proto-lint proto-gen

proto-gen:
	@echo "Create docker image for Protobuf files"
	@DOCKER_BUILDKIT=1 docker build -t scalarorg/xchains-indexer-proto-gen -f ./Dockerfile.protocgen .
	@echo "Generate Protobuf files"
	@$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace scalarorg/xchains-indexer-proto-gen sh ./scripts/protocgen.sh
	@echo "Generating Protobuf Swagger endpoint"
	@$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace scalarorg/xchains-indexer-proto-gen sh ./scripts/protoc-swagger-gen.sh
	@statik -src=./client/docs/static -dest=./client/docs -f -m

proto-format:
	@echo "Formatting Protobuf files"
	@$(DOCKER) run --rm -v $(CURDIR):/workspace \
	--workdir /workspace tendermintdev/docker-build-proto \
	$( find ./ -not -path "./third_party/*" -name "*.proto" -exec clang-format -i {} \; )

proto-lint:
	@echo "Linting Protobuf files"
	@$(DOCKER_BUF) lint

proto-check-breaking:
	@$(DOCKER_BUF) breaking --against $(HTTPS_GIT)#branch=main

TM_CRYPTO_TYPES     	= third_party/proto/tendermint/crypto
TM_ABCI_TYPES       	= third_party/proto/tendermint/abci
TM_TYPES            	= third_party/proto/tendermint/types
TM_VERSION          	= third_party/proto/tendermint/version
TM_LIBS             	= third_party/proto/tendermint/libs/bits
TM_P2P              	= third_party/proto/tendermint/p2p

GOGO_PROTO_TYPES    	= third_party/proto/gogoproto
GOOGLE_API_TYPES		= third_party/proto/google/api
GOOGLE_PROTOBUF_TYPES	= third_party/proto/google/protobuf
COSMOS_PROTO_TYPES  	= third_party/proto/cosmos_proto		
# For some reason ibc expects confio proto files to be in the main folder
CONFIO_TYPES        	= third_party/proto