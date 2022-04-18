mysql:
	docker run --name mysqlCont -e MYSQL_ROOT_PASSWORD=grandmaster002 -e MYSQL_DATABASE=kkmoney -p 3307:3306 mysql:8.0.28-oracle

openmysql:
	docker exec -it mysqlCont mysql -u root --password kkmoney

sqlc:
	sqlc generate

migrateup:
	migrate -path db/migration -database "mysql://root:grandmaster002@tcp(localhost:3307)/kkmoney" -verbose up

migratedown:
	migrate -path db/migration -database "mysql://root:grandmaster002@tcp(localhost:3307)/kkmoney" -verbose down

server:
	go run main.go

resetdb: migrateup migratedown

test:
	go test -v -cover -timeout 30s ./... 

mock:
	mockgen -destination db/mock/store.go -package mockdb github.com/AbdulkarimOgaji/kkmoney3/db/sqlc Store


.PHONY: openmysql mysql migrateup migratedown test resetdb server mock