package manager

import (
	"Forester/config"
	proto "Forester/grpc"
	context "context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func newServer(conf *config.Config) {
	grpcServer := grpc.NewServer()
	proto.RegisterApiServerServer(grpcServer, &service{})
	lis, err := net.Listen("tcp", conf.Grpc.Addr)
	if err != nil {
		fmt.Println(err)
	}
	grpcServer.Serve(lis)
}

type service struct {
}

func (s service) AddUrl(ctx context.Context, list *proto.UrlList) (*proto.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Limit(ctx context.Context, list *proto.UrlList) (*proto.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) GetTaskCount(ctx context.Context, empty *proto.Empty) (*proto.TaskCount, error) {
	//TODO implement me
	panic("implement me")
}
