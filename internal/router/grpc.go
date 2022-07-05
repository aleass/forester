package router

import (
	"Forester/config"
	proto "Forester/grpc"
	"google.golang.org/grpc"
)

func newClient(conf *config.Config) proto.ApiClient {
	conn, err := grpc.Dial(conf.ApiGrpc.Addr, grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	return proto.NewApiClient(conn)
}
