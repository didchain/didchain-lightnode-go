package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)


var stop chan os.Signal

func main() {

	cfg:=InitNodeConf()

	node:=NewNode(cfg)

	node.Start()

	signal.Notify(stop,
		syscall.SIGKILL,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	s:=<-stop

	log.Println("get signal:",s)

	node.Stop()

}

func init()  {
	stop = make(chan os.Signal,8)
}