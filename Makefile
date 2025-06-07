SHELL := /bin/bash

run:
	go run ./tg/cmd/bot/main.go

lint:
	golangci-lint run ./tg/...

test_race:
	go test -race -timeout=60s -count 1 ./...

test:
	go clean -testcache
	go test ./...

tidy:
	find . -type f -name "go.mod" -exec dirname {} \; | xargs -I {} sh -c 'echo "Running go mod tidy in {}"; cd {} && go get -u ./... && go mod tidy'

db_backup:
	docker exec -t ssbot_postgres mkdir -p /backup
	docker exec -t ssbot_postgres pg_dump -U user -d ssbot_db -F c -f /backup/db.dump
	docker cp ssbot_postgres:/backup/db.dump ./var/backup/db.dump
	zip ./var/backup/db_$(shell date +"%Y%m%d-%H%M%S").zip ./var/backup/db.dump
	unzip -l ./var/backup/db_*.zip
	rm ./var/backup/db.dump


.PHONY: run lint test_race test tidy db_backup