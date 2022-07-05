package router

import (
	"Forester/config"
	"fmt"
	"github.com/gin-gonic/gin"
)

func NewHttp(c *config.Config) *gin.Engine {
	engine := gin.New()
	engine.Use(recoverHandler)
	go func() {
		if err := engine.Run(c.Http.Addr); err != nil {
			panic(err)
		}
	}()
	return engine
}
func (s *RouteServer) initRouter() {
	s.Http.GET("/add_url/", s.AddUrl)
	s.Http.GET("/limit", s.Limit)
	s.Http.GET("/get_task_count", s.GetTaskCount)
}
func (s *RouteServer) AddUrl(c *gin.Context) {
	fmt.Println("abc")
	//s.Etcd.AddUrl(context.Background())
}
func (s *RouteServer) GetTaskCount(c *gin.Context) {

}
func (s *RouteServer) Limit(c *gin.Context) {

}
