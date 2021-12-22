.PHONY: build test docker api

PROTO_BUILD_DIR = ../../..
DOCKER_OPTS ?= --rm
TARGET_DIR ?= ./
# TEST_ARGS = -v | grep -c RUN

VERSION := $(shell git describe --tags --abbrev=0)

help:
	@echo "Service building targets"
	@echo "  build : build service command"
	@echo "  test  : run test suites"
	@echo "  docker: build docker image"
	@echo "  api: compile protobuf files for go"
	@echo "Env:"
	@echo "  DOCKER_OPTS : default docker build options (default : $(DOCKER_OPTS))"
	@echo "  TEST_ARGS : Arguments to pass to go test call"

api:
	if [ ! -d "./pkg/api" ]; then mkdir -p "./pkg/api"; else  find "./pkg/api" -type f -delete &&  mkdir -p "./pkg/api"; fi
	find ./api/logging_service/*.proto -maxdepth 1 -type f -exec protoc {} --proto_path=./api --go_out=plugins=grpc:$(PROTO_BUILD_DIR) \;

logging_service:
	go build -o $(TARGET_DIR) ./cmd/logging_service_app

build: logging_service
	# Dummy

test:
	./test/test.sh $(TEST_ARGS)

docker:
	docker build -t  github.com/influenzanet/logging-service:$(VERSION)  -f build/docker/Dockerfile $(DOCKER_OPTS) .
