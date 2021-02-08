package main

import (
	"github.com/didchain/didchain-lightnode-go/node"
	"log"
	"os"
	"os/signal"
	"syscall"
)


var stop chan os.Signal

func main() {

	cfg:= node.InitNodeConf()

	node:= node.NewNode(cfg)

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