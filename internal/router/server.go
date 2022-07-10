package router

import (
	proto "Forester/grpc"
	"Forester/internal/pkg"
	"github.com/gin-gonic/gin"
)

type RouteServer struct {
	Config *pkg.Config
	Http   *gin.Engine
	Client proto.ApiClient
	log    *pkg.MyLog
}

func ServerInit(path string) *RouteServer {
	server := new(RouteServer)
	conf, err := pkg.InitConfig(path)
	if err != nil {
		panic(err.Error())
	}
	server.Config = conf
	server.Http = NewHttp(conf)
	server.Client = newClient(conf)
	server.initRouter()
	server.log = pkg.New("router", true)
	return server
}
func (r *RouteServer) Close() {

}
