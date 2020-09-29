package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/sangeetk/microservice-template/config"
	"github.com/sangeetk/microservice-template/data"
	"github.com/sangeetk/microservice-template/servers/grpc/pb"
	"github.com/sangeetk/microservice-template/service"

	"google.golang.org/grpc"
)

type server struct{}

var svc service.Service

func (s *server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	var request = data.GreetRequest{Name: req.Name}

	response, err := svc.Greet(ctx, &request)
	if err != nil {
		return nil, err
	}

	return &pb.GreetResponse{Message: response.Message}, nil
}

func Server(ctx context.Context, c config.AppConfig) {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", c.GrpcServerPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
