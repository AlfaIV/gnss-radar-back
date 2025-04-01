FROM --platform=linux/amd64 golang:1.23.0-alpine AS build

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /source
COPY .. .

RUN apk --no-cache add ca-certificates tzdata

RUN apk add --no-cache \
    protobuf-dev \
    protobuf \
    make

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

RUN make generate

RUN go mod download && \
    go mod tidy && \
    go build -ldflags="-s -w" -o /app/user ./gnss-user/cmd/main.go

FROM debian:bookworm-slim

RUN apk --no-cache add ca-certificates tzdata

# Установка grpc_health_probe
RUN wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.4.37/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

ENV USER_ADDR=:8002 \
    TZ=UTC
COPY --from=build /app/user /user

RUN useradd -u 1001 -m appuser && \
    chown -R adduser:appuser /user

USER appuser

EXPOSE 8002

CMD ["/user"]