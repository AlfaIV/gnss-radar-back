FROM golang:1.22-alpine AS builder
WORKDIR /build
COPY ../.. .

RUN go build -o app ./cmd/gnss-radar/main.go

CMD ./app