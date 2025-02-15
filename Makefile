postgres:
	docker run --name db-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=postgres@5502 -d postgres:17-alpine

createdb:
	docker exec -it db-postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it db-postgres dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:postgres@5502@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:postgres@5502@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY:	postgres createdb dropdb migrateup migratedown sqlc
