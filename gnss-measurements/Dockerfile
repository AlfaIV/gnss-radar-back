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
    go build -ldflags="-s -w" -o /app/measurements ./gnss-measurements/cmd/main.go

FROM debian:bookworm-slim

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    ca-certificates \
    tzdata \
    wget && \
    rm -rf /var/lib/apt/lists/*

RUN wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.4.37/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

ENV MEASUREMENTS_ADDR=:8003 \
    TZ=UTC \
    MINIO_API_ADDR=s3:9002 \
    SATELLITES_ADDR=satellites:8010
    
COPY --from=build /app/measurements /measurements

RUN useradd -u 1001 -m appuser && \
    chown -R appuser:appuser /measurements

USER appuser

EXPOSE 8003

CMD ["/measurements"]