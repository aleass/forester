package router

import (
	config2 "Forester/config"
	"Forester/internal/config"
	redis "cloud.google.com/go/redis/apiv1"
	"fmt"
)

type RouteServer struct {
	Config *config2.Config
	Redis  *redis.CloudRedisClient
}

func ServerInit(path string) (*RouteServer, error) {
	server := new(RouteServer)
	conf, err := config.InitConfig(path)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	server.Config = conf
	return server, nil
}
