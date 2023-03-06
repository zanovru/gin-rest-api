-include .env
export $(shell sed 's/=.*//' .env)

.PHONY: build
build:
	go build -o ./build/apiserv ./cmd/api/main.go

run: build
	./build/apiserv

DOCKER_COMPOSE=docker compose

up:
	$(DOCKER_COMPOSE) up -d

down:
	$(DOCKER_COMPOSE) down

start: up
stop: down

docker-rebuild:
	$(DOCKER_COMPOSE) up --no-deps --build -d


DB_MIGRATIONS_SOURCE=db/migrations
DATABASE_URL='postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DATABASE}?sslmode=disable'


migration:
	migrate create -ext sql -dir $(DB_MIGRATIONS_SOURCE) -seq $(NAME)

migrate:
	migrate -path $(DB_MIGRATIONS_SOURCE) -database $(DATABASE_URL) up

migrate-drop:
	migrate -path $(DB_MIGRATIONS_SOURCE) -database $(DATABASE_URL) drop

migrate-down:
	migrate -path $(DB_MIGRATIONS_SOURCE) -database $(DATABASE_URL) down