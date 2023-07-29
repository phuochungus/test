DOCKER_RUN = docker run --rm -v "$(CURDIR)/db:/src" -w /src kjconroy/sqlc generate --experimental
DBNAME = library


createdb:
	createdb --username=postgres --owner=postgres $(DBNAME)

dropdb:
	dropdb --username=postgres $(DBNAME)

run:
	go run main.go
	
gen:
	$(DOCKER_RUN)

mup:
	migrate -path db/migration -database "postgresql://postgres:123123123@localhost:5432/$(DBNAME)?sslmode=disable" -verbose up

mdown:
	migrate -path db/migration -database "postgresql://postgres:123123123@localhost:5432/$(DBNAME)?sslmode=disable" -verbose down

cm:
	migrate create -ext sql -dir db/migration -seq $(name)

f:
	migrate -path db/migration -database "postgresql://postgres:123123123@localhost:5432/$(DBNAME)?sslmode=disable" force $(v)
	





