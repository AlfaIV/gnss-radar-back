

.generate-graphql:
	go run github.com/99designs/gqlgen generate

up:
	docker-compose --env-file .env up --build