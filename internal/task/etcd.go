package task

import (
	"Forester/internal/pkg"
	"context"
	client "go.etcd.io/etcd/client/v3"
	"time"
)

func newEtcd(conf *pkg.Config) (*client.Client, error) {
	cli, err := client.New(client.Config{
		Endpoints:   []string{conf.Etcd.Addr},
		DialTimeout: 3 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	return cli, nil
}

func (s *Server) register() {
	info := s.Config.Etcd.TaskPre
	key := info + s.addr
	// 创建一个租约
	resp, err := s.Etcd.Grant(context.TODO(), 5)
	if err != nil {
		panic(err.Error())
	}

	_, err = s.Etcd.Put(context.TODO(), key, s.addr, client.WithLease(resp.ID))

	ch, _ := s.Etcd.KeepAlive(context.TODO(), resp.ID)
	for {
		select {
		case <-ch:
		}
	}

}
