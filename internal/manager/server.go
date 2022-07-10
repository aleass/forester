package manager

import (
	proto "Forester/grpc"
	"Forester/internal/pkg"
	"context"
	"fmt"
	client "go.etcd.io/etcd/client/v3"
)

type Server struct {
	Config *pkg.Config
	Etcd   *client.Client
	url    chan string
	crawl  map[string]*clients
	limit  int64
	c, a   int
	log    *pkg.MyLog
}

var server *Server

func ServerInit(path string) *Server {
	server = new(Server)
	server.crawl = make(map[string]*clients, 100)
	conf, err := pkg.InitConfig(path)
	if err != nil {
		panic(err.Error())
	}
	server.Config = conf
	server.url = make(chan string, 1000)
	server.log = pkg.New("manager", true)
	err = server.newEtcd()
	if err != nil {
		panic(err.Error())
	}
	go newServer(conf)
	server.GetClient()
	go server.watch()
	go server.run()
	return server
}

func (s *Server) run() {
	for true {
		select {
		case url := <-s.url:
			var min int
			var cli *clients
			for _, res := range s.crawl {
				if !res.isDoing {
					res.isDoing = true
					go s.OrderClient(res, url)
					break
				}

				if val := len(res.taskList); val < min || min == 0 {
					min = val
					cli = res
				}
			}
			if cli != nil {
				cli.taskList <- url
			}
		}
	}
}

func (s *Server) OrderClient(client *clients, url string) {
	obj := *client.client
	res, err := obj.SendTask(context.Background())
	if err != nil {
		fmt.Println(err.Error())
	}
	err = res.Send(&proto.TaskReq{
		Url:  url,
		Uuid: 0,
	})
	if err != nil {
		s.log.Error("manager.OrderClient:send task err:" + err.Error())
	}
	go func() {
		for true {
			task, err := res.Recv()
			if err != nil {
				s.log.Error("manager.OrderClient:Rec err:" + err.Error())
				return
			}
			fmt.Println(task.Uuid)
			count, err := obj.GetTaskCount(context.Background(), &proto.Empty{})
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			client.taskCount = count.Num
		}
	}()
	for url := range client.taskList {
		err = res.Send(&proto.TaskReq{
			Url:  url,
			Uuid: 0,
		})
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

func (s *Server) Limit() {
	for _, c := range s.crawl {
		(*c.client).Limit(context.Background(), &proto.LimitDown{Rate: s.limit})
	}
}

func (s *Server) Close() {
}
