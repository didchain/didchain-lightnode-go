package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/didchain/didCard-go/account"
	"github.com/didchain/didchain-lightnode-go/protocol"
	"github.com/ethereum/go-ethereum/crypto"
	act2 "github.com/didchain/didchain-lightnode-go/test/account"
	"github.com/kprc/nbsnetwork/tools"
	"github.com/kprc/nbsnetwork/tools/httputil"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main()  {



	//testinterface()


	//testloadwallet()

	//testw1,_:=account.NewWallet("123")
	//fmt.Println(testw1.Did().String())
	//
	//
	//testw1,_=account.NewWallet("123")
	//fmt.Println(testw1.Did().String())
	//
	//
	//testw1,_=account.NewWallet("123")
	//fmt.Println(testw1.Did().String())
	//
	//testw1,_=account.NewWallet("123")
	//fmt.Println(testw1.Did().String())
	//
	//testw1,_=account.NewWallet("123")
	//fmt.Println(testw1.Did().String())



	accesstoken,err:=gettoken()
	if err!=nil{
		fmt.Println(err)
		return
	}

	fmt.Println(accesstoken)

	var (
		w act2.Wallet

	)

	walletfile:="/Users/rickeyliao/gowork/src/github.com/didchain/didchain-lightnode-go/test/testwallet4"

	if tools.FileExists(walletfile){
		w,_=act2.LoadWallet(walletfile)

		w.Open("123")

	}else{
		w,_=act2.NewWallet("123")
		w.SaveToPath(walletfile)
	}

	fmt.Println(w.MainAddress().String())

	to := "\x19Ethereum Signed Message:\n"
	to += strconv.Itoa(len(accesstoken))
	to += accesstoken

	hash := crypto.Keccak256([]byte(to))

	sig,_:=w.Sign(hash)

	verifysig(sig,accesstoken)

	var wl account.Wallet
	wp:="/Users/rickeyliao/gowork/src/github.com/didchain/didchain-lightnode-go/test/tw2"
	if tools.FileExists(wp){
		wl,_=account.LoadWallet(wp)
		fmt.Println(wl.String())


		wl.Open("123")

	}else{
		wl,_=account.NewWallet("123")
		wl.SaveToPath(wp)
	}


	addUser(wl.Did().String(),accesstoken)
	delUser(wl.Did().String(),accesstoken)
	countUser(accesstoken)
	listUser(accesstoken)
	listunauth(accesstoken)
}



func listUser(token string)  {
	type UserReqParam struct {
		PageSize     int  `json:"page_size"`
		PageNum      int  `json:"page_num"`
	}

	type Request struct {
		AccessToken string      `json:"access_token"`
		Data        interface{} `json:"data,omitempty"`
	}
	urp:=&UserReqParam{
		PageSize: 20,
		PageNum: 0,
	}

	r:=&Request{
		AccessToken: token,
		Data: urp,
	}

	url:="http://39.99.198.143:50999/api/user/listUser"
	j,_:=json.MarshalIndent(r," ","\t")
	fmt.Println(url,"send to ",string(j))
	resp,_,err:=httputil.Post(url,string(j),false)
	if err!=nil{
		fmt.Println(err)
		return
	}
	type UserDesc struct {
		Name string	`json:"name"`
		UnitName string `json:"unit_name"`
		SerialNumber string `json:"serial_number"`
		Did string `json:"did"`
	}
	type UserListDetails struct {
		PageSize     int  `json:"page_size"`
		PageNum      int  `json:"page_num"`
		Uds []*UserDesc   `json:"uds"`
	}

	uld := &UserListDetails{}

	type Response struct {
		ResultCode int         `json:"result_code"`
		Message    string      `json:"message"`
		Data       interface{} `json:"data,omitempty"`
	}
	res:=&Response{
		Data: uld,
	}

	json.Unmarshal([]byte(resp),res)

	fmt.Println("response",resp)

}

