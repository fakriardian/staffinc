# staffinc

## General info

API for emas digital with go, kafka, postgres

## Table of contents

- [General info](#general-info)
- [Technologies](#technologies)
- [Setup](#setup)

## Technologies

Project is created with:

- Go
- GCC or TDM-GCC if using windows
- Kafka
- Zookeeper
- Kafka-ui
- Postgres

## Setup

To run this project, install it locally using npm:

```
$ cd ./misc && mv .env-example .env (set value .env)
$ docker-compose --env-file .env up -d
$ cd ../ && mv .env-example .env (set value .env)
$ go mod tidy
$ go run ./cmd/main.go
```

for postman collection in [here](https://github.com/fakriardian/staffinc/blob/main/staffinc.postman_collection.json)
