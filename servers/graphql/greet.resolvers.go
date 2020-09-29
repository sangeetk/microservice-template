package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/sangeetk/microservice-template/data"
	"github.com/sangeetk/microservice-template/servers/graphql/generated"
	"github.com/sangeetk/microservice-template/servers/graphql/model"
	"github.com/sangeetk/microservice-template/service"
)

var svc service.Service

func (r *queryResolver) Greet(ctx context.Context, req model.GreetRequest) (*model.GreetResponse, error) {
	var request = data.GreetRequest{Name: req.Name}

	response, err := svc.Greet(ctx, &request)
	if err != nil {
		return nil, err
	}

	return &model.GreetResponse{Message: response.Message}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
