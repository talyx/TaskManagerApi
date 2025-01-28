.PHONY: up down logs psql reset-db status run run-debug run-production
include .env
export $(shell sed 's/=.*//' .env)

up:
	sudo docker-compose up -d
down:
	sudo docker-compose down

psql:
	sudo docker exec -it $(CONTAINER_NAME) psql -U postgres

reset-db:
	sudo docker-compose down -v
	sudo docker-compose up -d

status:
	sudo docker ps

run-debug:
	go run ./cmd/server/main.go --log-level=debug

run-production:
	go run ./cmd/server/main.go --log-level=info --log-output=logs/app.log