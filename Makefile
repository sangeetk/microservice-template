all:
	go build

gqlgen:
	go run github.com/99designs/gqlgen generate

protoc:
	protoc --go_out=plugins=grpc:servers/grpc servers/grpc/greet.proto

clean:
	go clean
