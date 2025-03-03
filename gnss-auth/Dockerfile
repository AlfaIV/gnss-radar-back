FROM golang:1.23.0-alpine AS build

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /source
COPY . .

RUN go mod download && \
    go build -ldflags="-s -w" -o /app/auth ./cmd/main.go


FROM alpine:3.18

RUN apk --no-cache add ca-certificates tzdata

# Установка grpc_health_probe
RUN wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.4.37/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

ENV AUTH_ADDR=:8001 \
    REDIS_ADDR=db-redis-auth:6379 \
    TZ=UTC 

COPY --from=build /app/auth /auth

RUN adduser -D -g '' appuser && \
    chown -R appuser:appuser /auth

USER appuser

EXPOSE 8001

CMD ["/auth"]