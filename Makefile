.PHONY: run build migrate swagger

COMPOSE_FILE=docker-compose.yml
DOCKER_COMPOSE=docker-compose -f $(COMPOSE_FILE)

run:
	$(DOCKER_COMPOSE) up --build

down:
	$(DOCKER_COMPOSE) down

migrate:
	# контейнер базы данных для выполнения миграций
	docker exec -it $$(docker ps -qf "name=music-library-db-1") bash -c "go run internal/db/migrate.go"

swagger:
	swag init -g cmd/main.go
