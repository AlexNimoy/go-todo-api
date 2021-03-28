include .env
export

postgres_url = postgres://$(POSTGRES_PASSWORD):$(POSTGRES_USER)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_PASSWORD)?sslmode=disable

migrate_up:
	migrate -path ./schema -database $(postgres_url) up

migrate_down:
	migrate -path ./schema -database $(postgres_url) down

db_cli:
	docker-compose exec postgres psql -U $(POSTGRES_USER)

run_srv:
	go run ./cmd/server/main.go