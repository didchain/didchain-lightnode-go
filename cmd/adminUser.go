package cmd

import (
	"context"
	"encoding/json"
	"github.com/didchain/didchain-lightnode-go/config"
	"github.com/didchain/didchain-lightnode-go/pbs"
)

func (cs *cmdService)ShowAllAdminUser(ctx context.Context, r *pbs.EmptyRequest) (*pbs.CommonResponse, error)  {
	//msg := node.SrvNode().UserManagement().ShowAllUser()
	//
	//return &pbs.CommonResponse{
	//	Msg: msg,
	//},
	//nil

	msg:="no user"

	au:=config.GAdminUser
	if au == nil{
		msg = "load admin user failed"
	}else{
		s:=au.ListUser()
		j,_:=json.MarshalIndent(s," ","\t")
		msg = string(j)
	}

	return &pbs.CommonResponse{
		Msg: msg,
	},nil

}

func (cs *cmdService)ChgUser(ctx context.Context,r  *pbs.AccessAddress) (*pbs.CommonResponse, error)  {
	msg:="success"

	au:=config.GAdminUser
	if au == nil{
		msg = "load admin user failed"
	}else{
		if r.Op == 1{
			err := au.AddUser(r.Adddr)
			if err!=nil{
				msg = err.Error()
			}
		}else{
			if r.Op == 2{
				err:=au.DelUser(r.Adddr)
				if err!=nil{
					msg = err.Error()
				}
			}
		}

	}

	return &pbs.CommonResponse{
		Msg: msg,
	},nil

}