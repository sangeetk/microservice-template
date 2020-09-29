package main

import (
	"context"
	"log"

	"github.com/sangeetk/microservice-template/config"
	"github.com/sangeetk/microservice-template/servers/graphql"
	"github.com/sangeetk/microservice-template/servers/grpc"
	"github.com/sangeetk/microservice-template/servers/rest"

	"github.com/sethvargo/go-envconfig"
)

func main() {
	ctx := context.Background()

	var c config.AppConfig
	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal(err)
	}

	// REST Server
	go rest.Server(ctx, c)

	// GraphQL Server
	go graphql.Server(ctx, c)

	// GRPC Server
	grpc.Server(ctx, c)
}
