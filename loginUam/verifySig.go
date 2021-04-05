package loginUam

import (
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/didchain/didCard-go/account"
	"github.com/didchain/didchain-lightnode-go/config"
	"github.com/didchain/didchain-lightnode-go/protocol"
	"github.com/didchain/didchain-lightnode-go/user/storage"
	"github.com/didchain/didchain-lightnode-go/user/webapi"
	"github.com/kprc/nbsnetwork/tools"
	"io/ioutil"
	"net/http"
)

type UamAPI struct {
	userStorage   *storage.Storage
	sessStorage *SessStorage
	cfg *config.NodeConfig
}

func NewUamAPI(us *storage.Storage, ss *SessStorage,cfg *config.NodeConfig) *UamAPI {
	return &UamAPI{
		userStorage: us,
		sessStorage: ss,
		cfg: cfg,
	}
}

func (ua *UamAPI)Auth(w http.ResponseWriter, r *http.Request)  {
	id:=ua.sessStorage.NewSession()

	token:=&protocol.AuthContent{
		AuthUrl: ua.cfg.LoginUrl,
		RandomToken: base58.Encode(id[:]),
	}

	j,_:=json.Marshal(*token)

	w.WriteHeader(200)

	w.Write(j)
}

func internalErr(w http.ResponseWriter)  {
	resp:=&protocol.UAMResponse{}
	resp.ResultCode = 999
	resp.Message = protocol.InternalErrMsg
	w.WriteHeader(200)
	j,_:=json.Marshal(*resp)
	w.Write(j)
}

func userNotFound(w http.ResponseWriter)  {
	resp:=&protocol.UAMResponse{}
	resp.ResultCode = protocol.UserNotFound
	resp.Message = protocol.UserNotFoundMsg
	w.WriteHeader(200)
	j,_:=json.Marshal(*resp)
	w.Write(j)
}

func sigNotCorrect(w http.ResponseWriter)  {
	resp:=&protocol.UAMResponse{}
	resp.ResultCode = protocol.SignatureNotCorrect
	resp.Message = protocol.SigNatureErrMsg
	w.WriteHeader(200)
	j,_:=json.Marshal(*resp)
	w.Write(j)
}

func (ua *UamAPI)Verify(w http.ResponseWriter, r *http.Request)  {
	resp:=&protocol.UAMResponse{
		Message: protocol.SuccessMsg,
	}
	if r.Method != "POST"{
		internalErr(w)
		return
	}

	var (
		content []byte
		err error
	)
	if content,err = ioutil.ReadAll(r.Body);err!=nil{
		internalErr(w)
		return
	}
	fmt.Println("receive verify:",string(content))
	sig:=&protocol.UAMSignature{}
	err = json.Unmarshal(content,sig)
	if err!=nil{
		internalErr(w)
		return
	}

	b:=account.VerifySig(account.ID(sig.Content.DID),base58.Decode(sig.Signature),*sig.Content)
	if !b{
		sigNotCorrect(w)
		return
	}

	var udstr string

	if udstr=ua.userStorage.FindUser(sig.Content.DID);udstr == ""{
		userNotFound(w)
		webapi.Glist4add.Add(sig.Content.DID, tools.GetNowMsTime())
		return
	}

	id := Str2SessID(sig.Content.RandomToken)
	ud := &protocol.UserDesc{}
	json.Unmarshal([]byte(udstr),ud)

	si:=&SessionInfo{
		UserDesc: ud,
		Signature: sig,
	}

	ua.sessStorage.AddSession(id,si)

	w.WriteHeader(200)
	j,_:=json.Marshal(*resp)
	w.Write(j)
}

func (ua *UamAPI)Check(w http.ResponseWriter, r *http.Request) {
	resp:=&protocol.UAMResponse{
		Message: protocol.SuccessMsg,
	}
	if r.Method != "POST"{
		internalErr(w)
		return
	}

	var (
		content []byte
		err error
	)
	if content,err = ioutil.ReadAll(r.Body);err!=nil{
		internalErr(w)
		return
	}
	fmt.Println("receive check:",string(content))
	sig:=&protocol.AuthContent{}
	err = json.Unmarshal(content,sig)
	if err!=nil{
		internalErr(w)
		return
	}
	id := Str2SessID(sig.RandomToken)
	var v interface{}
	if v=ua.sessStorage.FindSession(id);v==nil{
		sigNotCorrect(w)
		return
	}

	si:=v.(*SessionInfo)

	uc:=&protocol.UAMCheck{
		RedirUrl: ua.cfg.RedirUrl,
		Signature: si.Signature,
		UserDesc: si.UserDesc,
	}

	resp.Data = uc

	ua.setCookie(id,w)

	w.WriteHeader(200)
	j,_:=json.Marshal(*resp)
	w.Write(j)

}

func (ua *UamAPI)setCookie(id SessID,w http.ResponseWriter)  {

	fmt.Println("set cookie",base58.Encode(id[:]))
	cookie:=&http.Cookie{
		Name: "sessid",
		Value: base58.Encode(id[:]),
		MaxAge: int(ua.cfg.SessiontimeOut/1000),
	}

	http.SetCookie(w,cookie)
}

func (ua *UamAPI)deleteCookie(id SessID,w http.ResponseWriter)  {
	cookie:=&http.Cookie{
		Name: "sessid",
		Value: base58.Encode(id[:]),
		MaxAge: -1,
	}

	http.SetCookie(w,cookie)

}

func (ua *UamAPI)CheckLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET"{
		internalErr(w)
		return
	}

	id,err:=r.Cookie("sessid")
	if err!=nil{
		internalErr(w)
		return
	}
    sessid:=Str2SessID(id.Value)

    var v interface{}

	if v=ua.sessStorage.FindSession(sessid);v == nil{
		internalErr(w)
		return
	}

	resp:=&protocol.UAMResponse{
		Message: protocol.SuccessMsg,
	}

	ua.setCookie(sessid,w)

	si:=v.(*SessionInfo)

	uc:=&protocol.UAMCheck{
		RedirUrl: ua.cfg.RedirUrl,
		Signature: si.Signature,
		UserDesc: si.UserDesc,
	}

	resp.Data = uc

	w.WriteHeader(200)
	j,_:=json.Marshal(*resp)
	w.Write(j)
}

func (ua *UamAPI)Logout(w http.ResponseWriter, r *http.Request)  {
	if r.Method != "GET"{
		internalErr(w)
		return
	}
	id,err:=r.Cookie("sessid")
	if err!=nil{
		internalErr(w)
		return
	}
	sessid:=Str2SessID(id.Value)

	ua.sessStorage.DelSession(sessid)

	ua.deleteCookie(sessid,w)

	resp:=&protocol.UAMResponse{
		Message: protocol.SuccessMsg,
	}

	w.WriteHeader(200)
	j,_:=json.Marshal(*resp)
	w.Write(j)
}