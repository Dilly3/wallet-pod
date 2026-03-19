up:
	docker compose up -d
build-up:
	docker compose up -d --build
down:
	docker compose down
logs:
	docker compose logs -f
migrate-dev:
	@for file in sql/migrations/*.sql; do \
		echo "Running $$file..."; \
		docker exec -i wallet_pod_db psql -U wallet_pod -d walletpod < $$file || exit 1; \
	done
	@echo "All migrations completed"
migrate-prod:
	migrate -path sql/migrations -database "postgres://$$DB_USER:$$DB_PASSWORD@$$DB_HOST:$$DB_PORT/$$DB_NAME?sslmode=$$DB_SSLMODE" up