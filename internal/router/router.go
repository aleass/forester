package router

import (
	"Forester/config"
	proto "Forester/grpc"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func NewHttp(c *config.Config) *gin.Engine {
	engine := gin.New()
	engine.Use(recoverHandler)
	go func() {
		fmt.Println("http:", c.Http.Addr)
		if err := engine.Run(c.Http.Addr); err != nil {
			panic(err)
		}
	}()
	return engine
}
func (s *RouteServer) initRouter() {
	s.Http.GET("/add_url", s.AddUrl)
	s.Http.GET("/limit", s.Limit)
	s.Http.GET("/get_task_count", s.GetTaskCount)
}
func (s *RouteServer) AddUrl(c *gin.Context) {
	value, ok := c.GetQuery("url")
	if !ok {
		return
	}
	var list = &proto.UrlList{Url: []string{value}}
	s.Client.AddUrl(context.Background(), list)
}

func (s *RouteServer) GetTaskCount(c *gin.Context) {
	s.Client.GetTaskCount(context.Background(), &proto.Empty{})
}

func (s *RouteServer) Limit(c *gin.Context) {
	value, ok := c.GetQuery("rate")
	if !ok {
		return
	}
	res, _ := strconv.Atoi(value)
	s.Client.Limit(context.Background(), &proto.LimitDown{Rate: int64(res)})
}
