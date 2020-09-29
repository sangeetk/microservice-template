package service

import (
	"context"

	"github.com/sangeetk/microservice-template/data"
)

func (s *Service) Greet(ctx context.Context, req *data.GreetRequest) (*data.GreetResponse, error) {
	res := &data.GreetResponse{
		Message: "Hello " + req.Name + "!",
	}
	return res, nil
}
