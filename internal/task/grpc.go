package task

import (
	proto "Forester/grpc"
	"context"
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

func (s serviceGrpc) GetTaskCount(ctx context.Context, empty *proto.Empty) (*proto.Count, error) {
	fmt.Println("task run get task count", len(server.task))
	return &proto.Count{Num: int64(len(server.task))}, nil
}

func (s serviceGrpc) Limit(ctx context.Context, down *proto.LimitDown) (*proto.Response, error) {
	if down.Rate == -1 {
		limit = maxLimit
	} else {
		limit = uint64(down.Rate)
	}
	fmt.Println("task run limit :", limit)
	return &proto.Response{}, nil
}

func (s serviceGrpc) SendTask(res proto.Task_SendTaskServer) error {
	//accept
	go func() {
		for true {
			task, err := res.Recv()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			server.task <- &TasksObj{
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
