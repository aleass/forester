package task

import (
	config2 "Forester/config"
	"Forester/internal/config"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Server struct {
	Config *config2.Config
	Etcd   *clientv3.Client
	task   chan *TaskInfo
}

var server *Server

func ServerInit(path string) (*Server, error) {
	server = new(Server)
	server.task = make(chan *TaskInfo, 1000)
	conf, err := config.InitConfig(path)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	server.Config = conf
	server.Etcd, _ = newEtcd(conf)
	go newGrpc(conf)
	server.register()
	return server, nil
}
func (s *Server) Close() {

}