func listunauth(token string)  {
	var b bool

	type Request struct {
		AccessToken string      `json:"access_token"`
		Data        interface{} `json:"data,omitempty"`
	}
	r := &Request{
		AccessToken: token,
		Data: &b,
	}


	url:="http://39.99.198.143:50999/api/user/listUser4Add"
	j,_:=json.MarshalIndent(r," ","\t")
	fmt.Println(url,"send to ",string(j))
	resp,_,err:=httputil.Post(url,string(j),false)
	if err!=nil{
		fmt.Println(err)
		return
	}

	type ListItem struct {
		Did string `json:"did"`
		T int64 `json:"t"`
	}

	type ListUser4Add struct {
		Dids []*ListItem `json:"dids"`
	}

	lua:=&ListUser4Add{}

	type Response struct {
		ResultCode int         `json:"result_code"`
		Message    string      `json:"message"`
		Data       interface{} `json:"data,omitempty"`
	}
	res:=&Response{
		Data: lua,
	}


	json.Unmarshal([]byte(resp),res)

	fmt.Println("response",resp)
}


func countUser(token string)  {
	var b bool

	type Request struct {
		AccessToken string      `json:"access_token"`
		Data        interface{} `json:"data,omitempty"`
	}
	r := &Request{
		AccessToken: token,
		Data: &b,
	}


	url:="http://39.99.198.143:50999/api/user/count"
	j,_:=json.MarshalIndent(r," ","\t")
	fmt.Println(url,"send to ",string(j))
	resp,_,err:=httputil.Post(url,string(j),false)
	if err!=nil{
		fmt.Println(err)
		return
	}

	var count int

	type Response struct {
		ResultCode int         `json:"result_code"`
		Message    string      `json:"message"`
		Data       interface{} `json:"data,omitempty"`
	}
	res:=&Response{
		Data: &count,
	}


	json.Unmarshal([]byte(resp),res)

	fmt.Println("response",resp)
}

func delUser(did string,token string)  {

	type Request struct {
		AccessToken string      `json:"access_token"`
		Data        interface{} `json:"data,omitempty"`
	}
	r := &Request{
		AccessToken: token,
		Data: &did,
	}


	url:="http://39.99.198.143:50999/api/user/del"
	j,_:=json.MarshalIndent(r," ","\t")
	fmt.Println(url,"send to ",string(j))
	resp,_,err:=httputil.Post(url,string(j),false)
	if err!=nil{
		fmt.Println(err)
		return
	}

	type Response struct {
		ResultCode int         `json:"result_code"`
		Message    string      `json:"message"`
		Data       interface{} `json:"data,omitempty"`
	}
	res:=&Response{}


	json.Unmarshal([]byte(resp),res)

	fmt.Println("response",resp)
}


func addUser(did string,token string)  {
	type UserDesc struct {
		Name string	`json:"name"`
		UnitName string `json:"unit_name"`
		SerialNumber string `json:"serial_number"`
		Did string `json:"did"`
	}

	type Request struct {
		AccessToken string      `json:"access_token"`
		Data        interface{} `json:"data,omitempty"`
	}

	ud:=&UserDesc{
		Name: "rickey",
		UnitName: "unit 5",
		SerialNumber: "112233",
		Did: did,
	}

	r := &Request{
		AccessToken: token,
		Data: ud,
	}

	url:="http://39.99.198.143:50999/api/user/add"
	j,_:=json.MarshalIndent(r," ","\t")

	fmt.Println("send to ",string(j))

	resp,_,err:=httputil.Post(url,string(j),false)
	if err!=nil{
		fmt.Println(err)
		return
	}

	type Response struct {
		ResultCode int         `json:"result_code"`
		Message    string      `json:"message"`
		Data       interface{} `json:"data,omitempty"`
	}

	res:=&Response{}


	json.Unmarshal([]byte(resp),res)

	fmt.Println("response",resp)

}

func verifysig(sig []byte, token string)  {
	url:="http://39.99.198.143:50999/api/auth/verify"

	type AccessSig struct {
		Sig        string `json:"sig"`
		AccesToken string `json:"acces_token"`
	}

	as:=&AccessSig{
		Sig: base58.Encode(sig),
		AccesToken: token,
	}

	j,_:=json.MarshalIndent(as," ","\t")

	fmt.Println("send to ",string(j))

	resp,_,err:=httputil.Post(url,string(j),false)
	if err!=nil{
		fmt.Println(err)
		return
	}

	type ValidSigResult struct {
		ResultCode  int    `json:"result_code"` //0 success, 1 session not found, 2 signature not correct, 3 other error
		Message     string `json:"message"`
		AccessToken string `json:"access_token"`
	}

	vr:=&ValidSigResult{}

	json.Unmarshal([]byte(resp),vr)

	fmt.Println("response",resp)

}


