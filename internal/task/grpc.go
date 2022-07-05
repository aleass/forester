package task

import (
	"Forester/config"
	proto "Forester/grpc"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func newGrpc(conf *config.Config) {
	grpcServer := grpc.NewServer()
	proto.RegisterTaskServiceServer(grpcServer, &service{})
	lis, err := net.Listen("tcp", conf.Grpc.Addr)
	if err != nil {
		fmt.Println(err)
	}
	grpcServer.Serve(lis)
}

type service struct {
}

func (s service) SendTask(server proto.TaskService_SendTaskServer) error {
	//TODO implement me
	panic("implement me")
}


