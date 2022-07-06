package manager

import (
	config2 "Forester/config"
	proto "Forester/grpc"
	"Forester/internal/config"
	"context"
	"fmt"
	client "go.etcd.io/etcd/client/v3"
)

type Server struct {
	Config *config2.Config
	Etcd   *client.Client
	url    chan string
	crawl  map[string]*clients
	c, a   int
}

var server *Server

func ServerInit(path string) (*Server, error) {
	server = new(Server)
	server.crawl = make(map[string]*clients, 100)
	conf, err := config.InitConfig(path)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	server.Config = conf
	server.url = make(chan string, 1000)

	err = server.newEtcd()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	go newServer(conf)
	server.GetClient()
	go server.watch()
	go server.run()
	return server, nil
}

func (s Server) run() {
	for true {
		select {
		case url := <-s.url:
			var min int
			var cli *clients
			for _, res := range s.crawl {
				if !res.isDoing {
					res.isDoing = true
					go OrderClient(res, url)
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

func OrderClient(client *clients, url string) {
	res, err := (*client.client).SendTask(context.Background())
	if err != nil {
		fmt.Println(err.Error())
	}
	err = res.Send(&proto.TaskReq{
		Url:  url,
		Uuid: 0,
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	go func() {
		for true {
			task, err := res.Recv()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println(task.Uuid)
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

func (s Server) Close() {
}
