KEEP_IMAGES = nginx postgres redis

.generate-graphql:
	go run github.com/99designs/gqlgen generate

build:
#	go get github.com/99designs/gqlgen/graphql/handler/lru@v0.17.49
#	go get github.com/99designs/gqlgen/graphql/handler/extension@v0.17.49

	docker compose up --build

up:
	docker compose up

docker-clear:
	@echo "Остановка всех запущенных контейнеров..."
	docker stop $(docker ps -aq)

	@echo "Удаление всех контейнеров..."
	docker rm $(docker ps -aq)

	@echo "Удаление всех томов..."
	docker volume rm $(docker volume ls -q)

	@echo "Удаление всех Docker-образов, кроме: $(KEEP_IMAGES)"
	docker images --format '{{.Repository}}:{{.Tag}} {{.ID}}' | grep -v -e 'nginx' -e 'postgres' -e 'redis' | awk '{print $2}' | xargs -r docker rmi

run-main:
	go run cmd/gnss-radar/main.go

init:
	go mod init github.com/Gokert/gnss-radar

generate:
	# go generate ./api/proto/gnss-radar/gnss-radar.proto
	#protoc --go_out=./pb --go-grpc_out=./pb ./api/proto/gnss-radar/gnss-radar.proto
	#go generate ./api/proto/gnss-radar/gnss-radar.proto
	rm -rf gen
	buf generate

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


