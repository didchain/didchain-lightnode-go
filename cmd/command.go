package cmd

import (
	"context"
	"fmt"
	"github.com/didchain/didchain-lightnode-go/pbs"

	"github.com/spf13/cobra"
)

var AdminUserCmd = &cobra.Command{
	Use:   "admin-user",
	Short: "show admin user",
	Long:  `TODO::.`,
	Run:   showAllAdminUser,
}

var AdminAddUserCmd  = &cobra.Command{
	Use:   "add",
	Short: "add admin user",
	Long:  `TODO::.`,
	Run:   addAdminUser,
}
var AdminDelUserCmd  = &cobra.Command{
	Use:   "del",
	Short: "del admin user",
	Long:  `TODO::.`,
	Run:   delAdminUser,
}

var adminDid string

func init()  {
	AdminUserCmd.AddCommand(AdminAddUserCmd)
	AdminUserCmd.AddCommand(AdminDelUserCmd)

	AdminAddUserCmd.Flags().StringVarP(&adminDid,"eth-address","a","","a eth address")
	AdminDelUserCmd.Flags().StringVarP(&adminDid,"eth-address","a","","a eth address")

}

func showAllAdminUser(_ *cobra.Command, _ []string) {
	c := DialToCmdService()
	msg, err := c.ShowAllAdminUser(context.TODO(), &pbs.EmptyRequest{})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(msg.Msg)
}


func addAdminUser(_ *cobra.Command, _ []string) {
	c := DialToCmdService()
	msg, err := c.ChgUser(context.TODO(), &pbs.AccessAddress{
		Adddr: adminDid,
		Op: 1,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(msg.Msg)
}

func delAdminUser(_ *cobra.Command, _ []string) {
	c := DialToCmdService()
	msg, err := c.ChgUser(context.TODO(),  &pbs.AccessAddress{
		Adddr: adminDid,
		Op: 2,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(msg.Msg)
}
