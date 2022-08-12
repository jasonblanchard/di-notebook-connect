.PHONY: protobuf

start:
	go run cmd/server/*.go

protobuf:
	buf generate

store:
	sqlc generate

migrate:
	migrate --source file://migrations --database postgres://postgres:mysecretpassword@localhost:50162/postgres?sslmode=disable up