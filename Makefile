KEEP_IMAGES = nginx postgres redis

build-images:
	docker build -t gateway-image -f gnss-api-gateway/Dockerfile .
	docker build -t auth-image -f gnss-auth/Dockerfile .

docker-clear:
	@echo "Остановка всех запущенных контейнеров..."
	docker stop $(docker ps -aq)

	@echo "Удаление всех контейнеров..."
	docker rm $(docker ps -aq)

	@echo "Удаление всех томов..."
	docker volume rm $(docker volume ls -q)

	@echo "Удаление всех Docker-образов, кроме: $(KEEP_IMAGES)"
	docker images --format '{{.Repository}}:{{.Tag}} {{.ID}}' | grep -v -e 'nginx' -e 'postgres' -e 'redis' | awk '{print $2}' | xargs -r docker rmi

start-networks:
	@if [ -z "$$(docker network ls --filter name=gnss-radar-net -q)" ]; then \
		docker network create --driver bridge gnss-radar-net; \
	fi

start-services:
	docker compose -f deployments/docker-compose.yml up -d

stop-services:
	docker compose -f deployments/docker-compose.yml down

deploy: docker-clear build-images start-networks start-services