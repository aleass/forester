package task

import (
	proto "Forester/grpc"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

func (s *Server) newGrpc() {
	grpcServer := grpc.NewServer()
	proto.RegisterTaskServer(grpcServer, &serviceGrpc{})
	lis, err := net.Listen("tcp", s.addr)
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
	var i int
	for true {
		for info := range server.task {
			fmt.Println("start download:", info.Url)
			httpLimit(info.Url, strconv.Itoa(i))
			res.Send(&proto.Finish{Uuid: info.Uuid})
			i++
		}
	}
	return nil
}
