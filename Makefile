#!/usr/bin/make
GOHOSTOS := $(shell go env GOHOSTOS)
GOPATH := $(shell go env GOPATH)

# remove hash "sed 's/-g[a-z0-9]\{7\}//'" or use "--dirty --broken" parameter
VERSION := $(shell git describe --tags --match "v*" --always | sed 's/-g[a-z0-9]\{7\}//')
BRANCH=$(shell git symbolic-ref -q --short HEAD)
REVISION=$(shell git rev-parse --short HEAD)
BUILT_AT := $(shell date +%FT%T%z)

# $(shell a=`basename $$PWD` && echo $$a)
APP_NAME = sample

BUILD_VARS := ${APP_NAME}/internal
LDFLAGS := -ldflags "-X ${BUILD_VARS}.Version=${VERSION} -X ${BUILD_VARS}.Branch=$(BRANCH) -X ${BUILD_VARS}.Revision=$(REVISION) -X ${BUILD_VARS}.BuiltAt=${BUILT_AT} -s -w"
# darwin:amd64 darwin:arm64
OS_ARCH=linux:amd64 linux:arm64

DOCKER_NAMESPACE := cr.uasse.com/${APP_NAME}
DOCKER_IMAGE := $(shell echo ${APP_NAME}  | awk -F '@' '{print "${DOCKER_NAMESPACE}/" $$0}')
DOCKER_ARGS := --build-arg APP_NAME=${APP_NAME}
DOCKER_FILE := infra/docker/Dockerfile
DOCKER_VERSION := ${VERSION}


PROTOBUF := "./api"
THIRD_PARTY := "./third_party"
PB_OUT_PATH := "."

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	CONF_PROTO_FILES=$(shell $(Git_Bash) -c "find **/conf -name *.proto")
	ENT_SCHEMA_PATH := $(shell find **/data -type d -name 'schema')
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find ${PROTOBUF} -name *.proto ! -path '**/shared/*'")
else
	CONF_PROTO_FILES=$(shell find **/conf -name *.proto)
	ENT_SCHEMA_PATH := $(shell find **/data -type d -name 'schema')
	API_PROTO_FILES=$(shell find ${PROTOBUF} -name *.proto ! -path '**/shared/*')
endif

.PHONY: init
# init env
init: install-tools

.PHONY: install-tools
# install tools
install-tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: config
# generate config proto
config:
	@protoc --proto_path=. \
		--proto_path=${THIRD_PARTY} \
		--go_out=paths=source_relative:${PB_OUT_PATH} \
		${CONF_PROTO_FILES}

.PHONY: _api_shared
# generate shared protobuf
_api_shared:
	@protoc --proto_path=. \
		--proto_path=${THIRD_PARTY} \
		--go_out=paths=source_relative:${PB_OUT_PATH} \
		--validate_out=paths=source_relative,lang=go:${PB_OUT_PATH} \
		--go-errors_out=paths=source_relative:${PB_OUT_PATH} \
		${PROTOBUF}/shared/**/*.proto

.PHONY: api
# generate api proto
api: _api_shared
	@protoc --proto_path=. \
		--proto_path=${THIRD_PARTY} \
		--go_out=paths=source_relative:${PB_OUT_PATH} \
		--go-http_out=paths=source_relative:${PB_OUT_PATH} \
		--go-grpc_out=paths=source_relative:${PB_OUT_PATH} \
		--validate_out=paths=source_relative,lang=go:${PB_OUT_PATH} \
		--go-errors_out=paths=source_relative:${PB_OUT_PATH} \
		$(API_PROTO_FILES)

.PHONY: swagger
# generate swagger
swagger:
	@protoc --proto_path=. \
		--proto_path=${PROTOBUF} \
		--proto_path=${THIRD_PARTY} \
		--openapiv2_out ${PB_OUT_PATH} \
		--openapiv2_opt logtostderr=true \
		--openapiv2_opt json_names_for_fields=false \
		${PROTOBUF}/interface/**/*.proto

.PHONY: build
# build
build:
	@mkdir -p bin/ && go build ${LDFLAGS} -trimpath -o ./bin/ ./...

.PHONY: build-cross
# build multi platform
build-cross:
	@$(foreach n, $(OS_ARCH), \
		os=$(shell echo "$(n)" | cut -d : -f 1); \
		arch=$(shell echo "$(n)" | cut -d : -f 2); \
		gomips=$(shell echo "$(n)" | cut -d : -f 3); \
		target_suffix=$${os}_$${arch}; \
		echo "Build $${os}-$${arch}..."; \
		mkdir -p ./bin && env CGO_ENABLED=0 GOOS=$${os} GOARCH=$${arch} GOMIPS=$${gomips} go build -trimpath ${LDFLAGS} -o ./bin/${APP_NAME}_$${target_suffix} ./cmd/${APP_NAME}; \
		echo "Build $${os}-$${arch} done"; \
	)

.PHONY: docker
# eg: make docker DOCKER_VERSION=0.1.0
docker:
	docker build -f ${DOCKER_FILE} ${DOCKER_ARGS} -t ${DOCKER_IMAGE}:${DOCKER_VERSION} . --load

.PHONY: docker-buildx
# docker multi platform build and push remote
docker-buildx:
	docker buildx build -f ${DOCKER_FILE} ${DOCKER_ARGS} -t ${DOCKER_IMAGE}:${DOCKER_VERSION} .  --platform linux/amd64,linux/arm64 --push

.PHONY: docker-push
# push docker images to registry
docker-push:
	docker push ${DOCKER_IMAGE}:${DOCKER_VERSION}


.PHONY: generate
# generate
generate:
	@go generate ./...

.PHONY: schema
# generate ent schema
schema:
ifeq ($(name),)
	@echo "schema is null. e.g: make schema name=User"
else
	@ent init --target ./data/schema $(name)
endif

.PHONY: ent
# generate ent code
ent:
	@$(foreach n, $(ENT_SCHEMA_PATH), \
		schema_path=$(shell echo "$(n)"); \
		out_path=$(shell echo "$(n)" | sed "s/$/schema/ent/g"); \
		go run entgo.io/ent/cmd/ent generate --feature sql/versioned-migration --target $${out_path} ./$${schema_path}; \
	)

.PHONY: wire
# generate wire
wire:
	@cd cmd/${APP_NAME} && wire

.PHONY: test
# test
test:
	@go test -v ./... -cover

.PHONY: run
# run app
run:
	@cd cmd/${APP_NAME} && go run . -

.PHONY: all
# generate all
all: ent api config wire swagger

.PHONY: clean
# clear all generated and binary the file
clean:
	@find . -type d -name 'ent' -o -name 'graphql' -o -name 'bin' ! -path "**/data/schema/*" ! -path './web' | xargs rm -rf;
	@find . -type f -name '*.pb.*' -o -name '*_gen.go' -o -name '*.swagger.json' ! -path './web' | xargs rm;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
