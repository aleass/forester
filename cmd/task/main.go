package main

import (
	"Forester/internal/task"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	server, err := task.ServerInit("config/config.yaml")
	if err != nil {
		panic(err.Error())
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	<-c
	server.Close()
}
