FROM --platform=linux/amd64 golang:1.23.0-alpine AS build

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /source
COPY .. .

# Установка зависимостей для Alpine (правильно для первого этапа)
RUN apk --no-cache add ca-certificates tzdata && \
    apk add --no-cache \
    protobuf-dev \
    protobuf \
    make

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

RUN make generate && \
    go mod download && \
    go mod tidy && \
    go build -ldflags="-s -w" -o /app/auth ./gnss-auth/cmd/main.go

# Финальный этап - исправления для Debian
FROM debian:bookworm-slim

# Установка зависимостей через apt (вместо apk)
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    ca-certificates \
    tzdata \
    wget && \
    rm -rf /var/lib/apt/lists/*

# Установка grpc_health_probe
RUN wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.4.37/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

ENV AUTH_ADDR=:8001 \
    REDIS_ADDR=db-redis-auth:6379 \
    TZ=UTC 

COPY --from=build /app/auth /auth

RUN useradd -u 1001 -m appuser && \
    chown -R appuser:appuser /auth

USER appuser

EXPOSE 8001

CMD ["/auth"]