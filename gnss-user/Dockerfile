FROM golang:1.23.0-alpine AS build

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /source
COPY .. .

RUN go mod download && \
    go build -ldflags="-s -w" -o /app/user ./gnss-user/cmd/main.go

FROM alpine:3.18

RUN apk --no-cache add ca-certificates tzdata

# Установка grpc_health_probe
RUN wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.4.37/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

ENV USER_ADDR=:8002 \
    TZ=UTC
COPY --from=build /app/user /user

RUN adduser -D -g '' appuser && \
    chown -R appuser:appuser /user

USER appuser

EXPOSE 8002

CMD ["/user"]