func gettoken() (string,error) {
	url:="http://39.99.198.143:50999/api/auth/token"
	fmt.Println(url,"get")
	resp,err:=http.Get(url)
	if err!=nil{
		fmt.Println(err)
		return "",err
	}

	r,e:=ioutil.ReadAll(resp.Body)
	if e!=nil{
		fmt.Println(e)
		return "",e
	}

	//fmt.Println(string(r))

	return string(r),nil
}


func SignMessage(did string, latitude, longitude float64, timestamp int64) string  {
	msg:= struct {
		DID       string `json:"did"` ///public key in string
		TimeStamp int64 `json:"time_stamp"`
		Latitude float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}{}

	msg.DID = did
	msg.TimeStamp = timestamp
	msg.Latitude = latitude
	msg.Longitude = longitude

	j,_:=json.Marshal(msg)

	return string(j)

}

func testloadwallet()  {
	////w,_:=account.NewWallet("123")
	//
	////fmt.Println("-----",hex.EncodeToString(w.PrivKey()))
	//wp:="/Users/rickeyliao/gowork/src/github.com/didchain/didchain-lightnode-go/test/testwallet"
	////w.SaveToPath(wp)
	//
	////fmt.Println(w.String())
	//
	////w.Close()
	//
	//wl,err:=account.LoadWallet(wp)
	//if err!=nil{
	//	panic(err.Error())
	//}
	//fmt.Println(wl.String())
	//
	//
	//
	//err = wl.Open("123")
	//if err!=nil{
	//	panic(err.Error())
	//}
	//
	//fmt.Println(hex.EncodeToString(wl.PrivKey()))
	//
	//fmt.Println("open wallet success")
	//
	//aesk,_:=wl.DriveAESKey("123")
	//
	//
	//wl.Close()
	//
	//wl.OpenWithAesKey(aesk)
	//
	//fmt.Println(hex.EncodeToString(wl.PrivKey()))


	fmt.Println(SignMessage("didaaa",11,12,11234343))

}

func testverifysig()  {
	w,_:=account.NewWallet("123")
	did := w.Did()

	vr:=&protocol.VerfiyPlainMsg{
		DID: did.String(),
		TimeStamp: 1612312312345,
		Latitude: 111.21,
		Longitude: 21.23,
	}

	sig:=w.SignJson(vr)

	req := &protocol.Request{
		Action: protocol.VerifySignature,
	}

	svr:=&protocol.VerifyReq{
		Signature: base58.Encode(sig),
		VerfiyPlainMsg:*vr,
	}

	req.PayLoad = svr

	j,_:=json.MarshalIndent(req," ","\t")

	fmt.Println(string(j))

	b:=account.VerifySig(did,sig,vr)

	fmt.Println(b)
}


func testinterface()  {
	req:=&protocol.Request{Action: protocol.VerifySignature}



	var sigbytes = make([]byte,64)
	rand.Read(sigbytes)

	vr:=&protocol.VerifyReq{
		Signature: base58.Encode(sigbytes),
		VerfiyPlainMsg:protocol.VerfiyPlainMsg{DID: base58.Encode(sigbytes[32:]),
			TimeStamp: 1234567123456,
			Latitude: 112.22,
			Longitude: 12.22,
		},
	}

	req.PayLoad = vr

	j,_:=json.MarshalIndent(req," ","\t")

	fmt.Println(string(j))

	fmt.Println(protocol.ResponseError(protocol.ErrDesc[protocol.MethodNotCorrect], protocol.MethodNotCorrect).String())

	fmt.Println(protocol.ResponseError(protocol.ErrDesc[protocol.InterIOErr], protocol.InterIOErr).String())
	fmt.Println(protocol.ResponseError(protocol.ErrDesc[protocol.UnmarshalJsonErr], protocol.UnmarshalJsonErr).String())

	fmt.Println(protocol.ResponseError(protocol.ErrDesc[protocol.ActionErr], protocol.ActionErr).String())

	resp:=&protocol.VerifyResp{Signature: vr}

	fmt.Println(protocol.ResponseSuccess(resp).String())


	fmt.Println(protocol.ResponseError(protocol.ErrDesc[protocol.SignatureNotCorrect], protocol.SignatureNotCorrect).String())
}