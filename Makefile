DOCKER_RUN = docker run --rm -v "$(CURDIR)/db:/src" -w /src kjconroy/sqlc generate --experimental

createdb:
	createdb --username=postgres --owner=postgres  simple_library

dropdb:
	dropdb --username=postgres simple_library

run:
	go run main.go
	
gen:
	$(DOCKER_RUN)




