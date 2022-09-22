.PHONY: protobuf

start:
	go run cmd/server/*.go

protobuf:
	buf generate

store:
	sqlc generate

migrate:
	migrate --source file://migrations --database postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable up

docker:
	docker-compose up -d