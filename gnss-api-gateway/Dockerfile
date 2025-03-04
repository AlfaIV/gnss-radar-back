# Stage 1: Build
FROM golang:1.23.0-alpine AS build

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /source
COPY .. .

RUN go mod download && \
    go build -ldflags="-s -w" -o /app/api-gateway ./gnss-api-gateway/cmd/main.go
FROM alpine:3.18

RUN apk --no-cache add ca-certificates tzdata

ENV GATEWAY_ADDR=:8000 \
    TZ=UTC \
    AUTH_ADDR=gnss-auth:8001 \
    USER_ADDR=gnss-user:8002 

COPY --from=build /app/gateway /gateway
COPY --from=build /source/config /config/

RUN adduser -D -g '' appuser && \
    chown -R appuser:appuser /gateway /config

USER appuser

EXPOSE 8000

CMD ["/gateway"]