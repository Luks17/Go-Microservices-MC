migrate-up:
	migrate -path db/migrations -database "postgresql://root:123456@0.0.0.0:5432/bank?sslmode=disable" -verbose up

migrate-down:
	migrate -path db/migrations -database "postgresql://root:123456@0.0.0.0:5432/bank?sslmode=disable" -verbose down

sqlc:
	@if [ -d "./db/sqlc" ]; then \
		rm ./db/sqlc/*; \
	fi;
	sqlc generate;

.PHONY: migrate-up migrate-down sqlc
