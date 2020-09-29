package graphql

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/sangeetk/microservice-template/config"
	"github.com/sangeetk/microservice-template/servers/graphql/generated"
	"github.com/sangeetk/microservice-template/service"
)

var svc service.Service

func Server(ctx context.Context, c config.AppConfig) {

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%d/ for GraphQL playground", c.GraphqlServerPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", c.GraphqlServerPort), nil))

}
