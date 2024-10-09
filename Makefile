migrate-up-all:
	migrate -path db/migrations -database "postgresql://root:123456@0.0.0.0:5432/bank?sslmode=disable" -verbose up

migrate-down-all:
	migrate -path db/migrations -database "postgresql://root:123456@0.0.0.0:5432/bank?sslmode=disable" -verbose down

migrate-up:
	migrate -path db/migrations -database "postgresql://root:123456@0.0.0.0:5432/bank?sslmode=disable" -verbose up 1

migrate-down:
	migrate -path db/migrations -database "postgresql://root:123456@0.0.0.0:5432/bank?sslmode=disable" -verbose down 1

sqlc:
	find db/sqlc -type f ! -name '*_test.go' -delete
	sqlc generate

test:
	go test -v -cover ./...

clean-test-cache:
	go clean -testcache

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/Luks17/Go-Microservices-MC/db/repository Store

.PHONY: migrate-up migrate-down migrate-up-all migrate-down-all sqlc test clean-test-cache server mock
