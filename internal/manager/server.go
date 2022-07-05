package manager

import (
	config2 "Forester/config"
	proto "Forester/grpc"
	"Forester/internal/config"
	"fmt"
	client "go.etcd.io/etcd/client/v3"
)

type Server struct {
	Config *config2.Config
	Etcd   *client.Client
	Client proto.TaskClient
	url    chan string
}

var server *Server

func ServerInit(path string) (*Server, error) {
	server = new(Server)
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
	server.Client = newClient(conf)
	return server, nil
}

func (s Server) run() {
	for true {

	}
}

func (s Server) Close() {

}
