FROM golang:alpine AS builder

RUN mkdir /src
ADD . /src
WORKDIR /src

RUN CGO_ENABLED=0 GOOS=linux go build -o api cmd/api/main.go

FROM alpine:latest AS production

COPY --from=builder /src/api .
CMD ["./api"]