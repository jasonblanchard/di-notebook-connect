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

lambda:
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/lambda cmd/lambda/*.go
	zip -j ./bin/lambda.zip ./bin/lambda

# pushlambda: lambda
# 	aws s3 cp ./bin/lambda.zip s3://di-notebook-prod-b287d59/${GIT_SHA}/lambda.zip
