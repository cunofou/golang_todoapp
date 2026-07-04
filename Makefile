include .env
export

<<<<<<< Updated upstream
PROJECT_ROOT := $(shell echo $$PWD)

env-up:
	docker compose up -d todoapp-postgres

env-down:
	docker compose down
=======
# Для Windows используем cd вместо pwd
PROJECT_ROOT := $(shell cd)
>>>>>>> Stashed changes

env-clean:
	powershell -Command "$$ans = Read-Host 'очистить? [y/n]'; if ($$ans -eq 'y') { docker compose down -v; Remove-Item -Path 'out/pgdata' -Recurse -Force -ErrorAction SilentlyContinue; Write-Host 'done' } else { Write-Host 'canceled' }"

<<<<<<< Updated upstream
migrate-create:
	docker compose run --rm todoapp-postgres-migrate create -ext sql -dir /migrations -seq $(name)

migrate-up:
	@echo "Ждем запуск PostgreSQL..."
	@timeout /t 5 /nobreak >nul
	docker compose run --rm todoapp-postgres-migrate -path /migrations -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DB}?sslmode=disable" up
=======
env-up: ## env: Запустить окружение проекта
	@docker compose up -d todoapp-postgres
>>>>>>> Stashed changes

migrate-down:
	docker compose run --rm todoapp-postgres-migrate -path /migrations -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DB}?sslmode=disable" down

<<<<<<< Updated upstream
env-port-forward:
	docker compose up -d todoapp-postgres-port-forwarder
env-port-close:
	docker compose down port-forwarder
test-target:
	@echo "value: $(var)"
todoapp-run:
	set LOGGER_FOLDER=out\logs&& go mod tidy && go run cmd/todoapp/main.go tidy && go run cmd/todoapp/main.go
=======
env-cleanup: ## env: Очистить окружение проекта
	@echo "Очистить все volume файлы окружения? Опасность утери данных. [y/N]: " & set /p ans= & \
	if "!ans!" == "y" ( \
		docker compose down todoapp-postgres port-forwarder && \
		rmdir /s /q "$(PROJECT_ROOT)\out\pgdata" 2>nul && \
		echo Файлы окружения очищены \
	) else ( \
		echo Очистка окружения отменена \
	)

env-port-forward: ## env: Открыть порты сервисов окружения
	@docker compose up -d port-forwarder

env-port-close: ## env: Закрыть порты сервисов окружения
	@docker compose down port-forwarder

logs-cleanup: ## env: Очистить файлы логов из out/logs
	@echo "Очистить все log файлы? Опасность утери логов. [y/N]: " & set /p ans= & \
	if "!ans!" == "y" ( \
		rmdir /s /q "$(PROJECT_ROOT)\out\logs" 2>nul && \
		echo Файлы логов очищены \
	) else ( \
		echo Очистка логов отменена \
	)

swagger-gen: ## env: Сгенерировать актуальную Swagger спецификацию
	@docker compose run --rm swagger \
		init \
		-g cmd/todoapp/main.go \
		-o docs \
		--parseInternal \
		--parseDependency

ps: ## env: Посмотреть запущенные Docker Compose сервисы
	@docker compose ps

migrate-create: ## PostgreSQL: Создать новую версию схемы данных
	@if "$(seq)" == "" ( \
		echo Отсутсвует необходимый параметр seq. Пример: make migrate-create seq=init && \
		exit 1 \
	) else ( \
		docker compose run --rm todoapp-postgres-migrate \
			create \
			-ext sql \
			-dir /migrations \
			-seq "$(seq)" \
	)

migrate-up: ## PostgreSQL: Накатить миграции
	@make migrate-action action=up

migrate-down: ## PostgreSQL: Откатить миграции
	@make migrate-action action=down

migrate-action:
	@if "$(action)" == "" ( \
		echo Отсутсвует необходимый параметр action. Пример: make migrate-action action=up && \
		exit 1 \
	) else ( \
		docker compose run --rm todoapp-postgres-migrate \
			-path /migrations \
			-database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@todoapp-postgres:5432/$(POSTGRES_DB)?sslmode=disable \
			$(action) \
	)

todoapp-run: ## Golang приложение: Запустить локально на хост-системе (для локальной разработки)
	@if not exist "$(PROJECT_ROOT)\out\logs" mkdir "$(PROJECT_ROOT)\out\logs"
	@set LOGGER_FOLDER=$(PROJECT_ROOT)\out\logs && \
	set POSTGRES_HOST=localhost && \
	set REDIS_HOST=localhost && \
	go mod tidy && \
	go run $(PROJECT_ROOT)\cmd\todoapp\main.go

todoapp-deploy: ## Golang приложение: Запустить в Docker Compose сервисе (для деплоя)
	@docker compose up -d --build todoapp

todoapp-undeploy: ## Golang приложение: Остановить Docker Compose сервис
	@docker compose down todoapp

help: ## Показать справку по командам
	@echo === Центр управления проектом ===
	@echo.
	@echo Доступные команды:
	@for /f "tokens=1,* delims=:" %%a in ('findstr /b /r "[a-zA-Z_-][a-zA-Z_-]*:.*?##" $(MAKEFILE_LIST)') do @echo   %%a
>>>>>>> Stashed changes
