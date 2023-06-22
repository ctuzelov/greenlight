run:
	go run ./cmd/api

.PHONY: db/psql
db/psql:
	@docker exec -it postgres psql -U admin -d postgres

## db/init: initialization of database from dump
.PHONY: db/init
db/init:
	@docker run --name greenlight_db_1 -e POSTGRES_USER=admin -e POSTGRES_DB=greenlight_db_1 -e POSTGRES_PASSWORD=123 -d -p 5433:5433 postgres
	@sleep 1 && cat db/data/db_dump.sql | docker exec -i greenlight_db_1 psql -U admin -d postgres

.PHONY: db/recover
db/recover:
	@sleep 1 && cat db/data/db_dump.sql | docker exec -i postgres psql -U admin -d postgres
## db/stop: dumping of db statement and removing container with db
.PHONY: db/stop
db/stop:
	@docker exec postgres pg_dump -U admin -d postgres > db/data/db_dump.sql
	@docker stop postgres
	@docker rm postgres