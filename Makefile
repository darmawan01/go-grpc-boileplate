DIND_PREFIX ?= $(HOME)
ifneq ($(HOST_PATH),)
DIND_PREFIX := $(HOST_PATH)
endif
ifeq ($(CACHE_PREFIX),)
	CACHE_PREFIX=/tmp
endif

PREFIX=$(shell echo $(PWD) | sed -e s:$(HOME):$(DIND_PREFIX):)

include .env

WORKDIR := $(PWD)

infra:
	@docker-compose up -d db
	@# docker-compose up -d adminer

run:
	@go run cmd/api/main.go

test:
	@go test ./... -coverprofile=cover.out && go tool cover -html=cover.out

build:
	@go build cmd/api/main.go

build-image:
	@docker build --no-cache\
		-f build/Dockerfile\
		-t go-grpc-boilerplate:v1 .

run-image:
	@docker run --rm \
		--name go-grpc-boilerplate-test \
		--env-file .env -p 8080:8080 \
		docker.io/library/go-grpc-boilerplate:v1
	

