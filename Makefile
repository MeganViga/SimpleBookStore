postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root book_store
accessdb:
	docker exec -it postgres12 psql -U root book_store
migratecreate:
	migrate create -ext sql -dir db/migration -seq init_schema
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/book_store?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/book_store?sslmode=disable" -verbose down
.Phony: postgres creatdb accessdb migratecreate migrateup migratedown