.PHONY: up down logs psql reset-db status run run-debug run-production

up:
	sudo docker-compose up -d
down:
	sudo docker-compose down
logs:
	sudo docker-compose logs -f db

psql:
	sudo docker exec -it poject_db psql -U postgres

reset-db:
	sudo docker-compose down -v
	sudo docker-compose up -d

status:
	sudo docker ps

run:
	go run cmd/server/main.go

run-debug:
	go run ./cmd/server/main.go --log-level=debug

run-production:
	go run ./cmd/server/main.go --log-level=info --log-output=logs/app.log