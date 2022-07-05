package router

import (
	config2 "Forester/config"
	proto "Forester/grpc"
	"Forester/internal/config"
	"fmt"
	"github.com/gin-gonic/gin"
)

type RouteServer struct {
	Config *config2.Config
	Http   *gin.Engine
	Etcd   proto.ApiServerClient
}

func ServerInit(path string) (*RouteServer, error) {
	server := new(RouteServer)
	conf, err := config.InitConfig(path)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	server.Config = conf
	server.Http = NewHttp(conf)
	server.Etcd = NewEtcd(conf)
	server.initRouter()
	return server, nil
}
func (r *RouteServer) Close() {

}
