KEEP_IMAGES = postgres redis

build-images:
	docker build -t gateway-image -f ./gnss-api-gateway/Dockerfile .
	docker build -t auth-image -f ./gnss-auth/Dockerfile .
	docker build -t user-image -f ./gnss-user/Dockerfile .

docker-clear:
	@echo "Остановка всех запущенных контейнеров..."
	docker stop $(docker ps -aq)

	@echo "Удаление всех контейнеров..."
	docker rm $(docker ps -aq)

	@echo "Удаление всех томов..."
	docker volume rm $(docker volume ls -q)

	@echo "Удаление всех Docker-образов, кроме: $(KEEP_IMAGES)"
	docker images --format '{{.Repository}}:{{.Tag}} {{.ID}}' | grep -v -e 'nginx' -e 'postgres' -e 'redis' | awk '{print $2}' | xargs -r docker rmi

PROTO_ROOT = api/proto
GO_OUT = .

generate:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/proto/auth/auth.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/proto/common/common.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/proto/user/user.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/proto/measurements/measurements.proto

start-networks:
	@if [ -z "$$(docker network ls --filter name=gnss-radar-net -q)" ]; then \
		docker network create --driver bridge gnss-radar-net; \
	fi

start-services:
	docker compose -f ./deployments/docker-compose.yaml up -d

stop-services:
	docker compose -f ./deployments/docker-compose.yaml down

deploy: build-images start-networks start-services