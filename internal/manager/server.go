package manager

import (
	config2 "Forester/config"
	"Forester/internal/config"
	redis "cloud.google.com/go/redis/apiv1"
	"fmt"
	client "go.etcd.io/etcd/client/v3"
)

type Server struct {
	Config *config2.Config
	Etcd   *client.Client
	Redis  *redis.CloudRedisClient
}

func ServerInit(path string) (*Server, error) {
	server := new(Server)
	conf, err := config.InitConfig(path)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	server.Config = conf

	err = server.newEtcd()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	newServer(conf)
	return server, nil
}
func (s Server) Close() {

}
