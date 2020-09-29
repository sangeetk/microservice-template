all:
	go build

graphql:
	go run github.com/99designs/gqlgen generate

grpc:
	protoc --go_out=plugins=grpc:servers/grpc servers/grpc/greet.proto