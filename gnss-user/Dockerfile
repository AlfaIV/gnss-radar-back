FROM golang:1.23.0-alpine AS build

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /source
COPY . .

RUN go mod download && \
    go build -ldflags="-s -w" -o /app/user ./cmd/main.go

FROM alpine:3.18

RUN apk --no-cache add ca-certificates tzdata

#Потом на CICD прокину

ENV USER_ADDR=:8002 \
    TZ=UTC \
    PG_URL=https://gp.com \

COPY --from=build /app/user /user

RUN adduser -D -g '' appuser && \
    chown -R appuser:appuser /auth

USER appuser

EXPOSE 8002

CMD ["/user"]