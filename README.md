# staffinc

How to run

- Install `gcc` or `tdm-gcc` if you using windows
- `go mod tidy`
- `cd misc && mv .env-example .env`
- set value .env
- `docker-compose --env-file .env up -d`
- `cd ../ && mv .env-example .env`
- set value .env
- `go run ./cmd/main.go`
