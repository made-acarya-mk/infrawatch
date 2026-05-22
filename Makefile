APP_NAME=infrawatch

build:
	docker compose build

up:
	docker compose up -d

down:
	docker compose down

restart:
	docker compose down
	docker compose up -d --build

logs:
	docker compose logs -f

health:
	./scripts/health-check.sh

backup:
	./scripts/backup.sh

test:
	curl -f http://localhost:8082/health