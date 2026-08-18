include .env
export

export PROJECT_ROOT=$(shell pwd)

env-up:
	@docker compose up -d letsgo-postgres

env-down:
	@docker compose down letsgo-postgres

env-cleanup:
	@read -p "Are you sure to cleanup pg volumes? [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down letsgo-postgres && \
		rm -rf out/pgdata && \
		echo "Pg volumes successfully cleared"; \
	else \
		echo "Cleanup rejected"; \
	fi

env-port-forward:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "seq parameter required"; \
		echo "Example: make migrate-create seq=init"; \
		exit 1; \
	fi;
	docker compose run --rm letsgo-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "action parameter required"; \
		echo "Example: make migrate-action action=up"; \
		exit 1; \
	fi;
	docker compose run --rm letsgo-postgres-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@letsgo-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"
