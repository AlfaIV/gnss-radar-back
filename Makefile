.generate-graphql:
	go run github.com/99designs/gqlgen generate

up:
	go get github.com/99designs/gqlgen/graphql/handler/lru@v0.17.49
	go get github.com/99designs/gqlgen/graphql/handler/extension@v0.17.49

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
	go get github.com/99designs/gqlgen@v0.17.49
	go get github.com/99designs/gqlgen/graphql/handler/lru@v0.17.49
	go get github.com/99designs/gqlgen/graphql/handler/extension@v0.17.49

	go run github.com/99designs/gqlgen generate

gg-init:
	go run github.com/99designs/gqlgen init
