package manager

import (
	proto "Forester/grpc"
	"context"
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
					s.crawl[key] = s.NewClient()
				} else {
					delete(s.crawl, key)
				}
			}
		}
	}
}

//func (s *Server) GetAllClient() {
//	res, err := s.Etcd.Get(context.Background(), s.Config.Etcd.TaskPre, client.WithPrefix())
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	for _, v := range res.Kvs {
//
//	}
//}

func (s *Server) NewClient() proto.TaskClient {
	conn, err := grpc.Dial(s.Config.ApiGrpc.Addr, grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	client := proto.NewTaskClient(conn)
	client.SendTask(context.Background())
	return client
}
