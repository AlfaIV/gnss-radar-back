FROM golang:1.23.0-alpine AS build

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
    go build -ldflags="-s -w" -o /app/gnss-tasks ./gnss-tasks/cmd/main.go

FROM alpine:3.18

ENV TASKS_ADDR=:8004 \
    TZ=UTC 
COPY --from=build /app/gnss-tasks /tasks

RUN adduser -D -g '' appuser && \
    chown -R appuser:appuser /gateway /tasks

USER appuser

EXPOSE 8004

CMD ["/tasks"]