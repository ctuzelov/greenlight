run:
	go run ./cmd/api

.PHONY: db/psql
db/psql:
	@docker exec -it greenlight_db_1 psql -U admin -d account

## db/init: initialization of database from dump
.PHONY: db/init
db/init:
	@docker run --name postgres -e POSTGRES_USER=admin -e POSTGRES_DB=greenlight_db_1 -e POSTGRES_PASSWORD=admin -d -p 5432:5432 postgres
	@sleep 1 && cat db/data/db_dump.sql | docker exec -i greenlight_db_1 psql -U admin -d account

## db/stop: dumping of db statement and removing container with db
.PHONY: db/stop
db/stop:
	@docker exec greenlight_db_1 pg_dump -U admin -d account > db/data/db_dump.sql
	@docker stop greenlight_db_1
	@docker rm greenlight_db_1