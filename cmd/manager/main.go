package main

import (
	"Forester/internal/manager"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	server := manager.ServerInit("config/config.yaml")
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	<-c
	server.Close()
}
