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
	proto.RegisterTaskServer(grpcServer, &serviceGrpc{})
	lis, err := net.Listen("tcp", conf.TaskGrpc.Addr)
	if err != nil {
		fmt.Println(err)
	}
	grpcServer.Serve(lis)
}

type serviceGrpc struct {
}

func (s serviceGrpc) SendTask(res proto.Task_SendTaskServer) error {
	//accept
	go func() {
		for true {
			task, err := res.Recv()
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			server.task <- &TaskInfo{
				Uuid: task.Uuid,
				Url:  task.Url,
			}
		}
	}()
	for true {
		for info := range server.task {
			httpLimit(info.Url)
			res.Send(&proto.Finish{Uuid: info.Uuid})
		}
	}
	return nil
}
