package manager

import (
	"Forester/config"
	proto "Forester/grpc"
	context "context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type clientLink struct {
	member *clients
}
type clients struct {
	client *proto.TaskClient
	next   *clients
	pre    *clients
}

func newServer(conf *config.Config) {
	grpcServer := grpc.NewServer()
	proto.RegisterApiServer(grpcServer, &service{})
	lis, err := net.Listen("tcp", conf.ApiGrpc.Addr)
	if err != nil {
		fmt.Println(err)
	}
	grpcServer.Serve(lis)
}
func (s Server) newClient(conf *config.Config) {
	conn, err := grpc.Dial(s.Config.ApiGrpc.Addr, grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	client := proto.NewTaskClient(conn)
	client.SendTask(context.Background())
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
