# Build stage
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
#RUN go build -o main cmd/grpc_server/main.go
RUN go build -o main ./cmd/http_server

RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xz
RUN mv migrate.linux-amd64 /usr/bin/migrate

# Run stage
FROM alpine
WORKDIR /app

COPY --from=builder /app/main .
COPY configs/cfg.env . 

CMD ["./main"]


# docker build -t api-gateway .