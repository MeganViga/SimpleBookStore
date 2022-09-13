postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres
createdb:
	docker exec -it postgres createdb --username=root --owner=root book_store

dropdb:
	docker exec -it postgres dropdb book_store

accessdb:
	docker exec -it postgres psql -U root book_store

migratecreate:
	migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/book_store?sslmode=disable" -verbose up

migratedown:
		migrate -path db/migration -database "postgresql://root:secret@localhost:5432/book_store?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go


.PHONY: postgres createdb acessdb dropdb migratecreate migrateup migratedown sqlc test server