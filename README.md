# kkmoney
A simple imlementation of a bank REST api service for crud operations, user authentication & authorization and making basic transactions 

## Technologies
- Gin framework
- mysql
- Docker
- sqlc for autogeneration of crud code for interacting with sql databases
- mockdb for simulating database for test purposes
- golang-migrate for managing database migrations

## Test
run `make test` or `go test -v -cover -timeout 30s`

## Migrate
run `make migratedown` or `make migrateup`

## Start server
run `make server` or `go run main.go`

## Start Mysql server
run `make mysql`
