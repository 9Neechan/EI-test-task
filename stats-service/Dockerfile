# Build stage
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
#RUN go build -o main cmd/grpc_server/main.go
RUN go build -o main ./cmd/grpc_server

RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xz
RUN mv migrate.linux-amd64 /usr/bin/migrate

# Run stage
FROM alpine
WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /usr/bin/migrate /usr/bin/migrate
COPY internal/db/migrations ./migration

COPY configs/cfg.env . 
COPY scripts/start.sh ./start.sh
COPY scripts/wait-for.sh ./wait-for.sh

RUN chmod +x ./start.sh
RUN chmod +x ./wait-for.sh

CMD ["./main"]
ENTRYPOINT ["/app/start.sh"]


# docker build -t stats-service .