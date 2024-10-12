.generate-graphql:
	go run github.com/99designs/gqlgen generate

up:
	docker-compose up --build

run:
	go run cmd/gnss-radar/main.go

init:
	go mod init github.com/Gokert/gnss-radar

docker-drop:
	docker-compose down

lint:
	golangci-lint run

gg:
	go run github.com/99designs/gqlgen generate

gg-init:
	go run github.com/99designs/gqlgen init
