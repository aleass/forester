package task

import (
	"Forester/config"
	client "go.etcd.io/etcd/client/v3"
	"time"
)

func newEtcd(conf *config.Config) (*client.Client, error) {
	cli, err := client.New(client.Config{
		Endpoints:   []string{conf.Etcd.Addr},
		DialTimeout: 3 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	return cli, nil
}
