migrate-up:
	migrate -path db/migrations -database "postgresql://root:123456@0.0.0.0:5432/bank?sslmode=disable" -verbose up

migrate-down:
	migrate -path db/migrations -database "postgresql://root:123456@0.0.0.0:5432/bank?sslmode=disable" -verbose down

.PHONY: migrate-up migrate-down
