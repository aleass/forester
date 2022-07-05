package main

import (
	"Forester/internal/manager"
	"fmt"
)

func main() {
	server, err := manager.ServerInit("config/config.yaml")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(server)
}
