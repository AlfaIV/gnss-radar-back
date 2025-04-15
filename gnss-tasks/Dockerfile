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
    go build -ldflags="-s -w" -o /app/tasks ./gnss-tasks/cmd/main.go

FROM debian:bookworm-slim

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    ca-certificates \
    tzdata \
    wget && \
    rm -rf /var/lib/apt/lists/*

RUN wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.4.37/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

ENV TASKS_ADDR=:8004 \
    TZ=UTC 
COPY --from=build /app/tasks /tasks

RUN useradd -u 1001 -m appuser && \
    chown -R appuser:appuser /tasks

USER appuser

EXPOSE 8004

CMD ["/tasks"]