postgres:
	docker run --name pg16a -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16.0-alpine

createdb:
	docker exec -it pg16a createdb --username=root --owner=root bank_db

dropdb:
	docker exec -it pg16a dropdb bank_db

migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/bank_db?sslmode=disable" -verbose up

migratedown:
	 migrate -path db/migration -database "postgres://root:secret@localhost:5432/bank_db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown
