FROM golang:1.22-alpine AS builder

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -o app ./cmd/gnss-radar/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /build/app .
COPY --from=builder /build/configurations ./configurations

CMD ["./app"]
