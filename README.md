# Go Microservices

## How to get started?

Prerequisites

- `go version go1.19` or higher
- `build-esential`

[GRPC Quick Start](https://grpc.io/docs/languages/go/quickstart/)
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

```

Export GO bin:
```
export PATH="$PATH:$(go env GOPATH)/bin"
```

Generate proto:
```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    services/grpc/hello/hello.proto
```

## CMD's
`Note: Make sure you have 'make' command installed`

Run Project
```
make run
```

Run Test
```
make test
```

## How to contribute

No rules for now. Feel free to add issue first and optionally submit a PR. Cheers

[Conventional Commits Reference](https://www.conventionalcommits.org/en/v1.0.0/#specification)