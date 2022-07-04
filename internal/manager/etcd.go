package manager

import (
	client "go.etcd.io/etcd/client/v3"
	"time"
)

func (s *Server) newEtcd() error {
	cli, err := client.New(client.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 3 * time.Second,
	})
	defer cli.Close()
	if err != nil {
	}
	s.Etcd = cli
	return nil
}
