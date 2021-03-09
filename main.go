package main

import (
	"fmt"
	"github.com/didchain/didchain-lightnode-go/cmd"
	"github.com/didchain/didchain-lightnode-go/config"
	"github.com/didchain/didchain-lightnode-go/node"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var rootCmd = &cobra.Command{
	Use:   "node",
	Short: "node",
	Long:  `usage description`,
	Run:   mainRun,
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var param struct {
	version bool
	did     string
}

var stop chan os.Signal

func mainRun(_ *cobra.Command, _ []string) {

	if param.version {
		fmt.Println("0.0.1")
		return
	}

	cfg := config.InitNodeConf()

	node := node.NewNode(cfg)

	go cmd.StartCmdService()

	node.Start()

	signal.Notify(stop,
		syscall.SIGKILL,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	s := <-stop

	log.Println("get signal:", s)

	node.Stop()

}

func init() {
	rootCmd.Flags().BoolVarP(&param.version, "version", "v", false, "node version")
	rootCmd.AddCommand(cmd.AdminUserCmd)
	//rootCmd.Flags().StringVarP(&param.did, "did",
	//	"d", "", "admin user")

	stop = make(chan os.Signal, 8)
}
