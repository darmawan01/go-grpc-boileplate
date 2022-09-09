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

seeds:
	

