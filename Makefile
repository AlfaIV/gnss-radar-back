.generate-graphql:
	go run github.com/99designs/gqlgen generate

build:
	go get github.com/99designs/gqlgen/graphql/handler/lru@v0.17.49
	go get github.com/99designs/gqlgen/graphql/handler/extension@v0.17.49

	docker compose up --build

up:
	docker compose up

run:
	go run cmd/gnss-radar/main.go

init:
	go mod init github.com/Gokert/gnss-radar

docker-drop:
	docker-compose down

lint:
	golangci-lint run

gg:
	go get github.com/99designs/gqlgen@v0.17.49
	go get github.com/99designs/gqlgen/graphql/handler/lru@v0.17.49
	go get github.com/99designs/gqlgen/graphql/handler/extension@v0.17.49

	go run github.com/99designs/gqlgen generate

gg-init:
	go run github.com/99designs/gqlgen init

migrate:
	psql postgres -c "drop database if exists gnss;"
	createdb gnss
	goose -allow-missing -dir migrations postgres "dbname=gnss sslmode=disable" up

SQL_NAME ?= new_sql

create-goose:
	goose -dir ./migrations -s create $(SQL_NAME) sql


