package task

import (
	"Forester/internal/pkg"
	client3 "go.etcd.io/etcd/client/v3"
)

type Server struct {
	Config *pkg.Config
	Etcd   *client3.Client
	task   chan *TasksObj
	addr   string
	log    *pkg.MyLog
}

var server *Server

func ServerInit(path string) *Server {
	server = new(Server)
	server.task = make(chan *TasksObj, 1000)
	conf, err := pkg.InitConfig(path)
	server.addr = pkg.InternalIP() + ":" + conf.TaskGrpc.Port
	if err != nil {
		panic(err.Error())
	}
	server.Config = conf
	server.Etcd, err = newEtcd(conf)
	if err != nil {
		panic(err.Error())
	}
	go server.newGrpc()
	server.register()
	server.log = pkg.New("task", true)
	return server
}
func (s *Server) Close() {

}
