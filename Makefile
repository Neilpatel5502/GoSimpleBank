postgres:
	docker run --name db-postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=postgres@5502 -d postgres:17-alpine

createdb:
	docker exec -it db-postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it db-postgres dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:postgres@5502@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:postgres@5502@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:postgres@5502@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:postgres@5502@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/Neilpatel5502/GoSimpleBank/db/sqlc Store

.PHONY:	postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc test server mock
