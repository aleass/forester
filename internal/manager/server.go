package manager

import (
	config2 "Forester/config"
	"Forester/internal/manager/config"
	redis "cloud.google.com/go/redis/apiv1"
	"fmt"
	client "go.etcd.io/etcd/client/v3"
)

type Server struct {
	Config *config2.Config
	Etcd   *client.Client
	Redis  *redis.CloudRedisClient
}

func ServerInit(path string) (server *Server) {
	server = new(Server)
	var (
		err error
	)
	err = config.InitConfig(server, path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = server.newEtcd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}
