# staffinc

how to run

- `cd misc && mv .env-example .env`
- set value .env
- `docker-compose --env-file .env up -d`
- `cd ../ && mv .env-example .env`
- set value .env
- `go run ./cmd/main.go`
