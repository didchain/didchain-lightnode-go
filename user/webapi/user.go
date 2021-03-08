package webapi

import (
	"encoding/json"
	"github.com/didchain/didchain-lightnode-go/user/storage"
	"net/http"
	"sync"
)

type UserAPI struct {
	sdb *storage.Storage
}

type ListItem struct {
	Did string `json:"did"`
	T int64 `json:"t"`
}

type ListUser4Add struct {
	Dids []*ListItem `json:"dids"`
}

var glist4add *ListUser4Add
var glistlock sync.Mutex

func init()  {
	glist4add = &ListUser4Add{}
}

func (lu *ListUser4Add)add(did string, t int64)  {
	glistlock.Lock()
	defer glistlock.Unlock()
	lu.Dids = append(lu.Dids, &ListItem{Did: did,T: t})
	l:=len(lu.Dids)
	if  l> 3{
		lu.Dids = lu.Dids[l-3:l]
	}
	return
}


func NewUserAPI(sdb *storage.Storage) *UserAPI  {
	return &UserAPI{sdb: sdb}
}

type UserDesc struct {
	Name string	`json:"name"`
	UnitName string `json:"unit_name"`
	SerialNumber string `json:"serial_number"`
	Did string `json:"did"`
}



func (ua *UserAPI)AddUser(w http.ResponseWriter,r *http.Request){
	var ud UserDesc
	 _,resp:=doRequest(r,&ud)
	if resp.ResultCode > 0 {
		j, _ := json.Marshal(*resp)
		w.WriteHeader(200)
		w.Write(j)
		return
	}

	ua.sdb.AddUser(ud.Did,&ud)
	j, _ := json.Marshal(*resp)
	w.WriteHeader(200)
	w.Write(j)
}



func (ua *UserAPI)DelUser(w http.ResponseWriter,r *http.Request){
	var user string
	_,resp:=doRequest(r,&user)
	if resp.ResultCode > 0 {
		j, _ := json.Marshal(*resp)
		w.WriteHeader(200)
		w.Write(j)
		return
	}

	ua.sdb.DelUser(user)
	j, _ := json.Marshal(*resp)
	w.WriteHeader(200)
	w.Write(j)
}

type UserReqParam struct {
	PageSize     int  `json:"page_size"`
	PageNum      int  `json:"page_num"`
}

type UserListDetails struct {
	PageSize     int  `json:"page_size"`
	PageNum      int  `json:"page_num"`
	Uds []*UserDesc   `json:"uds"`
}

func (ua *UserAPI)ListUser(w http.ResponseWriter,r *http.Request)  {
	mrp := &UserReqParam{}

	req, resp := doRequest(r, mrp)
	if resp.ResultCode > 0 {
		j, _ := json.Marshal(*resp)
		w.WriteHeader(200)
		w.Write(j)
		return
	}

	param := req.Data.(*UserReqParam)
	data := &UserListDetails{PageNum: param.PageNum, PageSize: param.PageSize}

	resp.Data = data


	uas:=ua.sdb.ListAllValue2(func(data []byte) interface{} {
		ud := &UserDesc{}
		json.Unmarshal(data,ud)
		return ud
	},param.PageNum*param.PageSize,param.PageSize)

	for i:=0;i<len(uas);i++{
		data.Uds = append(data.Uds,uas[i].(*UserDesc))
	}

	j, _ := json.Marshal(*resp)

	w.WriteHeader(200)
	w.Write(j)

}

func (ua *UserAPI)UserCount(w http.ResponseWriter,r *http.Request)   {
	var forceRefresh bool
	_, resp := doRequest(r, &forceRefresh)
	if resp.ResultCode > 0 {
		j, _ := json.Marshal(*resp)
		w.WriteHeader(200)
		w.Write(j)
		return
	}

	count:=len(ua.sdb.ListAll())

	resp.Data =  &count

	j, _ := json.Marshal(*resp)

	w.WriteHeader(200)
	w.Write(j)


}



func (ua *UserAPI)ListUnAuthorizeUser(w http.ResponseWriter,r *http.Request)  {
	var refresh bool
	_,resp:=doRequest(r,&refresh)
	if resp.ResultCode > 0 {
		j, _ := json.Marshal(*resp)
		w.WriteHeader(200)
		w.Write(j)
		return
	}
	glistlock.Lock()
	defer glistlock.Unlock()
	resp.Data = *glist4add

	j, _ := json.Marshal(*resp)
	w.WriteHeader(200)
	w.Write(j)

}

