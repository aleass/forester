package router

import (
	proto "Forester/grpc"
	"Forester/internal/pkg"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func NewHttp(c *pkg.Config) *gin.Engine {
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
func (r *RouteServer) initRouter() {
	r.Http.GET("/add_url", r.AddUrl)
	r.Http.GET("/limit", r.Limit)
	r.Http.GET("/get_task_count", r.GetTaskCount)
}
func (r *RouteServer) AddUrl(c *gin.Context) {
	value, ok := c.GetQuery("url")
	if !ok {
		return
	}
	var list = &proto.UrlList{Url: []string{value}}
	_, err := r.Client.AddUrl(context.Background(), list)
	if err != nil {
		c.JSON(500, err.Error())
	}
}

func (r *RouteServer) GetTaskCount(c *gin.Context) {
	res, err := r.Client.GetTaskCount(context.Background(), &proto.Empty{})
	if err != nil {
		c.JSON(500, err.Error())
	}
	c.JSON(200, res.Count)
}

func (r *RouteServer) Limit(c *gin.Context) {
	value, ok := c.GetQuery("rate")
	if !ok {
		return
	}
	res, _ := strconv.Atoi(value)
	_, err := r.Client.Limit(context.Background(), &proto.LimitDown{Rate: int64(res)})
	if err != nil {
		c.JSON(500, err.Error())
	}
	c.JSON(200, "ok")
}
