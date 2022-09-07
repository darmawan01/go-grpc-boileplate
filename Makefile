include .env

WORKDIR := $(PWD)

run:
	go run cmd/api/main.go

test:
	go test ./... -coverprofile=cover.out && go tool cover -html=cover.out

build:

seeds:
	

