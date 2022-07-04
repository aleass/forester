package main

import (
	"Forester/internal/manager"
	"fmt"
)

func main() {
	server := manager.ServerInit("config/config.yaml")
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	fmt.Println(server)

}
