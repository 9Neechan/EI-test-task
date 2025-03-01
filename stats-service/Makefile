DB_URL=postgresql://root:secret@localhost:5432/stats?sslmode=disable

postgres:
	 docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17.3-alpine

createdb:
	 docker exec -it postgres17 createdb --username=root --owner=root stats

dropdb:
	 docker exec -it postgres17 dropdb stats
	 
migrateup:
	migrate -path internal/db/migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path internal/db/migrations -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate -f configs/sqlc.yaml

cleanmodcache:
	 go clean -modcache 


test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/9Neechan/simple-bank/db/sqlc Store



evans:
	evans --host localhost --port 9090 -r repl

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock db_docs db_schema proto evans