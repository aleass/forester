package router

import (
	"Forester/config"
	proto "Forester/grpc"
	"fmt"
	"google.golang.org/grpc"
)

func NewEtcd(conf *config.Config) proto.ApiServerClient {
	conn, err := grpc.Dial(conf.Etcd.Addr, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	return proto.NewApiServerClient(conn)
}
