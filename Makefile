up:
	docker compose up -d
build-up:
	docker compose up -d --build
down:
	docker compose down
logs:
	docker compose logs -f
migrate-dev-up:
	@for file in sql/migrations/*.up.sql; do \
		echo "Running $$file..."; \
		docker exec -i wallet_pod_db psql -U walletpod -d walletpod_db < $$file || exit 1; \
	done
	@echo "All UP migrations completed"

migrate-dev-down:
	@files=$$(ls -1r sql/migrations/*.down.sql 2>/dev/null); \
	for file in $$files; do \
		echo "Running $$file..."; \
		docker exec -i wallet_pod_db psql -U walletpod -d walletpod_db < $$file || exit 1; \
	done
	@echo "All DOWN migrations completed"

migrate-dev: migrate-dev-up

migrate-prod-up:
	migrate -path sql/migrations -database "postgres://$$DB_USER:$$DB_PASSWORD@$$DB_HOST:$$DB_PORT/$$DB_NAME?sslmode=$$DB_SSLMODE" up

migrate-prod-down:
	migrate -path sql/migrations -database "postgres://$$DB_USER:$$DB_PASSWORD@$$DB_HOST:$$DB_PORT/$$DB_NAME?sslmode=$$DB_SSLMODE" down

migrate-prod: migrate-prod-up