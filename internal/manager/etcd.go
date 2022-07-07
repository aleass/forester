package manager

import (
	proto "Forester/grpc"
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	client "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"time"
)

func (s *Server) newEtcd() error {
	cli, _ := client.New(client.Config{
		Endpoints:   []string{s.Config.Etcd.Addr},
		DialTimeout: 3 * time.Second,
	})
	s.Etcd = cli
	return nil
}

func (s *Server) watch() {
	watchChan := s.Etcd.Watch(context.Background(), s.Config.Etcd.TaskPre, client.WithPrefix())
	for true {
		select {
		case response := <-watchChan:
			for _, event := range response.Events {
				key := string(event.Kv.Key)
				if event.Type == mvccpb.PUT {
					fmt.Println("manager:watch put:", key)
					s.crawl[key] = s.NewClient(string(event.Kv.Value))
				} else {
					fmt.Println("manager:watch delete:", key)
					delete(s.crawl, key)
				}
			}
		}
	}
}

func (s *Server) GetClient() {
	res, err := s.Etcd.Get(context.Background(), s.Config.Etcd.TaskPre, client.WithPrefix())
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("发现服务：", len(res.Kvs))
	for _, v := range res.Kvs {
		s.crawl[string(v.Key)] = s.NewClient(string(v.Value))
	}
}

func (s *Server) NewClient(addr string) *clients {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	cli := proto.NewTaskClient(conn)
	return &clients{
		client:   &cli,
		taskList: make(chan string, 1000),
	}
}
