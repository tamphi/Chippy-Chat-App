postgresinit:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:16-alpine

postgres:
	docker exec -it postgres16 psql

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root chippy-chat

dropdb:
	docker exec -it postgres16 dropdb chippy-chat

migrateup:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5432/chippy-chat?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5432/chippy-chat?sslmode=disable" -verbose down
build-server-image:
	docker build -f Dockerfile -t chippy-chat-server --no-cache .
docker-run-server:
	docker run -dp 127.0.0.1:8080:8080 chippy-chat-server
.PHONY: postgresinit postgres createdb dropdb migrateup migratedown build-server-image docker-run-server