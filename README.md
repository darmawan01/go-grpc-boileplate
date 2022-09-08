# Go Microservices

```
> go verison
go version go1.19 darwin/arm64
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

## How to get started?

Requirements:

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
    services/hello/hello.proto
```