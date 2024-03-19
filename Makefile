up:
	@echo "Starting containers"
	docker compose -f docker-compose.yaml up --remove-orphans

up-prod:
	@echo "Starting production containers"
	 docker compose -f docker-compose.production.yaml up -d

down:
	@echo "Stopping containers"
	docker compose -f docker-compose.yaml down

down-prod:
	@echo "Stopping production containers"
	docker compose -f docker-compose.production.yaml down

run-api:
	@echo "Running API"
	go run api/cmd/main.go

.PHONY: up down up-prod down-prod
