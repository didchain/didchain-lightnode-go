package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/didchain/didchain-lightnode-go/node"
	"github.com/didchain/didchain-lightnode-go/pbs"
	"github.com/spf13/cobra"
)

var HostCmd = &cobra.Command{
	Use: "host",
	Short: "show host",
	Long: "show host",
	Run: showAllHost,
}

var HostDelCmd = &cobra.Command{
	Use: "delete",
	Short: "delete host",
	Long: "delete host",
	Run: delAllHost,
}
var delhostname string
func init() {
	HostCmd.AddCommand(HostDelCmd)
	HostDelCmd.Flags().StringVarP(&delhostname, "host", "a", "", "host name")
}

func showAllHost(_ *cobra.Command, _ []string) {
	c := DialToCmdService()
	msg, err := c.ShowAllHost(context.TODO(), &pbs.EmptyRequest{})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(msg.Msg)
}

func delAllHost(_ *cobra.Command, _ []string) {

	if delhostname == ""{
		fmt.Println("please set host name")
		return
	}

	c := DialToCmdService()
	msg, err := c.DelHost(context.TODO(), &pbs.HostRequest{
		Host: delhostname,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	delhostname = ""
	fmt.Println(msg.Msg)
}

func (cs *cmdService)ShowAllHost(ctx context.Context, r *pbs.EmptyRequest) (*pbs.CommonResponse, error) {
	msg:="no host"

	allhosts := node.LightNodeInstance.GetUserStorage().ListAllHost()

	if len(allhosts) > 0{
		msgbytes,_:=json.Marshal(allhosts)
		msg = string(msgbytes)
	}

	return &pbs.CommonResponse{
		Msg: msg,
	},nil
}

func (cs *cmdService)DelHost(ctx context.Context, r *pbs.HostRequest) (*pbs.CommonResponse, error) {
	node.LightNodeInstance.GetUserStorage().DelHost(r.Host)
	return &pbs.CommonResponse{
		Msg: "delete success",
	},nil

}