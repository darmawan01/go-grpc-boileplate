# Go Microservices

# Example architecture
![Architecture](https://github.com/darmawan01/assets/blob/main/integration.png)

## How to get started?

Prerequisites

- `go version go1.19` or higher
- `build-esential`

`To run this app please take a look for required environment variable`

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
make protoc
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

Build executable binary
```
make build
```

Build docker image version
```
make build-image
```

Run docker image version
```
make run-image
```
```
Notes: 
When running, this app will looking for .env if not found it will take host variable instead.
Other than that, you also can get the configs from vault. See the environment variable
```

## Environment Variable
Required:
- `ENV` App Environment, Eg: development, staging, production
- `APP_PORT` App port
- `GRPC_PORT` GRPC port
- `DB_HOST` Db Host
- `DB_PORT` Db port
- `DB_USER` Db user
- `DB_PASS` Db Pass
- `DB_NAME` Db name
- `DB_MAX_OPEN_CONN` Db max open connection
- `DB_MAX_IDLE_CONN` Db max idle connection
- `DB_MAX_LIFE_TIME` Db max life time
- `JWT_SECRET_KEY`  JWT secret key

Redis:
- `REDIS_ENABLED` (true|false) to enable or disable redis connection
- `REDIS_HOST` Redis host
- `REDIS_PORT` Redis port
- `REDIS_USER` Redis user
- `REDIS_PASS` Redis password

Vault:
- `VAULT_ENABLED` (true|false) to enable or disable to get configs from vault
- `VAULT_ADDRESS` Vault address
- `VAULT_TOKEN` Vault token
- `VAULT_SERVICE_NAME` Vault service name
- `VAULT_SECRET_NAME` Vault secret name


## How to contribute

No rules for now. Feel free to add issue first and optionally submit a PR. Cheers

[Conventional Commits Reference](https://www.conventionalcommits.org/en/v1.0.0/#specification)