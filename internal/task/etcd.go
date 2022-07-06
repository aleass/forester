package task

import (
	"Forester/config"
	"context"
	client "go.etcd.io/etcd/client/v3"
	"net"
	"strings"
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

func (s *Server) register() {
	info := s.Config.Etcd.TaskPre
	key := info + InternalIP()
	// 创建一个租约
	resp, err := s.Etcd.Grant(context.TODO(), 5)
	if err != nil {
		panic(err.Error())
	}

	_, err = s.Etcd.Put(context.TODO(), key, InternalIP(), client.WithLease(resp.ID))

	ch, _ := s.Etcd.KeepAlive(context.TODO(), resp.ID)
	for {
		select {
		case <-ch:
		}
	}

}

// InternalIP return internal ip.
func InternalIP() string {
	inters, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, inter := range inters {
		if inter.Flags&net.FlagUp != 0 && !strings.HasPrefix(inter.Name, "lo") {
			addrs, err := inter.Addrs()
			if err != nil {
				continue
			}
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String()
					}
				}
			}
		}
	}
	return ""
}
