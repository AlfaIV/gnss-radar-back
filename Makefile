.generate-graphql:
	go run github.com/99designs/gqlgen generate

build:
#	go get github.com/99designs/gqlgen/graphql/handler/lru@v0.17.49
#	go get github.com/99designs/gqlgen/graphql/handler/extension@v0.17.49

	docker compose up --build

up:
	docker compose up

docker-clear:
	docker stop $(docker ps -aq)
	docker rm $(docker ps -aq)
	docker volume rm $(docker volume ls -q)

run-main:
	go run cmd/gnss-radar/main.go

init:
	go mod init github.com/Gokert/gnss-radar

generate:
	# go generate ./api/proto/gnss-radar/gnss-radar.proto
	protoc --go_out=./pb --go-grpc_out=./pb ./api/proto/gnss-radar/gnss-radar.proto
	go generate ./api/proto/gnss-radar/gnss-radar.proto

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


