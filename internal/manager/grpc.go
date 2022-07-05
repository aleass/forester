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
	proto.RegisterApiServer(grpcServer, &service{})
	lis, err := net.Listen("tcp", conf.ApiGrpc.Addr)
	if err != nil {
		fmt.Println(err)
	}
	grpcServer.Serve(lis)
}
func newClient(conf *config.Config) proto.TaskClient {
	conn, err := grpc.Dial(conf.ApiGrpc.Addr, grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	return proto.NewTaskClient(conn)
}

type service struct {
}

func (s service) Limit(ctx context.Context, down *proto.LimitDown) (*proto.Response, error) {
	fmt.Println(down.Rate)
	return &proto.Response{}, nil
}

func (s service) GetTaskCount(ctx context.Context, empty *proto.Empty) (*proto.TaskCount, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) AddUrl(ctx context.Context, list *proto.UrlList) (*proto.Response, error) {
	go func() {
		for _, url := range list.Url {
			server.url <- url
		}
	}()
	return &proto.Response{}, nil
}
