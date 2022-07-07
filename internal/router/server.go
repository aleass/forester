package router

import (
	proto "Forester/grpc"
	"Forester/internal/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
)

type RouteServer struct {
	Config *pkg.Config
	Http   *gin.Engine
	Client proto.ApiClient
}

func ServerInit(path string) (*RouteServer, error) {
	server := new(RouteServer)
	conf, err := pkg.InitConfig(path)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	server.Config = conf
	server.Http = NewHttp(conf)
	server.Client = newClient(conf)
	server.initRouter()
	return server, nil
}
func (r *RouteServer) Close() {

}
