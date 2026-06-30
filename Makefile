include .env
export

PROJECT_ROOT := $(shell echo $$PWD)

env-up:
	docker compose up -d todoapp-postgres

env-down:
	docker compose down

env-clean:
	powershell -Command "$$ans = Read-Host 'очистить? [y/n]'; if ($$ans -eq 'y') { docker compose down -v; Remove-Item -Path 'out/pgdata' -Recurse -Force -ErrorAction SilentlyContinue; Write-Host 'done' } else { Write-Host 'canceled' }"

migrate-create:
	docker compose run --rm todoapp-postgres-migrate create -ext sql -dir /migrations -seq $(name)

migrate-up:
	@echo "Ждем запуск PostgreSQL..."
	@timeout /t 5 /nobreak >nul
	docker compose run --rm todoapp-postgres-migrate -path /migrations -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DB}?sslmode=disable" up

migrate-down:
	docker compose run --rm todoapp-postgres-migrate -path /migrations -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DB}?sslmode=disable" down

env-port-forward:
	docker compose up -d todoapp-postgres-port-forwarder
env-port-close:
	docker compose down port-forwarder
test-target:
	@echo "value: $(var)"