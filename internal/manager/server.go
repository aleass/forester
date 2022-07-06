package manager

import (
	config2 "Forester/config"
	proto "Forester/grpc"
	"Forester/internal/config"
	"fmt"
	client "go.etcd.io/etcd/client/v3"
)

type Server struct {
	Config *config2.Config
	Etcd   *client.Client
	url    chan string
	crawl  map[string]proto.TaskClient
	c, a   int
}

var server *Server

func ServerInit(path string) (*Server, error) {
	server = new(Server)
	conf, err := config.InitConfig(path)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	server.Config = conf
	server.url = make(chan string, 1000)
	server.crawl = make(map[string]proto.TaskClient, 1000)

	err = server.newEtcd()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	go newServer(conf)
	go server.watch()
	return server, nil
}

func (s Server) run() {
	for true {
		select {
		case <-s.url:
			//client := *s.crawl[s.c]
			//client.SendTask(context.Background(),)
			s.c++
			if s.c == s.a {
				s.c = 0
			}
		}
	}
}

func (s Server) Close() {

}
