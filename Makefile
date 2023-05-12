postgres:
	docker run --name golang01 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:latest
createdb:
	docker exec -it golang01 createdb --username=root --owner=root simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5433/simple_bank?sslmode=disable" -verbose up 
migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5433/simple_bank?sslmode=disable" -verbose down 
dropdb:
	docker exec -it golang01 dropdb simple_bank
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
.PHONY: postgres createdb dropdb migrateup migratedown