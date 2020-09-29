package config

type AppConfig struct {
	ApiServerPort     int `env:"API_SERVER_PORT,default=8080"`
	GraphqlServerPort int `env:"GRAPHQL_SERVER_PORT,default=9090"`
	GrpcServerPort    int `env:"GRPC_SERVER_PORT,default=50051"`
}
