package service

import (
	"context"

	"github.com/sangeetk/microservice-template/data"
)

type GreetService interface {
	Greet(context.Context, *data.GreetRequest) (*data.GreetResponse, error)
}

type Service struct{}
