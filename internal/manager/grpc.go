package manager

import (
	proto "Forester/grpc"
	"Forester/internal/pkg"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type clientLink struct {
	member *clients
}
type clients struct {
	client    *proto.TaskClient
	isDoing   bool
	taskCount int64
	taskList  chan string
	next      *clients
	pre       *clients
}

func newServer(conf *pkg.Config) {
	grpcServer := grpc.NewServer()
	proto.RegisterApiServer(grpcServer, &service{})
	lis, err := net.Listen("tcp", conf.ApiGrpc.Addr)
	if err != nil {
		panic(err)
	}
	grpcServer.Serve(lis)
}

type service struct {
}

func (s service) Limit(ctx context.Context, down *proto.LimitDown) (*proto.Response, error) {
	fmt.Println("manager set limit rate:", down.Rate)
	server.limit = down.Rate
	server.Limit()
	return &proto.Response{}, nil
}

func (s service) GetTaskCount(ctx context.Context, empty *proto.Empty) (*proto.TaskCount, error) {
	var count int64
	for _, c := range server.crawl {
		count += c.taskCount
	}
	fmt.Println("manager get count:", count)
	return &proto.TaskCount{Count: count}, nil
}

func (s service) AddUrl(ctx context.Context, list *proto.UrlList) (*proto.Response, error) {
	go func() {
		fmt.Println("manager add url:", len(list.Url))
		for _, url := range list.Url {
			server.url <- url
		}
	}()
	return &proto.Response{}, nil
}

func (g *clientLink) Add(s *clients) {
	if g.member != nil {
		s.next = g.member.next
		g.member.pre = s
	}
	g.member = s
}

func (g *clientLink) Del(s *clients) {
	if s == nil {
		return
	}
	if s.pre == nil { //第一位
		g.member = s.next
		if s.next != nil { //有第二位
			s.next.pre = nil
		}
	} else if s.next == nil { //末尾
		s.pre.next = nil
	} else { //中间
		s.pre.next = s.next
		s.next.pre = s.pre
	}
}
