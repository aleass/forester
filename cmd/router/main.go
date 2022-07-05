package main

import (
	"Forester/internal/router"
	"fmt"
)

func main() {
	server, err := router.ServerInit("config/config.yaml")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(server)
}
