postgres:
	docker run --name postgres17 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -v postgres_data:/var/lib/postgresql/data --rm -d postgres:17-alpine

postgresstop:
	docker stop postgres17

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres17 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose down

sqlc:
	@echo "generate code by slc check /sqlc"
	sqlc generate

test:
	@echo "test all layer and show coverage"
	go test -v -cover ./...

.PHONY: postgres postgresstop createdb dropdb migrateup migratedown sqlc test