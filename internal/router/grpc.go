package router

import (
	proto "Forester/grpc"
	"Forester/internal/pkg"
	"google.golang.org/grpc"
)

func newClient(conf *pkg.Config) proto.ApiClient {
	conn, err := grpc.Dial(conf.ApiGrpc.Addr, grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	return proto.NewApiClient(conn)
}
