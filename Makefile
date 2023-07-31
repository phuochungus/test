DBNAME = library

createdb:
	createdb --username=postgres --owner=postgres $(DBNAME)

dropdb:
	dropdb --username=postgres $(DBNAME)

build:
	go build -o bin/main.exe main.go

mup:
	migrate -path db/migration -database "postgresql://postgres:123123123@localhost:5432/$(DBNAME)?sslmode=disable" -verbose up

mdown:
	migrate -path db/migration -database "postgresql://postgres:123123123@localhost:5432/$(DBNAME)?sslmode=disable" -verbose down

cm:
	migrate create -ext sql -dir db/migration -seq $(name)

f:
	migrate -path db/migration -database "postgresql://postgres:123123123@localhost:5432/$(DBNAME)?sslmode=disable" force $(v)
	





